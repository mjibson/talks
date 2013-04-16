package goapp

import (
	"fmt"
	mpg "github.com/mjibson/MiniProfiler/go/miniprofiler_gae" // HL
	"net/http"
)

func init() {
	http.Handle("/", miniprofiler.NewHandler(Index)) // HL
}

func Index(c mpg.Context, w http.ResponseWriter, r *http.Request) { // HL
	fmt.Fprintf(w, "<html><body>%v</body></html>", c.Includes()) // HL
}
