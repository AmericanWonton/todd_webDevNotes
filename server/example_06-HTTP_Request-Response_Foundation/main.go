package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	/* Start by listening for the TCP server on localhost. On the web, it will just be port 80 */
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()
	/* Run an eternal loop to accept anything that comes in */
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn) //Run a goroutine to handle the incoming connection
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// read request
	request(conn)

	// write response
	respond(conn)
}

func request(conn net.Conn) {
	i := 0 //This is just a flag signaler in the program
	scanner := bufio.NewScanner(conn)
	/* This  is going to scan the entirety of the connection. When we get to the first line of that scan,
	it should be the request line first, (i == 0, part a).*/
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line (Part a)
			/* strings.Fields is a string function that splits a series of string characters up by the white space.
			So what Todd is doing here is just getting 'GET', cuz it's spot 0 in that slice of byte sections */
			m := strings.Fields(ln)[0]
			fmt.Println("***METHOD", m) //So at line 0, it should print ***METHOD GET. Based on what M is, we
			//can make some dope conditional logic based on what request it is.
			/* This is for the exercise Todd asked us to do, getting the URL */
			theURL := "localhost:8080" + strings.Fields(ln)[1]
			fmt.Println("***URL:", theURL)
		}
		if ln == "" {
			/* there's the request line, then the headers...and then a blank line, meaning we are done. */
			break
		}
		i++
	}
}

func respond(conn net.Conn) {
	/* This is a base-ass method for html, no parsing files. */
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello 
	World</strong><br><br><strong>Easy peasy, lemon Febreezy</strong></body></html>`
	/* Here we print, TO THE CONNECTION, our http response */
	/* This is formatted to HTTP Protocol. Status Line:
	Version: HTTP/1.1
	Code: 200
	Reason Phrase: OK
	CharacterReturn/LineFeed: \r\n */
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	/* Next is our header we send to the browser */
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	/* Here is our content type */
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	/* A blank line */
	fmt.Fprint(conn, "\r\n")
	/* The body */
	fmt.Fprint(conn, body)

}
