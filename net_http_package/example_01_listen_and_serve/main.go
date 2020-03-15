package main

import (
	"fmt"
	"net/http"
)

type example int //The underlying type of 'example' is an int...BUT
//Here we attatch the ServeHTTP method to 'example'. Now 'example', or any value of hotdog is ALSO
//of type Handler.
func (ex example) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("Any code you want in this func")
}

func main() {
	var exampleVariable example //This is our example 'type' variable, that has the Handler interface added now.
	http.ListenAndServe(":8080", exampleVariable)
	/* What the above is doing: Listening and accepting anything from localhost 8080. Then, it passes back something to serve,
	at whatever is passed in, (in our case, our example handler we made) */
}

/* Just a heads up running this code, chrome launches multiple requests...we don't know why, it's weird for example's sake.
Todd uses Firefox
*/

/* This is kind of messy code...there's a better way of doing this...see the next videos/files. */
