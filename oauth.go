package notion

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type OAuthAccessToken struct {
	AccessToken   string `json:"access_token,omitempty"`
	WorkspaceName string `json:"workspace_name,omitempty"`
	WorkspaceIcon string `json:"workspace_icon,omitempty"`
	BotID         string `json:"bot_id,omitempty"`
}

type OAuthClient struct {
	clientID     string
	clientSecret string
	redirectURI  string
	httpclient   *http.Client
}

func NewOAuthClient(clientID, clientSecret, redirectURI string) *OAuthClient {
	return &OAuthClient{
		clientID:     clientID,
		clientSecret: clientSecret,
		redirectURI:  redirectURI,
		httpclient:   &http.Client{},
	}
}

func (c *OAuthClient) ExchangeAccessToken(ctx context.Context, code string) (*OAuthAccessToken, error) {
	b, err := json.Marshal(&struct {
		GrantType   string `json:"grant_type,omitempty"`
		Code        string `json:"code,omitempty"`
		RedirectURI string `json:"redirect_uri,omitempty"`
	}{
		GrantType:   "authorization_code",
		Code:        code,
		RedirectURI: c.redirectURI,
	})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://api.notion.so/v1/oauth/token", bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.clientID, c.clientSecret)
	req.Header.Add("Content-Type", "application/json")

	rsp, err := c.httpclient.Do(req)
	if err != nil {
		return nil, err
	}

	defer rsp.Body.Close()

	var body OAuthAccessToken

	if err := json.NewDecoder(rsp.Body).Decode(&body); err != nil {
		return nil, err
	}
	return &body, nil
}
