package handlers

import (
	"net/http"

	"github.com/rickalon/GoWebScraper/util"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Mensaje de prueba"
	util.MsgToJson(w, msg)
}
