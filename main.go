package main

import (
	"log"
	"net/http"

	"github.com/trickierstinky/slack-invite-api/routes"
)

func main() {
	router := routes.Router()

	log.Fatal(http.ListenAndServe(":8080", router))

}
