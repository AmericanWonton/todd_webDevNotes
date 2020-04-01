package main

import (
	"fmt"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

// For this code to run, you will need this package:
// go get github.com/satori/go.uuid

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	/* Request this cookie */
	cookie, err := req.Cookie("session")
	if err != nil {
		/* If there's an error, create a new cookie to work with */
		id, err2 := uuid.NewV4()
		if err2 != nil {
			log.Printf("Yo, there's an error making the uuid: %v\n", err2)
		}
		/* Above, we create a new uuid,(this method takes no parameters), to work with */
		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			/* Add in below if you are using https...this only makes the cookie if there's a secure connection */
			// Secure: true,
			HttpOnly: true,
			/* Above means you can't access this cookie with javascript, only https...it's kinda confusing */
			Path: "/",
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}
