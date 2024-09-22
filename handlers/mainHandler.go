package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/rickalon/GoWebScraper/data"
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

}
