func Index(c mpg.Context, w http.ResponseWriter, r *http.Request) {
	someFunc()
	otherFunc()
	a := another()
	c.Step("a.list loop", func() { // HL
		for _, v := range a.list {
			v.maybeSlowThing()
		}
	}) // HL
	fmt.Fprintf(w, "<html><body>%v</body></html>", c.Includes())
}
