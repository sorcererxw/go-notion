package notion

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Client struct {
	token      string
	httpclient *http.Client
}

func NewClient(token string) API {
	return &Client{
		token: token,
		httpclient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) RetrieveDatabase(ctx context.Context, databaseID string) (*Database, error) {
	var database Database
	if err := c.request(ctx, http.MethodGet, "/v1/databases/"+databaseID, nil, &database); err != nil {
		return nil, err
	}
	return &database, nil
}

func (c *Client) QueryDatabase(ctx context.Context, databaseID string, param QueryDatabaseParam) ([]*Page, string, bool, error) {
	var result List
	if err := c.request(ctx, http.MethodPost, "/v1/databases/"+databaseID+"/query", param, &result); err != nil {
		return nil, "", false, err
	}
	return result.Results.Pages(), result.NextCursor, result.HasMore, nil
}

func (c *Client) ListDatabases(ctx context.Context, pageSize int32, startCursor string) ([]*Database, string, bool, error) {
	var result List
	if err := c.request(ctx, http.MethodPost, "/v1/databases", nil, &result, c.concatPagination(pageSize, startCursor)); err != nil {
		return nil, "", false, err
	}
	return result.Results.Databases(), result.NextCursor, result.HasMore, nil
}

func (c *Client) RetrievePage(ctx context.Context, pageID string) (*Page, error) {
	var page Page
	if err := c.request(ctx, http.MethodGet, "/v1/pages/"+pageID, nil, &page); err != nil {
		return nil, err
	}
	return &page, nil
}

func (c *Client) CreatePage(ctx context.Context, parent Parent, properties map[string]*Property, children ...*Block) (*Page, error) {
	body := struct {
		Parent     Parent               `json:"parent,omitempty"`
		Properties map[string]*Property `json:"properties,omitempty"`
		Children   []*Block             `json:"children,omitempty"`
	}{
		Parent:     parent,
		Properties: properties,
		Children:   children,
	}
	var page Page
	if err := c.request(ctx, http.MethodGet, "/v1/pages", body, &page); err != nil {
		return nil, err
	}
	return &page, nil
}

func (c *Client) UpdatePageProperties(ctx context.Context, pageID string, properties map[string]*Property) (*Page, error) {
	body := struct {
		Properties map[string]*Property `json:"properties,omitempty"`
	}{
		Properties: properties,
	}
	var page Page
	if err := c.request(ctx, http.MethodPatch, "/v1/pages/"+pageID, body, &page); err != nil {
		return nil, err
	}
	return &page, nil
}

func (c *Client) RetrieveBlockChildren(ctx context.Context, blockID string, pageSize int32, startCursor string) ([]*Block, string, bool, error) {
	var result List
	if err := c.request(ctx, http.MethodGet, "/v1/blocks/"+blockID+"/children", nil, &result, c.concatPagination(pageSize, startCursor)); err != nil {
		return nil, "", false, err
	}
	return result.Results.Blocks(), result.NextCursor, result.HasMore, nil
}

func (c *Client) AppendBlockChildren(ctx context.Context, blockID string, children ...*Block) error {
	body := struct {
		Children []*Block `json:"children"`
	}{
		Children: children,
	}
	var block Block
	if err := c.request(ctx, http.MethodPatch, "/v1/blocks/"+blockID+"/children", body, &block); err != nil {
		return err
	}
	return nil
}

func (c *Client) RetrieveUser(ctx context.Context, userID string) (*User, error) {
	var user User
	if err := c.request(ctx, http.MethodGet, "/v1/users/"+userID, nil, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *Client) ListAllUsers(ctx context.Context, pageSize int32, startCursor string) ([]*User, string, bool, error) {
	var result List
	if err := c.request(ctx, http.MethodGet, "/v1/users", nil, &result, c.concatPagination(pageSize, startCursor)); err != nil {
		return nil, "", false, err
	}
	return result.Results.Users(), result.NextCursor, result.HasMore, nil
}

func (c *Client) Search(ctx context.Context, param SearchParam) ([]*Object, string, bool, error) {
	var result List
	if err := c.request(ctx, http.MethodGet, "/v1/search", param, &result); err != nil {
		return nil, "", false, err
	}
	return result.Results, result.NextCursor, result.HasMore, nil
}

func (c *Client) request(ctx context.Context, method string, path string, in interface{}, out interface{}, fns ...func(req *http.Request)) error {
	var body io.Reader
	if in != nil {
		b, err := json.Marshal(in)
		if err != nil {
			return err
		}
		body = bytes.NewBuffer(b)
	}
	req, err := http.NewRequestWithContext(ctx, method, endpoint+path, body)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Bearer "+c.token)
	req.Header.Add("Notion-Version", apiVersion)

	for _, fn := range fns {
		fn(req)
	}

	rsp, err := c.httpclient.Do(req)
	if err != nil {
		return err
	}

	defer rsp.Body.Close()

	if rsp.StatusCode >= 400 {
		if rsp.Header.Get("Content-Type") != "application/json" {
			return fmt.Errorf("HTTP/%d: http request error", rsp.StatusCode)
		}
		var e Error
		if err := json.NewDecoder(rsp.Body).Decode(&e); err != nil {
			return err
		}
		return &e
	}

	if out != nil {
		if err := json.NewDecoder(rsp.Body).Decode(out); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) concatPagination(pageSize int32, startCursor string) func(req *http.Request) {
	return func(req *http.Request) {
		if req == nil {
			return
		}
		q := req.URL.Query()
		q.Add("page_size", strconv.Itoa(int(pageSize)))
		q.Add("start_cursor", startCursor)
		req.URL.RawQuery = q.Encode()
	}
}
