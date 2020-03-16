package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "doggy doggy doggy")
	case "/cat":
		io.WriteString(w, "kitty kitty kitty")
	}
}

func main() {
	// d is type handler, type hotdog, AND type int!
	var d hotdog
	http.ListenAndServe(":8080", d) //8080 is our dev route for testing
}

/* This is NOT the best way to write this code! THis is just an example */
