package server

import (
	"net/http"
)

import (
	"dummy_ai/pkg/env"
)

func Run() {

	if err := http.ListenAndServe(env.ServerAddress, nil); err != nil {

		panic(err)
	}
}

func serveFile(route string, file string) {

	http.HandleFunc(route, func(responseWriter http.ResponseWriter, request *http.Request) {

		http.ServeFile(responseWriter, request, file)
	})
}
