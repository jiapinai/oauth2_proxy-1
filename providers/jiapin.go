package providers

import (
	"log"
	"net/http"
	"net/url"

	"github.com/bitly/oauth2_proxy/api"
)

type JiapinProvider struct {
	*ProviderData
}

func NewJiapinProvider(p *ProviderData) *JiapinProvider {
	p.ProviderName = "Jiapin"
	if p.LoginURL == nil || p.LoginURL.String() == "" {
		p.LoginURL = &url.URL{
			Scheme: "https",
			Host:   "usercenter.jiapinai.com",
			Path:   "/oauth/auth/authorize",
		}
	}
	if p.RedeemURL == nil || p.RedeemURL.String() == "" {
		p.RedeemURL = &url.URL{
			Scheme: "https",
			Host:   "usercenter.jiapinai.com",
			Path:   "/oauth/auth/token",
		}
	}
	if p.ValidateURL == nil || p.ValidateURL.String() == "" {
		p.ValidateURL = &url.URL{
			Scheme: "https",
			Host:   "usercenter.jiapinai.com",
			Path:   "/oauth/auth/resource",
		}
	}
	if p.Scope == "" {
		p.Scope = "admin"
	}
	return &JiapinProvider{ProviderData: p}
}

func (p *JiapinProvider) GetEmailAddress(s *sessions.SessionState) (string, error) {

	req, err := http.NewRequest("GET",
		p.ValidateURL.String()+"?access_token="+s.AccessToken, nil)
	if err != nil {
		log.Printf("failed building request %s", err)
		return "", err
	}
	json, err := api.Request(req)
	if err != nil {
		log.Printf("failed making request %s", err)
		return "", err
	}
	return json.Get("user").Get("username").String()
}
