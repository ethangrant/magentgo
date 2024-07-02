package magentgo

import "context"

type AuthService struct {
	client *Client
}

type AuthPost struct {
	username string `json:"username"`
	password string `json:"password"`
	otp string `json:"otp"`
}

// auth endpoints respond with "string"
type AuthResponse string

func (a *AuthService) Admin (username string, password string, ctx context.Context) (AuthResponse, error) {
	var res AuthResponse = "";
	authPost := AuthPost{username: username, password: password}
	_, err := a.client.call("integration/customer/token", "POST", &authPost, &res, ctx)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (a *AuthService) Customer (username string, password string) {

}

func (a *AuthService) AdminWithGoogleAuthenticator (username string, password string, otp string) {
// TODO:
}

func (a *AuthService) AdminWithAuthy (username string, password string, otp string) {
// TODO:
}

func (a *AuthService) AdminWithU2fKey (username string, password string, otp string) {
	// TODO:
}