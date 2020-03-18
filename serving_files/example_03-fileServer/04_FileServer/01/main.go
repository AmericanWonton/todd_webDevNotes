package main

import (
	"io"
	"net/http"
)

func main() {
	/* This is saying 'serve everything in this current directory */
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/dog/", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	/* Set the response header below */
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	/* write back this html. If the fileServer has that file at this directory, (/), it will serve it to us */
	io.WriteString(w, `<img src="/toby.jpg">`)
}

/* So if we go to localhost:8080/dog/, we get the dog. But if we just do localhost:8080, we get served ALL the files! */
