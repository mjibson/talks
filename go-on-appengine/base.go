package goapp

import (
	"appengine"
	"net/http"
)

func init() {
	http.HandleFunc("/", Main)
}

func Main(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from go on app engine"))
}
