package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("carregionbot/webhook", func(w http.ResponseWriter, r *http.Request) {
		HandleTelegramWebHook(w, r);
		fmt.Fprintf(w, "<h1>Telegram webhook!\n</h1>")
	})

	http.ListenAndServe(":80", r)
}