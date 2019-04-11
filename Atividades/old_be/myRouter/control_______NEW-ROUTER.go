package myRouter

import (
	"github.com/gorilla/mux"
	"lfkk/be/myLogger"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range allRoutes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = myLogger.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}
