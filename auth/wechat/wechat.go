package wechat

import (
	"fmt"

	"github.com/medivhzhan/weapp/v2"
)

type service struct {
	secret string
	appID  string
}

func (s *service) ResolveCode(code string) (string, error) {
	res, err := weapp.Login(s.appID, s.secret, code)
	if err != nil {
		return "", fmt.Errorf("weapp.Login error: %v", err)
	}

	err = res.GetResponseError()
	if err != nil {
		return "", fmt.Errorf("weapp response error: %v", err)
	}

	return res.OpenID, nil
}

func NewWechatService(secret string, appID string) *service {
	return &service{
		secret: secret,
		appID:  appID,
	}
}
