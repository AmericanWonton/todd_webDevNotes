package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	ctx = context.WithValue(ctx, "userID", 777)
	ctx = context.WithValue(ctx, "fname", "Bond")

	results, err := dbAccess(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		return
	}

	fmt.Fprintln(w, results)
}

func dbAccess(ctx context.Context) (int, error) {
	/* This says, 'we're going to time this out in one second if it isn't finished! */
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel() //We don't get an error with context, we get a cancel func

	ch := make(chan int)
	/* For testing purposes, we launch a go routine */
	go func() {
		// ridiculous long running task
		uid := ctx.Value("userID").(int) //Force this value to be an int
		time.Sleep(10 * time.Second)

		// check to make sure we're not running in vain
		// if ctx.Done() has
		if ctx.Err() != nil {
			return
		}

		ch <- uid //Continuously put on that int channel the uid
	}()
	/* The code below will still be running as the go routine goes off on it's own! */
	select {
	case <-ctx.Done(): //When ctx get's 'canceled' after taking 1 second too long
		return 0, ctx.Err()
	case i := <-ch: //If not canceled, put a value on that channel
		return i, nil
	}
}

func bar(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}

// per request variables
// good candidate for putting into context
