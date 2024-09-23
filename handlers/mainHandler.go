package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rickalon/GoWebScraper/data"
	"github.com/rickalon/GoWebScraper/db"
	"github.com/rickalon/GoWebScraper/services"
	"github.com/rickalon/GoWebScraper/settings"
	"github.com/rickalon/GoWebScraper/util"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	util.MsgToJson(w, "Test endpoint")
}

func Callback(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	state := r.FormValue("state")
	if state != settings.OauthStateString {
		str := fmt.Sprintf("Invalid OAuth state, expected '%s', got '%s'\n", settings.OauthStateString, state)
		util.ErrorStringToJson(w, http.StatusBadRequest, str)
		return
	}

	code := r.FormValue("code")
	token, err := settings.OauthConfig.Exchange(r.Context(), code)
	if err != nil {
		str := fmt.Sprintf("oauthConfig.Exchange() failed: %s\n", err)
		util.ErrorStringToJson(w, http.StatusBadRequest, str)
		return
	}

	client := settings.OauthConfig.Client(r.Context(), token)
	userInfo, err := client.Get("https://api.github.com/user")
	if err != nil {
		str := fmt.Sprintf("client.Get() failed: %s\n", err)
		util.ErrorStringToJson(w, http.StatusBadRequest, str)
		return
	}
	defer userInfo.Body.Close()
	str := fmt.Sprintf("ACCESS TOKEN:%v", token.AccessToken)
	util.MsgToJson(w, str)

}

func Oauth2Login(w http.ResponseWriter, r *http.Request) {
	url := settings.OauthConfig.AuthCodeURL(settings.OauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func URLDBHandler(persistance db.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		strH := r.Header.Get("Authorization")
		strP := r.FormValue("token")
		if strH != "" {
			err := settings.Authentication(strH)

			if err != nil {
				util.ErrorStringToJson(w, 400, "Token not valid")
				return
			}
		} else if strP != "" {
			err := settings.Authentication(strH)
			if err != nil {
				util.ErrorStringToJson(w, 400, "Token not valid")
				return
			}
		} else {
			util.ErrorStringToJson(w, 400, "Token not valid")
			return
		}
		body, err := io.ReadAll(r.Body)

		if err != nil {
			util.ErrorToJson(w, 500, err)
			return
		}
		url := data.NewURL()
		err = json.Unmarshal(body, url)
		if err != nil {
			util.ErrorToJson(w, 500, err)
			return
		}
		ctx, cancel := context.WithCancel(r.Context())
		defer cancel()

		ch := make(chan *data.URL)
		//fan in-fan out
		go services.UrlProc(ctx, url, ch, persistance)
		//orDone pattern
		arrUrl := []*data.URL{}

		for val := range services.OrDone(ctx, ch) {
			arrUrl = append(arrUrl, val)
		}
		//return them
		util.WriteURLtoJson(w, arrUrl)
	}
}
