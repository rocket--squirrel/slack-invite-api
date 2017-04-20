package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/trickierstinky/slack-invite-api/config"
	"github.com/trickierstinky/slack-invite-api/data"
	"github.com/trickierstinky/slack-invite-api/routes"
)

func main() {
	router := routes.Router()
	port := fmt.Sprintf(":%s", config.Env("port"))

	data.SetupDatabase()

	fmt.Printf("Started Listenting on PORT=%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
