package magentgo

import (
	"errors"
	"net/http"
	"testing"
)

// TODO: revist once implemented more of the client.
func TestValidate(t *testing.T) {
	tests := []struct {
		name        string
		client      *Client
		expectedErr error
	}{
		{
			name: "valid client",
			client: &Client{
				client:      &http.Client{},
				baseUrl:     "https://www.magento2store.co.uk/",
				bearerToken: "somebearertoken",
				version:     1,
			},
			expectedErr: nil,
		},
		{
			name: "invalid client no base url",
			client: &Client{
				client:      &http.Client{},
				baseUrl:     "",
				bearerToken: "somebearertoken",
				version:     1,
			},
			expectedErr: errors.New("base url cannot be an empty string"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.client.validate()
			if (err != nil && tt.expectedErr == nil) || (err == nil && tt.expectedErr != nil) {
				t.Errorf("expected error %v, got %v", tt.expectedErr, err)
			} else if err != nil && err.Error() != tt.expectedErr.Error() {
				t.Errorf("expected error %v, got %v", tt.expectedErr.Error(), err.Error())
			}
		})
	}
}

func TestSetBaseUrl(t *testing.T) {
	tests := []struct{
		name string
		url string
		expectedBaseUrl string
		expectedErr error
	} {
		{
			name: "valid base url",
			url: "https://www.magento2store.co.uk/",
			expectedBaseUrl: "https://www.magento2store.co.uk/",
			expectedErr: nil,
		},
		{
			name: "base url no trailing slash",
			url: "https://www.magento2store.co.uk",
			expectedBaseUrl: "https://www.magento2store.co.uk/",
			expectedErr: nil,
		},
		{
			name: "invalid base url",
			url: "invalidbaseurl.co",
			expectedBaseUrl: "",
			expectedErr: errors.New("base URL is invalid"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				client:      &http.Client{},
				baseUrl:     "",
				bearerToken: "",
				version:     1,
			}

			err := client.setBaseUrl(tt.url)
			if (err != nil && tt.expectedErr == nil) || (err == nil && tt.expectedErr != nil) {
				t.Errorf("expected error %v, got %v", tt.expectedErr, err)
			} else if err != nil && err.Error() != tt.expectedErr.Error() {
				t.Errorf("expected error %v, got %v", tt.expectedErr.Error(), err.Error())
			}

			if client.baseUrl != tt.expectedBaseUrl {
				t.Errorf("expected base url %s, got %s", tt.expectedBaseUrl, client.baseUrl)
			}
		})
	}
}