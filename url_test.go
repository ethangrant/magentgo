package magentgo

import (
	"strconv"
	"testing"
)

func TestIsUrl(t *testing.T) {
	urls := map[string]bool{
		"https://www.example.com":                           true,
		"ftp://example.com/resource.txt":                    false,
		"http://www.invalid-url-with-spaces .com":           false,
		"https://sub.domain.example.com/path?query=string":  true,
		"htp://misspelled-protocol.com":                     false,
		"https://www.valid-url.com/page#section":            true,
		"http://":                                           false,
		"https://www.example-with-dash.com":                 true,
		"http://www.example_with_underscore.com":            false,
		"https://example.com:8080":                          true,
		"://missing-protocol.com":                           false,
		"http://www.example.com?query=param&another=param2": true,
		"https://example.com/":                              true,
		"https://example.com/path/to/resource":              true,
		"http://.example.com":                               false,
		"https://example.com#fragment":                      true,
		"invalid://example.com":                             false,
	}

	for url, valid := range urls {
		isUrl := IsUrl(url)

		if valid != isUrl {
			t.Errorf("Result was incorrect for url %s. Got: %s, Want: %s", url, strconv.FormatBool(isUrl), strconv.FormatBool(valid))
		}
	}
}
