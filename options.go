package magentgo

type OptionFunc func(*Client) error

// set base url for all api requests, should be the store base URL
func WithBaseURl(url string) OptionFunc {
	return func(c *Client) error {
		return c.setBaseUrl(url)
	}
}

// set bearer token if using it
func WithBearerToken(token string) OptionFunc {
	return func(c *Client) error {
		return c.setBearerToken(token)
	}
}

// set api version number for all requests
func WithVersion(version int) OptionFunc {
	return func(c *Client) error {
		return c.setVersion(version)
	}
}
