package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type uberHandler int //Here is our base handler
type joe struct {
	Name         string
	Age          int
	Iscool       bool
	SenorRequest string
}

//Handler Declaration
func (namedHandler uberHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	/* Let's start by parsing all of our form data, (if there is any) */
	err1 := request.ParseForm()
	if err1 != nil {
		log.Fatalln(err1) //This takes an error...the Fatalln shuts down the program if an error occurs
	}

	//Declare the uberJoeStruct to be passed in data for our request!
	uberJoeStructExample := joe{
		Name:         "Joe",
		Age:          25,
		Iscool:       true,
		SenorRequest: request.FormValue("fname"),
	}
	fmt.Println(uberJoeStructExample.SenorRequest)
	/* Template execution, based upon a new Request coming in which is handled by our uberHanlder handler */
	template1.ExecuteTemplate(writer, "index.gohtml", uberJoeStructExample)
}

//Here is our main func
func main() {
	var aHandler uberHandler
	http.ListenAndServe(":8080", aHandler)
}

/* Template Declaration */
var template1 *template.Template

func init() {
	template1 = template.Must(template.ParseGlob("*.gohtml")) //Parse all files with gohtml
}

/* Used to check if form field is empty */
func isEmpty(j joe) bool {
	theName := j.SenorRequest
	judgement := false
	if theName == "" {
		judgement = false
	} else {
		judgement = true
	}

	return judgement
}
