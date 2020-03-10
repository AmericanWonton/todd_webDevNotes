package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	//Step one, create a listener
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close() //Defer closing this Listener

	//eternally loop until we get a connection
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		/* Launch a 'GoRoutine',(simultaneous go function) that listens to this connection when something comes in */
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	/* What does bufio scanner do?
	It scans stuff, returning to the 'next token', (returning a boolean. It returns false when scan stops at end or gets an error.)
	It basically scans line by line until finished.
	*/
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()

	// we never get here
	// we have an open stream connection so it's waiting on the phone like, "yo, got anything else to say? Can we close?"
	// how does the above reader know when it's done?
	fmt.Println("Code got here.")
}

/* Run this and then open a browser at: localhost:8080
Use CTRL + C to stop
You can also add stuff on like this: localhost:8080/elrequestdemio
*/
