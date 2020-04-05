package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

/* this has the username and the time of the last activity */
type session struct {
	Username     string
	lastActivity time.Time
}

var theTemplate *template.Template
var dbUsers = map[string]user{}       // user ID, user
var dbSessions = map[string]session{} // session ID, session structs we now have
var dbSessionsCleaned time.Time       //This keeps track of session times

const sessionLength int = 10

/* Role declaration */
const theUser string = "USER"
const theAdmin string = "ADMIN"
const theEmployee string = "EMPLOYEE"

/* Funciton parse declaration */
var funcMap = template.FuncMap{
	"upperCase": strings.ToUpper, //upperCase is a key we can call inside of the template html file
	//"rolePlay":    roleplay,
	"debug_print": debugPrint,
}

func init() {
	theTemplate = template.Must(template.New("").Funcs(funcMap).ParseGlob("templates/*"))
	dbSessionsCleaned = time.Now() //This is when the program starts, to keep track of the session length
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/roleplay", roleplay)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	refreshSessions(w, req)
	theUser := getUser(w, req)
	showSessions() // for demonstration purposes
	theTemplate.ExecuteTemplate(w, "index.gohtml", theUser)
}

func roleplay(w http.ResponseWriter, req *http.Request) {
	refreshSessions(w, req)
	/* This is the data this should be passed into our parsed template */
	type dataStruct struct {
		AUser   user
		Writer  http.ResponseWriter
		Request *http.Request
	}
	newUser := getUser(w, req) //get User with role permissions
	/* If they aren't already logged in, log out! */
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther) //Go back to the main screen to log in!
		log.Printf("This boy right here ain't logged in!\n")
		return
	}
	/* If the wrong access was here, we could use this...but we allow everyone on the roleplay page..
	for now.
	if newUser.Role != theAdmin {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	*/
	/* Below is just some debug stuff */
	if newUser.Role != theAdmin {
		log.Printf("We have some peasants in here...\n")
	}
	showSessions() // for demonstration purposes
	/* Create data to parse into file */
	passedData := dataStruct{
		AUser:   newUser,
		Writer:  w,
		Request: req,
	}
	theTemplate.ExecuteTemplate(w, "roleplay.gohtml", passedData)
}

func signup(w http.ResponseWriter, req *http.Request) {
	refreshSessions(w, req)
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		log.Printf("Bro, you're already logged in! Why you signing up?\n")
		return
	}
	var u user
	// process form submission
	if req.Method == http.MethodPost {
		log.Printf("Form posted, getting form values!")
		// get form values
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")
		// username taken?
		if _, ok := dbUsers[un]; ok {
			log.Printf("Oh no, this Username is already taken! %v\n", un)
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}
		// create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = sessionLength
		/* For the above, the session gets reset for another 30 seconds, which is that constant in main.go */
		http.SetCookie(w, c)
		/* In dbSessions, we are storing the new session struct with the most recent login time, (now) */
		dbSessions[c.Value] = session{un, time.Now()}
		// store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			log.Printf("Uh oh, we had a problem using bit crypt with our password...\n")
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		u = user{un, bs, f, l, r}
		dbUsers[un] = u
		// redirect
		log.Printf("Okay, we've made our User, back to the main menu!\n")
		log.Printf("Here is our cookie session name: %v\n", c.Name)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	showSessions() // for demonstration purposes
	theTemplate.ExecuteTemplate(w, "signup.gohtml", u)
}

func login(w http.ResponseWriter, req *http.Request) {
	refreshSessions(w, req)
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var u user
	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		// is there a username?
		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = sessionLength
		/* For the above, the session gets reset for another 30 seconds, which is that constant in main.go */
		http.SetCookie(w, c)
		dbSessions[c.Value] = session{un, time.Now()}
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	showSessions() // for demonstration purposes
	theTemplate.ExecuteTemplate(w, "login.gohtml", u)
}

func logout(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("session")
	// delete the session
	delete(dbSessions, c.Value)
	// remove the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	// clean up dbSessions
	/* This cleans out the session that aren't being used every 30 seconds */
	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 30) {
		go cleanSessions()
	}

	http.Redirect(w, req, "/login", http.StatusSeeOther)
}

/* You might not use this session cleaner in production */
/* debug zone */
func debugPrint() bool {
	log.Printf("Hey, you made it here\n")
	return false
}
