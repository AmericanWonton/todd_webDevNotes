package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog doggy")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat catty")
}

func main() {

	http.HandleFunc("/dog", d)
	http.HandleFunc("/cat", c)

	http.ListenAndServe(":8080", nil) //When we pass in nil, we are using the default serveMux
}

/* This is the best way to do this. */
