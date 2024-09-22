package handlers

import (
	"context"
	"encoding/json"
	"fmt"
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

	ch := make(chan string)
	//fan in-fan out
	for _, val := range url.Data {
		fmt.Println("Procesamos 1 url", val)
	}
	//orDone pattern
	for val := range services.OrDone(ctx, ch) {
		fmt.Println(val)
	}

}
