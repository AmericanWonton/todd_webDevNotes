package main

import (
	"net/http"
)

func getUser(req *http.Request) user {
	var u user

	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		return u
	}

	// if the user exists already, get user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}
	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	/* This is going to return true or false...if the cookie.value HAS THE VALUE, then ok should be true and we return true from this */
	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}
