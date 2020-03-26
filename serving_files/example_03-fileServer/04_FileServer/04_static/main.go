package main

import (
	"log"
	"net/http"
)

func main() {
	/* This is done in log.Fatal, (it's something different if you watch 46 and 47).
	Listen and serve serves an error. We wanna catch that error and log it. */
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}

/* This is for making a fileserver, (just a static website, not doing anything special) */
/* You need an index.html. Also, it will no longer show your files in the root */
/* Here we are just serving all of our files in the directory. */
/* Try changing index.html to inde.html to see what happens */
