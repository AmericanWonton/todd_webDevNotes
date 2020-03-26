package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("insult")
	/* Let's write some logic based on what you enter. It's kind of funky, but it
	works! */
	isSpaghetti(w, req)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="GET">
	 <input type="text" name="insult">
	 <input type="submit">
	</form>
	<br>`+v)
}

/* This is kind of funky, but it works! */
func isSpaghetti(w http.ResponseWriter, req *http.Request) {
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		log.Fatal(err)
	}
	dumpString := (string(requestDump))

	i := 0 //This is just a flag signaler in the program
	scanner := bufio.NewScanner(strings.NewReader(dumpString))
	/* This  is going to scan the entirety of the connection. When we get to the first line of that scan,
	it should be the request line first, (i == 0, part a).*/
	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			// request line (Part a)
			/* strings.Fields is a string function that splits a series of string characters up by the white space.
			So what Todd is doing here is just getting 'GET', cuz it's spot 0 in that slice of byte sections */
			/* This is for the exercise Todd asked us to do, getting the URL */
			thequery := strings.Fields(ln)[1]
			fmt.Println(thequery)
			/* Check to see if the entry is larger than 9 so we can format the slice */
			if len(thequery) >= 9 {
				substring := thequery[9:len(thequery)]
				if strings.Compare(substring, "Spaghetti") == 0 {
					fmt.Printf("Our query we entered is: %s\n", substring)
					fmt.Printf("Hey, you're  a cool dude\n")
				} else {
					fmt.Printf("Where is my spaghetti?\n")
					fmt.Printf("All I got was %s\n", substring)
				}
			}
		}
		if ln == "" {
			/* there's the request line, then the headers...and then a blank line, meaning we are done. */
			break
		}
		i++
	}
}
