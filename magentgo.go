package magentgo

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	apiBaseUrl             string
	client                 *http.Client
	clientConfigValidators []ClientConfigValidationFunc
	baseUrl                string
	bearerToken            string
	// defaults to
	storeCode string
	// defaults to 1
	version int

	AuthService    *AuthService
	ProductService *ProductService
}

// create a new instance of api client, function options for configuration
func New(options ...OptionFunc) (*Client, error) {
	client := &Client{
		client:      &http.Client{},
		baseUrl:     "",
		bearerToken: "",
		storeCode:   "",
		version:     1,
	}

	// option validation files just return an error
	for _, option := range options {
		err := option(client)
		if err != nil {
			return nil, err
		}
	}

	err := client.validate()
	if err != nil {
		return nil, err
	}

	client.setApiBaseUrl()
	client.assignServices()

	return client, nil
}

func (c *Client) assignServices() {
	c.AuthService = &AuthService{client: c}
	c.ProductService = &ProductService{client: c}
}

func (c *Client) validate() error {
	c.clientConfigValidators = append(c.clientConfigValidators,
		ValidateBaseUrl,
		ValidateBearerToken,
		ValidateVersion,
	)

	for _, validator := range c.clientConfigValidators {
		err := validator(c)
		if err != nil {
			return err
		}
	}

	return nil
}

// build base url used for all requests
func (c *Client) setApiBaseUrl() *Client {
	apiType := "rest/"
	storeCode := ""
	if c.storeCode != "" {
		storeCode = c.storeCode + "/"
	}
	version := fmt.Sprintf("V%d/", c.version)

	c.apiBaseUrl = c.baseUrl + apiType + storeCode + version

	return c
}

// configure base url
func (c *Client) setBaseUrl(url string) error {
	if !strings.HasSuffix(url, "/") {
		url += "/"
	}

	if !IsUrl(url) {
		return errors.New("base URL is invalid")
	}

	c.baseUrl = url

	return nil
}

// configure bearertoken
func (c *Client) setBearerToken(token string) error {
	if token == "" {
		return errors.New("empty string is not a valid bearer token")
	}

	c.bearerToken = token

	return nil
}

// set store code in API url
func (c *Client) setStoreCode(storeCode string) error {
	c.storeCode = storeCode

	return nil
}

// configure api version
func (c *Client) setVersion(version int) error {
	c.version = version

	return nil
}

// make http request from the client. Attempts to marshal against type struct, returns raw result as byte slice
func (c *Client) call(endpoint string, httpVerb string, bodyType any, responseType any, ctx context.Context) ([]byte, error) {
	var err error

	requestUrl := fmt.Sprintf("%s%s", c.apiBaseUrl, endpoint)

	fmt.Println(requestUrl)

	body, err := c.processBodyType(bodyType)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest(ctx, httpVerb, requestUrl, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	c.setRequestHeaders(req, c.bearerToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if responseType != nil {
		err = json.Unmarshal(body, responseType)
		if err != nil {
			return body, err
		}
	}

	return body, nil
}

// type assertions to handle body type, converts whichever type provided into a byte slice
func (c *Client) processBodyType(body any) ([]byte, error) {
	var marshalled []byte
	var err error
	switch v := body.(type) {
	case string:
		marshalled = []byte(v)
	case []byte:
		marshalled = v
	default:
		marshalled, err = json.Marshal(&body)
		if err != nil {
			return nil, err
		}
	}

	return marshalled, nil
}

// create new request with context to be used in main call function
func (c *Client) newRequest(ctx context.Context, verb string, url string, reader *bytes.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, verb, url, reader)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// implement request headers
func (c *Client) setRequestHeaders(req *http.Request, token string) *Client {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	return c
}
