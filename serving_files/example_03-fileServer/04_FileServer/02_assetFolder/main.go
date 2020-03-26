package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	/* The second and third arguments are BOTH handlers
	Anything that comes in at the /resources/ url is going to be served everything under assets folder.
	Strip prefix is stripping off the /resources section and now searches for anything passed as an argument.
	Our dog, for example. */
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	/* We're asking for the toby image from the handled assets folder */
	io.WriteString(w, `<img src="/resources/toby.jpg">`)
}

/*

./assets/toby.jpg

*/
