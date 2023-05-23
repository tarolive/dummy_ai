package server

import (
	"net/http"
)

import (
	"dummy_ai/pkg/env"
)

func Run() {

	err := http.ListenAndServe(env.ServerAddress, nil)

	if err != nil {
		panic(err)
	}
}
