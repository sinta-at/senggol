package httpserver

import (
	"log"
	"net/http"
	"senggol/service"
)

func StartServer(port string, services service.Services) {
	router := makeRouter(services)

	log.Fatal(http.ListenAndServe(":" + port, router))
}