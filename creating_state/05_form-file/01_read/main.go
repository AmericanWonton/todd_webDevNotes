package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	/* This below has the 'not found' handler, because we're too lazy to search for it and include it so we handle that */
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

	var fileString string //This is the result of our file upload, we will add it to our form result later.
	fmt.Println(req.Method)
	/* In the if statement, these are constants, per the documenation of http and req. If the request is the 'POST' constant */
	if req.Method == http.MethodPost {

		// open
		/* This is for catching a file coming in by a User */
		f, h, err := req.FormFile("the_form")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		// for your information
		fmt.Println("\nfile:", f, "\nheader:", h, "\nerr", err)

		// read
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fileString = string(bs)
	}
	/* Even if this ISN'T POST, this still runs, setting the response header a text to HTML Form.
	Additionally, if the contents of the file are uploaded */
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="the_form">
	<input type="submit">
	</form>
	<br>`+fileString)
	/* enctype is for uploading a file */
	/* We also don't need to specify the action in the form, since we just want to reload the page we were on */
}

/* May want to try this with a simple txt file for better looking results. */
