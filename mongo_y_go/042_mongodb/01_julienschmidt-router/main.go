package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	/* below means that any GET paramater that comes in at index is handledd with index */
	r.GET("/", index)
	http.ListenAndServe("localhost:8080", r)
}

// note: using 'r' instead of 'req'
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

/* Enter this in Terminal to see what's up with curl:
curl http://localhost:8080/ */
