package main

import (
	"net/http"
)

func getUser(req *http.Request) warCriminal {
	var warBaddie warCriminal

	// get cookie
	theCookie, err := req.Cookie("session")
	if err != nil {
		return warBaddie //There's an error...
	}

	// if the user exists already, get user
	/* If in db sessions, when searching with the cookie value, if we find something NOT BLANK,
	That's a User and we can assign it to warBaddie */
	if dbSessions[theCookie.Value] != "" {
		//Get the Username value at that cookie
		userName := dbSessions[theCookie.Value]
		warBaddie = dbUsers[userName]
	}
	/* I DON'T LIKE THIS FORMATTING
	if un, ok := dbSessions[theCookie.Value]; ok {
		warBaddie = dbUsers[un]
	}
	*/
	return warBaddie
}

func alreadyLoggedIn(req *http.Request) bool {
	cookie, err := req.Cookie("session")
	if err != nil {
		return false
	}
	/* This is going to return true or false...if the cookie.value HAS THE VALUE, then ok should be true and we return true from this */
	userName := dbSessions[cookie.Value] //Looks for the Username in the session
	/* I DON'T LIKE THIS FORMAT
	_, ok := dbUsers[userName]
	*/
	foundUser := true
	/* This sees if the War Criminal in dbUsers compares to an empty warCriminal{} struct,
	(my fill in for an empty or nil value) */
	if dbUsers[userName] == (warCriminal{}) {
		foundUser = true
	}
	return foundUser
}
