package main

import (
	"io"
	"net/http"
)

func main() {
	/* Here we handle anything with the URL of /toby.jpg, (whihc is served in our local file with main.go) */
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	/* Here we just write to the client what the html code should display */
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="toby.jpg">	`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "toby.jpg")
}
