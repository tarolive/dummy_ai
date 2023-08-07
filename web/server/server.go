package server

import (
	"net/http"
	"text/template"
)

import (
	"dummy_ai/pkg/env"
)

var (
	_serverTemplate = template.Must(template.ParseFiles("./web/server/server.tmpl"))
)

func Run() {

	_serveFile("/manifest.json" /*               */, "./web/static/config/manifest.json")
	_serveFile("/robots.txt" /*                  */, "./web/static/config/robots.txt")
	_serveFile("/sitemap.xml" /*                 */, "./web/static/config/sitemap.xml")
	_serveFile("/apple_touch_icon.png" /*        */, "./web/static/img/favicon/apple_touch_icon.png")
	_serveFile("/favicon.ico" /*                 */, "./web/static/img/favicon/favicon.ico")
	_serveFile("/favicon.svg" /*                 */, "./web/static/img/favicon/favicon.svg")
	_serveFile("/favicon_192x192.png" /*         */, "./web/static/img/favicon/favicon_192x192.png")
	_serveFile("/favicon_512x512.png" /*         */, "./web/static/img/favicon/favicon_512x512.png")
	_serveFile("/logo.svg" /*                    */, "./web/static/img/logo/logo.svg")
	_serveFile("/logo_white.svg" /*              */, "./web/static/img/logo/logo_white.svg")
	_serveFile("/rhds.min.css" /*                */, "./web/static/lib/rhds/rhds.min.css")
	_serveFile("/rhds.min.js" /*                 */, "./web/static/lib/rhds/rhds.min.js")
	_serveFile("/RedHatDisplay-Bold.woff2" /*    */, "./web/static/lib/rhds/font/RedHatDisplay-Bold.woff2")
	_serveFile("/RedHatDisplay-Medium.woff2" /*  */, "./web/static/lib/rhds/font/RedHatDisplay-Medium.woff2")
	_serveFile("/RedHatDisplay-Regular.woff2" /* */, "./web/static/lib/rhds/font/RedHatDisplay-Regular.woff2")
	_serveFile("/RedHatText-Medium.woff2" /*     */, "./web/static/lib/rhds/font/RedHatText-Medium.woff2")
	_serveFile("/RedHatText-Regular.woff2" /*    */, "./web/static/lib/rhds/font/RedHatText-Regular.woff2")
	_serveFile("/wasm_exec.js" /*                */, "./web/static/lib/wasm/wasm_exec.js")
	_serveFile("/wasm_run.js" /*                 */, "./web/static/lib/wasm/wasm_run.js")
	_serveFile("/error_404.wasm" /*              */, "./web/static/wasm/error_404.wasm")
	_serveFile("/index.wasm" /*                  */, "./web/static/wasm/index.wasm")

	_servePage("/" /* */, "/index.wasm")

	if err := http.ListenAndServe(env.ServerAddress, nil); err != nil {

		panic(err)
	}
}

func _serveFile(route string, file string) {

	http.HandleFunc(route, func(responseWriter http.ResponseWriter, request *http.Request) {

		http.ServeFile(responseWriter, request, file)
	})
}

func _servePage(route string, wasmRoute string) {

	http.HandleFunc(route, func(responseWriter http.ResponseWriter, request *http.Request) {

		if request.URL.Path != route {

			_error404(responseWriter)
			return
		}

		_executeServerTemplate(responseWriter, wasmRoute)
	})
}

func _error404(responseWriter http.ResponseWriter) {

	responseWriter.WriteHeader(http.StatusNotFound)
	_executeServerTemplate(responseWriter, "/error_404.wasm")
}

func _executeServerTemplate(responseWriter http.ResponseWriter, wasmRoute string) {

	if err := _serverTemplate.Execute(responseWriter, wasmRoute); err != nil {

		panic(err)
	}
}
