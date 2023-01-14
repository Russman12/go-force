# go-force

## Overview
This project is a documentation first implementation of a salesforce client for golang. OpenApi documents are used to generate teh documentation then cade is automatically generated using the documentation provided.

## How it works
All the code in this project is automatically generated via [openapi-generator](https://github.com/OpenAPITools/openapi-generator).
There are several openapi files which and are used to automatically generate the code necessary to make http requests to the various salesforce endpoints. The openapi documents can also be viewed to better understand how the APIs work in salesforce.

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import rest "github.com/russman12/go-force/rest"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), rest.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), rest.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), rest.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), rest.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Usage
To start using this library, first download it via the following command:

```bash
go get github.com/russman12/go-force
```

Next, add the following lines of code to your project to log in to salesforce:

```golang
//create auth client with default config
authClient := auth.NewAPIClient(auth.NewConfiguration())

//use test.salesforce.com server
authCtx := context.WithValue(context.Background(), auth.ContextServerIndex, 0)

//log in to salesforce
auth, _, _ := authClient.OAuthApi.
	OAuthUserPass(authCtx).
	ClientId("clientId").
	ClientSecret("clientSecret").
	Username("username").
	Password("password").
	GrantType("password").
	Execute()
```

Now that you are logged into salesforce, you can use the auth data to set a default header on your desired API. You can also use the returned instance url to set the instance url in a new context for your desired API. For example:

```golang
//create config with default auth header set
restCfg := rest.NewConfiguration()
restCfg.AddDefaultHeader("Authorization", fmt.Sprintf("%s %s", auth.TokenType, auth.AccessToken))

//create new context with instanceUrl injected into url
restCtx := context.WithValue(
	context.Background(),
	rest.ContextServerVariables,
	map[string]string{"instanceUrl": auth.InstanceUrl}
)

restClient := rest.NewAPIClient(restCfg)
```

See API documentation for implementation details

## Supported APIs
* [Auth](pkg/auth/README.md) (partial)
* [Bulk v2](pkg/bulkv2/README.md)
* [Rest](pkg/rest/README.md) (partial)
