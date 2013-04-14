package main

import "fmt"
import "github.com/mjibson/MiniProfiler/go/miniprofiler"
import "net/http"

func main() {
	profiles := make(map[string]*miniprofiler.Profile) // HL
	miniprofiler.Enable = func(r *http.Request) bool { return true } // HL
	miniprofiler.Store = func(r *http.Request, p *miniprofiler.Profile) { // HL
		profiles[string(p.Id)] = p // HL
	} // HL
	miniprofiler.Get = func(r *http.Request, id string) *miniprofiler.Profile { // HL
		return profiles[id] // HL
	} // HL
	http.Handle("/", miniprofiler.NewHandler(Index))
	http.ListenAndServe(":8888", nil)
}

func Index(p *miniprofiler.Profile, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body>%v</body></html>", p.Includes())
}
