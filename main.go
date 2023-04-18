package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/carregionbot/webhook", func(w http.ResponseWriter, r *http.Request) {
		HandleTelegramWebHook(w, r);
	})

	http.ListenAndServe(":80", r)
}