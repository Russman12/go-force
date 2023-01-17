# go-force

## Overview
This project is a documentation first implementation of a salesforce client for golang. OpenApi documents are used to generate teh documentation then cade is automatically generated using the documentation provided.

## How it works
All the code in this project is automatically generated via [openapi-generator](https://github.com/OpenAPITools/openapi-generator).
There are several openapi files which and are used to automatically generate the code necessary to make http requests to the various salesforce endpoints. The openapi documents can also be viewed to better understand how the APIs work in salesforce.

## Installation

Add the following import:
`import "github.com/russman12/go-force/pkg/rest"`

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

The server URL is automatically determined based on the oauth2 response from salesforce.

## Usage
The example below demonstrates how to send a request to the (REST Resources)[https://developer.salesforce.com/docs/atlas.en-us.api_rest.meta/api_rest/dome_discoveryresource.htm].

```golang
package main

import (
	"fmt"
	"golang.org/x/oauth2"
	"context"
)

func main() {
	oAuth2Cfg := clientcredentials.Config{}
	tokenSrc := oAuth2Cfg.TokenSource(context.Background())
	packageName.NewAPIClient(packageName.NewConfiguration(), tokenSrc)

	restClient := rest.NewAPIClient(rest.NewConfiguration(), tokenSrc)
	resources, _, err := restClient.OrgApi.GetResources(context.Background()).Execute()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", resources)
}
```

See API documentation for implementation details

## Supported APIs
* [Bulk v2](pkg/bulkv2/README.md)
* [Rest](pkg/rest/README.md) (partially implemented)
