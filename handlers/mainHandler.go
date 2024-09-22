package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/rickalon/GoWebScraper/data"
	"github.com/rickalon/GoWebScraper/services"
	"github.com/rickalon/GoWebScraper/util"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Mensaje de prueba"
	util.MsgToJson(w, msg)
}

func URLHandler(w http.ResponseWriter, r *http.Request) {

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

	ch := make(chan *data.UrlObj)
	//fan in-fan out
	go services.UrlProc(ctx, url, ch)
	//orDone pattern
	arrUrl := []*data.UrlObj{}
	for val := range services.OrDone(ctx, ch) {
		arrUrl = append(arrUrl, val)
	}
	//return them
	util.WriteURLtoJson(w, arrUrl)
}
