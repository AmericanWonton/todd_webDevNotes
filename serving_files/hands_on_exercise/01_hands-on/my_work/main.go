package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func dog(res http.ResponseWriter, req *http.Request) {
	//Making template
	dogTemplate, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	//Execute Template
	dogTemplate.ExecuteTemplate(res, "dog.gohtml", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo ran")
}

func chien(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "dog.jpg")
}

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/dog.jpg", chien)
	http.ListenAndServe(":8080", nil) //When we pass in nil, we are using the default serveMux
}

//template
