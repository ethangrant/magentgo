package magentgo

import "errors"

type ClientConfigValidationFunc func(*Client) error

// validates base url of client config
func ValidateBaseUrl(c *Client) error {
	isEmpty := validateEmptyString(c.baseUrl)
	if isEmpty {
		return errors.New("base url cannot be an empty string")
	}

	return nil
}

// validates bearer token of client config
func ValidateBearerToken(c *Client) error {
	isEmpty := validateEmptyString(c.baseUrl)
	if isEmpty {
		return errors.New("bearer token cannot be an empty string")
	}

	return nil
}

// one day there might be a v2...
func ValidateVersion(c *Client) error {
	if c.version != 1 {
		return errors.New("magento 2 currently only support v1 API")
	}
	return nil
}

// returns true if string is empty
func validateEmptyString(s string) bool {
	return s == ""
}
