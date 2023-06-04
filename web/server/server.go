package server

import (
	"net/http"
)

import (
	"dummy_ai/pkg/env"
)

func Run() {

	serveFile("/manifest.json" /*       */, "./web/static/config/manifest.json")
	serveFile("/robots.txt" /*          */, "./web/static/config/robots.txt")
	serveFile("/sitemap.xml" /*         */, "./web/static/config/sitemap.xml")
	serveFile("/apple_touch_icon.png" /**/, "./web/static/img/favicon/apple_touch_icon.png")
	serveFile("/favicon.ico" /*         */, "./web/static/img/favicon/favicon.ico")
	serveFile("/favicon.svg" /*         */, "./web/static/img/favicon/favicon.svg")
	serveFile("/favicon_192x192.png" /* */, "./web/static/img/favicon/favicon_192x192.png")
	serveFile("/favicon_512x512.png" /* */, "./web/static/img/favicon/favicon_512x512.png")
	serveFile("/logo.svg" /*            */, "./web/static/img/logo/logo.svg")
	serveFile("/logo_white.svg" /*      */, "./web/static/img/logo/logo_white.svg")
	serveFile("/index.wasm" /*          */, "./web/static/wasm/index.wasm")

	if err := http.ListenAndServe(env.ServerAddress, nil); err != nil {

		panic(err)
	}
}

func serveFile(route string, file string) {

	http.HandleFunc(route, func(responseWriter http.ResponseWriter, request *http.Request) {

		http.ServeFile(responseWriter, request, file)
	})
}
