package goapp

import (
	"appengine"
	"github.com/mjibson/appstats" // HL
	"net/http"
)

func init() {
	http.Handle("/", appstats.NewHandler(Main)) // HL
}

func Main(c appengine.Context, w http.ResponseWriter, r *http.Request) { // HL
	w.Write([]byte("hello from go on app engine"))
}
