package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

func getUser(w http.ResponseWriter, req *http.Request) user {
	// get cookie
	newCookie, err := req.Cookie("session")
	if err != nil {
		//Create cookie if error is returned and session cookie not found
		sID, _ := uuid.NewV4()
		newCookie = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		log.Printf("New cookie created named %v with value of %v\n", newCookie.Name, newCookie.Value)
	}
	newCookie.MaxAge = sessionLength
	/* For the above, the session gets reset for another 30 seconds, which is that constant in main.go */
	http.SetCookie(w, newCookie) //Set the new cookie created

	// if the user exists already, get user
	var aUser user
	if someone, ok := dbSessions[newCookie.Value]; ok {
		someone.lastActivity = time.Now() //Set the last activity to now
		dbSessions[newCookie.Value] = someone
		aUser = dbUsers[someone.Username] //Find the username from sessions to search our User database for...if that person is returned, asigned it to this User
	}
	return aUser
}

func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		log.Printf("Yo, there was an error logging in: %v\n", err)
		log.Printf("Our cookie's name is: %v\n", c.Name)
		return false
	}
	s, ok := dbSessions[c.Value]
	if ok {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
	}
	_, ok = dbUsers[s.Username]
	// refresh session
	c.MaxAge = sessionLength
	/* For the above, the session gets reset for another 10 seconds, which is that constant in main.go */
	http.SetCookie(w, c)
	return ok
}

func cleanSessions() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	showSessions()              // for demonstration purposes
	for k, v := range dbSessions {
		/* This says, for this current session in the loop, get the time it first logged on. If there's no activity in the last 30 seconds,
		remove this session at k from dbSessions */
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
			delete(dbSessions, k)
		}
	}
	dbSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	showSessions()             // for demonstration purposes
}

// for demonstration purposes
func showSessions() {
	fmt.Println("********")
	for k, aSession := range dbSessions {
		fmt.Println(k, aSession.Username) //This will print a User and the session
	}
	fmt.Println("")
}

//for debug purposes
func refreshSessions(w http.ResponseWriter, req *http.Request) {
	//Get Cookie
	newCookie, err := req.Cookie("session")
	if err != nil {
		log.Printf("Yo, there was an error logging in: %v\n", err)
		//Create cookie if error is returned and session cookie not found
		sID, _ := uuid.NewV4()
		newCookie = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		newCookie.MaxAge = sessionLength
		/* For the above, the session gets reset for another 30 seconds, which is that constant in main.go */
		http.SetCookie(w, newCookie) //Set the new cookie created
		log.Printf("New cookie created named %v with value of %v\n", newCookie.Name, newCookie.Value)
		return
	}
	// refresh session
	newCookie.MaxAge = sessionLength
	/* For the above, the session gets reset for another 10 seconds, which is that constant in main.go */
	http.SetCookie(w, newCookie)
	log.Printf("Session refreshed!\n")
}
