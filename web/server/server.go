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
	serveFile("/favicon.svg" /*         */, "./web/static/image/favicon/favicon.svg")
	serveFile("/favicon_192x192.png" /* */, "./web/static/image/favicon/favicon_192x192.png")
	serveFile("/favicon_512x512.png" /* */, "./web/static/image/favicon/favicon_512x512.png")
	serveFile("/logo.svg" /*            */, "./web/static/image/logo/logo.svg")
	serveFile("/logo_white.svg" /*      */, "./web/static/image/logo/logo_white.svg")

	if err := http.ListenAndServe(env.ServerAddress, nil); err != nil {

		panic(err)
	}
}

func serveFile(route string, file string) {

	http.HandleFunc(route, func(responseWriter http.ResponseWriter, request *http.Request) {

		http.ServeFile(responseWriter, request, file)
	})
}
