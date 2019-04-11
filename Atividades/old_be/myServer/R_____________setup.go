package myServer

import (
	"github.com/gorilla/handlers"
	"lfkk/be/myRouter"
	"log"
	"net/http"
)

func StartServer() {
	r := myRouter.NewRouter()
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	log.Fatal(http.ListenAndServe(SERVER_HOST,
		handlers.CORS(allowedOrigins, allowedMethods)(r)))
}
