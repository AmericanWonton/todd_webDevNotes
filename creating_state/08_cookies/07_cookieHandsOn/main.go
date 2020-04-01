package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("my-cookie")

	/* This is what is returned when no cookie is found. If it has that error, then let's make a cookie with a value of 0 */
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
			Path: "/",
		}
	}
	//Convert to a string to int
	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count++
	cookie.Value = strconv.Itoa(count)
	//Set the value to the new cookie
	http.SetCookie(res, cookie)

	io.WriteString(res, cookie.Value)
}
