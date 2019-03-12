package util

import (
	"github.com/solarhell/antsdk/api"
)

type AlipaySystemOauthTokenResponse struct {
	api.AlipayResponse
	AlipayUserId string `json:"alipay_user_id"`
	UserId       string `json:"user_id"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	ReExpiresIn  int    `json:"re_expires_in"`
}
