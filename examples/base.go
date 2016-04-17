package main

import (
	"log"
	"net/http"

	"github.com/maciekmm/messenger-platform-go-sdk"
)

var mess = &messenger.Messenger{
	AccessToken: "ACCESS_TOKEN",
}

func main() {
	http.HandleFunc("/webhook", mess.Handler)
	log.Fatal(http.ListenAndServe(":5646", nil))
}
