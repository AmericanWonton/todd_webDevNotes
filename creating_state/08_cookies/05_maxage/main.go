package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	//Click to set the cookie
	fmt.Fprintln(w, `<h1><a href="/set">set a cookie</a></h1>`)
}

func set(w http.ResponseWriter, req *http.Request) {
	/* Once clicked,  Se the cookie and write the read button to the request.*/
	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: "some value",
		Path:  "/",
	})
	fmt.Fprintln(w, `<h1><a href="/read">read</a></h1>`)
}

func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	/* If it errors, redirect back to set */
	if err != nil {
		http.Redirect(w, req, "/set", http.StatusSeeOther)
		return
	}

	fmt.Fprintf(w, `<h1>Your Cookie:<br>%v</h1><h1><a href="/expire">expire</a></h1>`, c)
}

func expire(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	/* If we look for the 'session' cookie and don't find it, redirect back to set */
	if err != nil {
		http.Redirect(w, req, "/set", http.StatusSeeOther)
		return
		/* Must return out or this code will run sequentillay! */
	}
	c.MaxAge = -1 // delete cookie
	http.SetCookie(w, c)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}
