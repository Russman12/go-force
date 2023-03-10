/*
Salesforce Platform REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 56.0
Contact: russell-laboe@outlook.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tooling

import (
    "fmt"
	"net/http"
	"strings"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextBasicAuth takes BasicAuth as authentication for the request.
	ContextBasicAuth = contextKey("basic")

	// ContextAccessToken takes a string oauth2 access token as authentication for the request.
	ContextAccessToken = contextKey("accesstoken")

	// ContextAPIKeys takes a string apikey as authentication for the request
	ContextAPIKeys = contextKey("apiKeys")

	// ContextHttpSignatureAuth takes HttpSignatureAuth as authentication for the request.
	ContextHttpSignatureAuth = contextKey("httpsignature")

	// ContextOperationServerIndices uses a server configuration from the index mapping.
	ContextOperationServerIndices = contextKey("serverOperationIndices")

	// ContextOperationServerVariables overrides a server configuration variables using operation specific values.
	ContextOperationServerVariables = contextKey("serverOperationVariables")
)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}

// ServerVariable stores the information about a server variable
type ServerVariable struct {
    description  string
    defaultValue string
    currentValue string
    enumValues   []string
}

// ServerConfiguration stores the information about a server
type ServerConfiguration struct {
    url         string
    description string
    variables   map[string]*ServerVariable
}

// ServerConfigurations stores multiple ServerConfiguration items
type ServerConfigurations []*ServerConfiguration

// Configuration stores the configuration of the API client
type Configuration struct {
	Host             string            `json:"host,omitempty"`
	Scheme           string            `json:"scheme,omitempty"`
	DefaultHeader    map[string]string `json:"defaultHeader,omitempty"`
	UserAgent        string            `json:"userAgent,omitempty"`
	Debug            bool              `json:"debug,omitempty"`
    servers           ServerConfigurations
    activeServerIndex int
    operationServers  map[string]ServerConfigurations
    httpClient        *http.Client
}

// NewConfiguration returns a new Configuration object
func NewConfiguration() *Configuration {
	cfg := &Configuration{
		DefaultHeader:    make(map[string]string),
		UserAgent:        "OpenAPI-Generator/1.0.0/go",
		Debug:            false,
		servers:          ServerConfigurations{
			{
				url: "{instanceUrl}/services/data/v{apiVersion}/tooling",
				description: "API Base URL",
				variables: map[string]*ServerVariable{
					"instanceUrl": &ServerVariable{
						description: "Salesforce server domain",
						defaultValue: "https://myorg.lightning.force.com",
                        currentValue: "https://myorg.lightning.force.com",
					},
					"apiVersion": &ServerVariable{
						description: "Salesforce api version",
						defaultValue: "56.0",
                        currentValue: "56.0",
					},
				},
			},
		},
        activeServerIndex: 0,
		operationServers: map[string]ServerConfigurations{
		},
	}
	return cfg
}

// NewConfigurationWithActiveServerVars returns a new Configuration object with the default server and variable replacements set
func NewConfigurationWithActiveServerVars(index int, variables map[string]string) (*Configuration, error) {
    cfg := NewConfiguration()
    if index < 0 || index >= len(cfg.servers) {
        return nil, fmt.Errorf("Server index out of range")
    }
    cfg.activeServerIndex = index
    for name, val := range variables {
        cfg.servers[cfg.activeServerIndex].SetServerVariable(name, val)
    }
    return cfg, nil
}

// SetActiveServer sets the current active server configuration
func (c *Configuration) SetActiveServer(index int) error {
    if index >= len(c.servers) || index < 0 {
        return fmt.Errorf("Server index out of range")
    }
    c.activeServerIndex = index
    return nil
}

// GetActiveServer returns the current active server configuration
func (c *Configuration) GetActiveServer() *ServerConfiguration {
    return c.servers[c.activeServerIndex]
}

// GetServer returns server configuration at given index
func (c *Configuration) GetServer(index int) (*ServerConfiguration, error) {
    if index >= len(c.servers) || index < 0 {
        return nil, fmt.Errorf("Server index out of range")
    }

    return c.servers[index], nil
}

func (sc *ServerConfiguration) SetServerVariable(name string, value string) {
    if _, ok := sc.variables[name]; !ok {
        sc.variables[name] = &ServerVariable{currentValue: value}
    }
    sc.variables[name].currentValue = value
}



func (sc *ServerConfiguration) GetServerVariable(name string) string {
    if srvVar, ok := sc.variables[name]; ok {
        return srvVar.currentValue
    }
    return ""
}

func (sc *ServerConfiguration) ResetServerVariable(name string) error {
    if _, ok := sc.variables[name]; !ok {
        return fmt.Errorf("unable to locate server variable")
    }
    sc.variables[name].currentValue = sc.variables[name].defaultValue
    return nil
}

// AddDefaultHeader adds a new HTTP header to the default header in the request
func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}

// GetURL gets the url for a server config with variables injected
func (sc ServerConfiguration) GetURL() string {
    urlVars := map[string]string{}
    for varName, serverVar := range sc.variables {
        urlVars[varName] = serverVar.currentValue
    }
    return injectUrlVars(sc.url, urlVars)
}

// injectUrlVars formats the given string with injected variables
func injectUrlVars(url string, variables map[string]string) string {
    // go through variables and replace placeholders
    for name, value := range variables {
        url = strings.Replace(url, "{"+name+"}", value, -1)
    }
    return url
}
