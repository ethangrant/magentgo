package magentgo

import (
	"context"
	"encoding/json"
)

type AuthService struct {
	client *Client
}

type AuthPost struct {
	ErrorResponse
	Username string `json:"username"`
	Password string `json:"password"`
	Otp      string `json:"otp"`
}

// auth endpoints respond with "string"
type AuthResponse string

// Auth end points just return strings so we can unmarshal to struct
type AuthResponseWithError struct {
	ErrorResponse
	Token string
}

func (a *AuthService) AuthToken(endpoint string, username string, password string, ctx context.Context) (AuthResponseWithError, error) {
	var authRes AuthResponse = ""
	authResponseWithError := &AuthResponseWithError{}

	authPost := AuthPost{Username: username, Password: password}
	raw, err := a.client.call(endpoint, "POST", &authPost, &authRes, ctx)
	if err != nil {
		err = json.Unmarshal(raw, &authResponseWithError)
		if err != nil {
			return *authResponseWithError, err
		}
		return *authResponseWithError, err
	}

	authResponseWithError.Token = string(authRes)

	// update bearer token on client so user does not have to.
	a.client.setBearerToken(authResponseWithError.Token)

	return *authResponseWithError, err
}

// request token with admin credentials
func (a *AuthService) AdminToken(username string, password string, ctx context.Context) (AuthResponseWithError, error) {
	res, err := a.AuthToken("integration/admin/token", username, password, ctx)
	if err != nil {
		return res, err
	}

	return res, nil
}

// request token with customer credentials
func (a *AuthService) CustomerToken(username string, password string, ctx context.Context) (AuthResponseWithError, error) {
	res, err := a.AuthToken("integration/customer/token", username, password, ctx)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (a *AuthService) AdminWithGoogleAuthenticator(username string, password string, otp string) {
	// TODO:
}

func (a *AuthService) AdminWithAuthy(username string, password string, otp string) {
	// TODO:
}

func (a *AuthService) AdminWithU2fKey(username string, password string, otp string) {
	// TODO:
}
