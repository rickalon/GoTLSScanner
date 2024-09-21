package util

import "net/http"

func MsgToJson(w http.ResponseWriter, msg string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}
