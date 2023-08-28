package server

import (
	"net/http"
	"text/template"
)

import (
	"dummy_ai/pkg/env"
)

var (
	serverTemplate = template.Must(template.ParseFiles("./web/server/server.tmpl"))
)

func Run() {

	serveFile("/manifest.json" /*                    */, "./web/static/config/manifest.json")
	serveFile("/robots.txt" /*                       */, "./web/static/config/robots.txt")
	serveFile("/sitemap.xml" /*                      */, "./web/static/config/sitemap.xml")
	serveFile("/apple_touch_icon.png" /*             */, "./web/static/img/favicon/apple_touch_icon.png")
	serveFile("/favicon.ico" /*                      */, "./web/static/img/favicon/favicon.ico")
	serveFile("/favicon.svg" /*                      */, "./web/static/img/favicon/favicon.svg")
	serveFile("/favicon_192x192.png" /*              */, "./web/static/img/favicon/favicon_192x192.png")
	serveFile("/favicon_512x512.png" /*              */, "./web/static/img/favicon/favicon_512x512.png")
	serveFile("/menu_close_white.svg" /*             */, "./web/static/img/icon/menu_close_white.svg")
	serveFile("/menu_open_white.svg" /*              */, "./web/static/img/icon/menu_open_white.svg")
	serveFile("/logo.svg" /*                         */, "./web/static/img/logo/logo.svg")
	serveFile("/logo_white.svg" /*                   */, "./web/static/img/logo/logo_white.svg")
	serveFile("/RedHatDisplay-Bold.woff2" /*         */, "./web/static/lib/font/RedHatDisplay-Bold.woff2")
	serveFile("/RedHatDisplay-BoldItalic.woff2" /*   */, "./web/static/lib/font/RedHatDisplay-BoldItalic.woff2")
	serveFile("/RedHatDisplay-Medium.woff2" /*       */, "./web/static/lib/font/RedHatDisplay-Medium.woff2")
	serveFile("/RedHatDisplay-MediumItalic.woff2" /* */, "./web/static/lib/font/RedHatDisplay-MediumItalic.woff2")
	serveFile("/RedHatDisplayVF.woff2" /*            */, "./web/static/lib/font/RedHatDisplayVF.woff2")
	serveFile("/RedHatDisplayVF-Italic.woff2" /*     */, "./web/static/lib/font/RedHatDisplayVF-Italic.woff2")
	serveFile("/RedHatMono-Italic.woff2" /*          */, "./web/static/lib/font/RedHatMono-Italic.woff2")
	serveFile("/RedHatMono-Regular.woff2" /*         */, "./web/static/lib/font/RedHatMono-Regular.woff2")
	serveFile("/RedHatMonoVF.woff2" /*               */, "./web/static/lib/font/RedHatMonoVF.woff2")
	serveFile("/RedHatMonoVF-Italic.woff2" /*        */, "./web/static/lib/font/RedHatMonoVF-Italic.woff2")
	serveFile("/RedHatText-Italic" /*                */, "./web/static/lib/font/RedHatText-Italic.woff2")
	serveFile("/RedHatText-Medium.woff2" /*          */, "./web/static/lib/font/RedHatText-Medium.woff2")
	serveFile("/RedHatText-MediumItalic.woff2" /*    */, "./web/static/lib/font/RedHatText-MediumItalic.woff2")
	serveFile("/RedHatText-Regular.woff2" /*         */, "./web/static/lib/font/RedHatText-Regular.woff2")
	serveFile("/RedHatTextVF.woff2" /*               */, "./web/static/lib/font/RedHatTextVF.woff2")
	serveFile("/RedHatTextVF-Italic.woff2" /*        */, "./web/static/lib/font/RedHatTextVF-Italic.woff2")
	serveFile("/rhds.min.css" /*                     */, "./web/static/lib/rhds/rhds.min.css")
	serveFile("/rhds.min.js" /*                      */, "./web/static/lib/rhds/rhds.min.js")
	serveFile("/wasm_exec.js" /*                     */, "./web/static/lib/wasm/wasm_exec.js")
	serveFile("/wasm_run.js" /*                      */, "./web/static/lib/wasm/wasm_run.js")
	serveFile("/error_404.wasm" /*                   */, "./web/static/wasm/error_404.wasm")
	serveFile("/index.wasm" /*                       */, "./web/static/wasm/index.wasm")

	servePage("/" /* */, "/index.wasm")

	if err := http.ListenAndServe(env.ServerAddress, nil); err != nil {

		panic(err)
	}
}

func serveFile(route string, file string) {

	http.HandleFunc(route, func(responseWriter http.ResponseWriter, request *http.Request) {

		http.ServeFile(responseWriter, request, file)
	})
}

func servePage(route string, wasmRoute string) {

	http.HandleFunc(route, func(responseWriter http.ResponseWriter, request *http.Request) {

		if request.URL.Path != route {

			error404(responseWriter)
			return
		}

		executeServerTemplate(responseWriter, wasmRoute)
	})
}

func error404(responseWriter http.ResponseWriter) {

	responseWriter.WriteHeader(http.StatusNotFound)
	executeServerTemplate(responseWriter, "/error_404.wasm")
}

func executeServerTemplate(responseWriter http.ResponseWriter, wasmRoute string) {

	if err := serverTemplate.Execute(responseWriter, wasmRoute); err != nil {

		panic(err)
	}
}
