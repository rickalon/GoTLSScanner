package settings

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var OauthConfig oauth2.Config

var OauthStateString = "random-string" //static for reference, should be diferent for each request

func NewOauth2Config() {
	OauthConfig = oauth2.Config{
		ClientID:     getenv("OAUTH2_ID", ""),
		ClientSecret: getenv("OAUTH2_SECRET", ""),
		Scopes:       []string{"read:user"},
		Endpoint:     github.Endpoint,
		RedirectURL:  "http://localhost:8080/oauth2/callback",
	}
}

func Authentication(accessToken string) error {
	url := fmt.Sprintf("https://api.github.com/applications/%s/token", OauthConfig.ClientID)

	data := map[string]string{
		"access_token": accessToken,
	}
	jsonData, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.SetBasicAuth(OauthConfig.ClientID, OauthConfig.ClientSecret)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return nil
	} else {
		return errors.New("Token not valid")
	}
}
