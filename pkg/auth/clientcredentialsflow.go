package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ClientCredentialConfig struct {
	ClientId     string
	ClientSecret string
	Host         string
}

func NewConfiguration(clientId string, clientSecret string, host string) *ClientCredentialConfig {
	return &ClientCredentialConfig{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		Host:         host,
	}
}

var _ AuthResponse = &ClientCredentialApi{}

type ClientCredentialApi struct {
	config *ClientCredentialConfig
	*clientCredentialsResp
}

func NewClientCredentialApi(config *ClientCredentialConfig) *ClientCredentialApi {
	return &ClientCredentialApi{
		config:                config,
		clientCredentialsResp: &clientCredentialsResp{},
	}
}

type clientCredentialsResp struct {
	AccessToken string `json:"access_token"`
	InstanceUrl string `json:"instance_url"`
	Id          string `json:"id"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
	IssuedAt    string `json:"issued_at"`
	Signature   string `json:"signature"`
}

func (c *ClientCredentialApi) GetAccessToken() string {
	return c.AccessToken
}

func (c *ClientCredentialApi) GetID() string {
	return c.Id
}

func (c *ClientCredentialApi) GetInstanceURL() string {
	return c.InstanceUrl
}

func (c *ClientCredentialApi) GetScope() string {
	return c.Scope
}

func (c *ClientCredentialApi) GetTokenType() string {
	return c.TokenType
}

func (c *ClientCredentialApi) Auth() AuthResponse {
	url := fmt.Sprintf("%s/services/oauth2/token", c.config.Host)

	reqBodyStr := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", c.config.ClientId, c.config.ClientSecret)

	req, err := http.NewRequest("POST", url, strings.NewReader(reqBodyStr))
	if err != nil {
		panic(err)
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	//json marshal body
	err = json.Unmarshal(body, c.clientCredentialsResp)
	if err != nil {
		panic(err)
	}

	return c
}
