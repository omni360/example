package main

import (
	//	"fmt"
	"net/http"
)

/*
type Hello struct{}

//htmstr:=
func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Print(w, "ddd")
}
func main() {
	var h Hello
	http.ListenAndServe(
		"localhost:4000", h)
}
*/
func main() {
	http.Handle("/", http.FileServer(http.Dir("web/")))
	http.ListenAndServe("localhost:4000", nil)
}
