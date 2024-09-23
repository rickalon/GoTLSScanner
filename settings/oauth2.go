package settings

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var OauthConfig oauth2.Config

var OauthStateString = "random-string"

func NewOauth2Config() {
	OauthConfig = oauth2.Config{
		ClientID:     getenv("OAUTH2_ID", ""),
		ClientSecret: getenv("OAUTH2_SECRET", ""),
		Scopes:       []string{"read:user"},
		Endpoint:     github.Endpoint,
		RedirectURL:  "http://localhost:8080/oauth2/callback",
	}
}
