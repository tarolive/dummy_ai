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

	serveFile("/manifest.json" /*                       */, "./web/static/config/manifest.json")
	serveFile("/robots.txt" /*                          */, "./web/static/config/robots.txt")
	serveFile("/sitemap.xml" /*                         */, "./web/static/config/sitemap.xml")
	serveFile("/apple_touch_icon.png" /*                */, "./web/static/img/favicon/apple_touch_icon.png")
	serveFile("/favicon.ico" /*                         */, "./web/static/img/favicon/favicon.ico")
	serveFile("/favicon.svg" /*                         */, "./web/static/img/favicon/favicon.svg")
	serveFile("/favicon_192x192.png" /*                 */, "./web/static/img/favicon/favicon_192x192.png")
	serveFile("/favicon_512x512.png" /*                 */, "./web/static/img/favicon/favicon_512x512.png")
	serveFile("/logo.svg" /*                            */, "./web/static/img/logo/logo.svg")
	serveFile("/logo_white.svg" /*                      */, "./web/static/img/logo/logo_white.svg")
	serveFile("/font_awesome.woff2" /*                  */, "./web/static/lib/font/font_awesome.woff2")
	serveFile("/red_hat_display_bold.woff2" /*          */, "./web/static/lib/font/red_hat_display_bold.woff2")
	serveFile("/red_hat_display_bold_italic.woff2" /*   */, "./web/static/lib/font/red_hat_display_bold_italic.woff2")
	serveFile("/red_hat_display_medium.woff2" /*        */, "./web/static/lib/font/red_hat_display_medium.woff2")
	serveFile("/red_hat_display_medium_italic.woff2" /* */, "./web/static/lib/font/red_hat_display_medium_italic.woff2")
	serveFile("/red_hat_display_vf.woff2" /*            */, "./web/static/lib/font/red_hat_display_vf.woff2")
	serveFile("/red_hat_display_vf_italic.woff2" /*     */, "./web/static/lib/font/red_hat_display_vf_italic.woff2")
	serveFile("/red_hat_mono_italic.woff2" /*           */, "./web/static/lib/font/red_hat_mono_italic.woff2")
	serveFile("/red_hat_mono_regular.woff2" /*          */, "./web/static/lib/font/red_hat_mono_regular.woff2")
	serveFile("/red_hat_mono_vf.woff2" /*               */, "./web/static/lib/font/red_hat_mono_vf.woff2")
	serveFile("/red_hat_mono_vf_italic.woff2" /*        */, "./web/static/lib/font/red_hat_mono_vf_italic.woff2")
	serveFile("/red_hat_text_italic.woff2" /*           */, "./web/static/lib/font/red_hat_text_italic.woff2")
	serveFile("/red_hat_text_medium.woff2" /*           */, "./web/static/lib/font/red_hat_text_medium.woff2")
	serveFile("/red_hat_text_medium_italic.woff2" /*    */, "./web/static/lib/font/red_hat_text_medium_italic.woff2")
	serveFile("/red_hat_text_regular.woff2" /*          */, "./web/static/lib/font/red_hat_text_regular.woff2")
	serveFile("/red_hat_text_vf.woff2" /*               */, "./web/static/lib/font/red_hat_text_vf.woff2")
	serveFile("/red_hat_text_vf_italic.woff2" /*        */, "./web/static/lib/font/red_hat_text_vf_italic.woff2")
	serveFile("/patternfly.css" /*                      */, "./web/static/lib/patternfly/patternfly.css")
	serveFile("/wasm_exec.js" /*                        */, "./web/static/lib/wasm/wasm_exec.js")
	serveFile("/wasm_run.js" /*                         */, "./web/static/lib/wasm/wasm_run.js")
	serveFile("/error_404.wasm" /*                      */, "./web/static/wasm/error_404.wasm")
	serveFile("/index.wasm" /*                          */, "./web/static/wasm/index.wasm")

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

func executeServerTemplate(responseWriter http.ResponseWriter, wasmRoute string) {

	if err := serverTemplate.Execute(responseWriter, wasmRoute); err != nil {

		panic(err)
	}
}

func error404(responseWriter http.ResponseWriter) {

	responseWriter.WriteHeader(http.StatusNotFound)
	executeServerTemplate(responseWriter, "/error_404.wasm")
}
