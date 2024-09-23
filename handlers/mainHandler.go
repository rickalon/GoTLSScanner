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
	msg := "test msg"
	fmt.Println(msg)
}

func Oauth2Login(w http.ResponseWriter, r *http.Request) {

	url := settings.OauthConfig.AuthCodeURL(settings.OauthStateString)
	fmt.Println(settings.OauthConfig)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func URLDBHandler(persistance db.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()

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
