func Index(c mpg.Context, w http.ResponseWriter, r *http.Request) {
	someFunc()
	otherFunc()
	a := another()
	for _, v := range a.list {
		v.maybeSlowThing()
	}
	fmt.Fprintf(w, "<html><body>%v</body></html>", c.Includes())
}
