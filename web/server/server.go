package server

import (
	"net/http"
)

import (
	"dummy_ai/pkg/env"
)

func Run() {

	serveFile("/apple_touch_icon.png" /**/, "./web/static/image/favicon/apple_touch_icon.png")
	serveFile("/favicon.ico" /*         */, "./web/static/image/favicon/favicon.ico")

	if err := http.ListenAndServe(env.ServerAddress, nil); err != nil {

		panic(err)
	}
}

func serveFile(route string, file string) {

	http.HandleFunc(route, func(responseWriter http.ResponseWriter, request *http.Request) {

		http.ServeFile(responseWriter, request, file)
	})
}
