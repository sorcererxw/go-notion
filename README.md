# go-notion

[![tests](https://github.com/sorcererxw/go-notion/actions/workflows/tests.yaml/badge.svg)](https://github.com/sorcererxw/go-notion/actions/workflows/tests.yaml)
[![Go Reference](https://pkg.go.dev/badge/github.com/sorcererxw/go-notion.svg)](https://pkg.go.dev/github.com/sorcererxw/go-notion)
[![Go Report Card](https://goreportcard.com/report/github.com/sorcererxw/go-notion)](https://goreportcard.com/badge/github.com/sorcererxw/go-notion)
[![codecov](https://codecov.io/gh/sorcererxw/go-notion/branch/master/graph/badge.svg?token=BUSFRL18RV)](https://codecov.io/gh/sorcererxw/go-notion)

> 🚧 Working In Progress

Go SDK for Notion Official API.

```shell
go get github.com/sorcererxw/go-notion
```

* [Overview](#overview)
* [Getting Started](#getting-started)
    - [Pagination](#pagination)
    - [Error Handling](#error-handling)
    - [Reverse Proxy](#reverse-proxy)
    - [OAuth](#oauth)
* [License](#license)

## Overview

go-notion is the Golang binding
for [Notion official API](https://developers.notion.com/).
This package provides:

- Easy-to-use and well-testing API wrappers.

- Complete type definition.

You can easily and quickly build notion integrations with
this package.

⚠️ Notion official API is still in public beta, it's hard to
guarantee forward compatibility in the future. This package
will be continuously updated according to the official
documentation.

## Getting Started

At the beginning, you should follow the
official [document](https://developers.notion.com/docs) to
create your workspace and integrations.

```go
package main

import (
	"context"

	"github.com/sorcererxw/go-notion"
)

func main() {
	client := notion.NewClient(notion.Settings{Token: "token"})

	database, err := client.RetrieveDatabase(context.Background(), "database_id")
}
```

### Pagination

```go
package main

func main() {
	var cursor string
	for {
		data, nextCursor, hasMore, err := client.ListAllUsers(context.Background(), 30, cursor)
		if err != nil {
			break
		}
		if !hasMore {
			break
		}
		cursor = nextCursor
	}
}
```

### Error Handling

go-notion
declares [error codes](https://developers.notion.com/reference/errors)
. You can compare error code to confirm which error
occurred.

```go
package main

import (
	"context"
	"fmt"

	"github.com/sorcererxw/go-notion"
)

func main() {
	user, err := client.RetrieveUser(context.Background(), "user_id")
	if err, ok := notion.AsError(err); ok {
		switch err.Code {
		case notion.ErrCodeRateLimited:
			fmt.Println("rate limited")
		}
	}
}
```

### Reverse Proxy

If you cannot access Notion server in your region(e.g. China)
directly, you can use reverse proxy to solve the problem:

```go
package main

import "github.com/sorcererxw/go-notion"

const proxyEndpoint = "https://1.1.1.1/notion"

func main() {
  client := notion.NewClient(notion.Settings{
      Token: "token",
      Endpoint: proxyEndpoint,
  })
}
```

### OAuth

```go
package main

import "net/http"

func main() {
  client := notion.NewOAuthClient("client_id", "client_secret", "redirect_uri")

  mux := http.NewServeMux()
  mux.HandleFunc("/oauth", func(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    token, _ := client.ExchangeAccessToken(r.Context(), code)
    
    // store token to db ...
    
    w.WriteHeader(http.StatusOK)
  })
}
```

## License

go-notion is distributed under [MIT](./LICENSE).