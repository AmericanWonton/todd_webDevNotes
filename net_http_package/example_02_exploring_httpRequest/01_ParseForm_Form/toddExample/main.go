package main

import (
	"html/template"
	"log"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	/* ParseForm() is one of those fields inside of the Request type. Look at GoDoc for more, but it basically
	handles FormData passed in through a Request */
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err) //This takes an error...the Fatalln shuts down the program if an error occurs
	}
	/* Whenever this Request hits our server, it asks for the form values,(if any are submitted. Otherwise, we just have
	that basic form you see when first starting this.). We pass the writer to our executed template, along with the
	req.Form,(which has all the form data from our Request). The fields in that Request are what we are passing! (the fnames)*/
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
