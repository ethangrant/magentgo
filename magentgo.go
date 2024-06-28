package magentgo

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type Client struct {
	apiBaseUrl string
	client                 *http.Client
	clientConfigValidators []ClientConfigValidationFunc
	baseUrl     string
	bearerToken string
	// defaults to 
	storeCode string
	// defaults to 1
	version int
}

// create a new instance of api client, function options for configuration
func New(options ...OptionFunc) (*Client, error) {
	client := &Client{
		client:      &http.Client{},
		baseUrl:     "",
		bearerToken: "",
		storeCode: "",
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

	return client, nil
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
