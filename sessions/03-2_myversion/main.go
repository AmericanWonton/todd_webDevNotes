package main

import (
	"html/template"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type warCriminal struct {
	FirstName string
	Lastname  string
	Age       string
	Password  string
}

var uberTemplate *template.Template //template definition
/* These are the Users and their info we store in the database; they have a unique User ID */
var dbUsers = map[string]warCriminal{} // user ID, warCriminal struct
/* These are the session ID of the Users */
var dbSessions = map[string]string{} // session ID, user ID

func init() {
	uberTemplate = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/gulag", gulag)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	/* When we enter our main page, let's see if someone is already signed in */
	user := getUser(req)
	uberTemplate.ExecuteTemplate(w, "index.gohtml", user)
}

func gulag(w http.ResponseWriter, req *http.Request) {
	/* get the User */
	newUser := getUser(req)
	/* If the User isn't logged in, bring them to the main screen! */
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	uberTemplate.ExecuteTemplate(w, "gulag.gohtml", newUser)
}

func signup(w http.ResponseWriter, req *http.Request) {
	/* If alreadyLoggedIn, (which examines our cookies to see if the User value is already in the cookies),
	returns true, we redirect to the 'logged-in' page */
	if alreadyLoggedIn(req) {
		log.Printf("Already logged in with user\n. Going to Log in screen.\n")
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
		/* After redirecting, we have to return out of this or we loop through the
		below code! */
	}

	log.Printf("We aren't logged in, let's see what we can do...\n")
	/* If you ain't logged in, we go through this... */
	// process form submission
	if req.Method == http.MethodPost {

		// get form values
		firstName := req.FormValue("firstname")
		lastName := req.FormValue("lastname")
		age := req.FormValue("age")
		password := req.FormValue("password")

		/* This first name is already input and someone created this User before!
		We could do the same for password, but there's already enough bad dictators in the world...
		so we'll say, no, go get creative. If we search with the first name and it does NOT
		turn up an empty dbUsers,(AKA, the name exist and so does a profile), we redirect */
		if dbUsers[firstName] != (warCriminal{}) {
			http.Error(w, "This dictator's first name already exists!", http.StatusForbidden)
			log.Printf("Username already taken!\n")
			return
		}
		// username taken?
		/*
			if _, ok := dbUsers[un]; ok {
				http.Error(w, "Username already taken", http.StatusForbidden)
				log.Printf("Username already taken!\n")
				return
			}
		*/
		/* If the Username is NOT already taken, we can create a new session! */
		// create session
		sessionID, _ := uuid.NewV4()
		cookie := &http.Cookie{
			Name:  "session",
			Value: sessionID.String(),
		}
		http.SetCookie(w, cookie) //Set the new cookie with the http package, with the new session value string
		dbSessions[cookie.Value] = firstName

		// store user in dbUsers
		newUser := warCriminal{firstName, lastName, age, password}
		dbUsers[firstName] = newUser
		/* Mkay, we've created the new User, they are stored in our 'database' and we can go
		return to show the User they're logged onto the main page! */
		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	uberTemplate.ExecuteTemplate(w, "signup.gothml", nil)
}
