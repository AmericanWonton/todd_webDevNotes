package main

import (
	"fmt"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

func getUser(w http.ResponseWriter, req *http.Request) user {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

	}
	c.MaxAge = sessionLength
	/* For the above, the session gets reset for another 30 seconds, which is that constant in main.go */
	http.SetCookie(w, c)

	// if the user exists already, get user
	var u user
	if s, ok := dbSessions[c.Value]; ok {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
		u = dbUsers[s.un]
	}
	return u
}

func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := dbSessions[c.Value]
	if ok {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
	}
	_, ok = dbUsers[s.un]
	// refresh session
	c.MaxAge = sessionLength
	/* For the above, the session gets reset for another 30 seconds, which is that constant in main.go */
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
	for k, v := range dbSessions {
		fmt.Println(k, v.un) //This will print a User and the session
	}
	fmt.Println("")
}
