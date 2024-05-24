package config

import (
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func SetupServer() (*httprouter.Router, *http.Server) {
	router := httprouter.New()
	port := os.Getenv("PORT")
	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	return router, &server
}
