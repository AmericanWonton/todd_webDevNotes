package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	/* This connects and writes to the connection, (if we run 02 at the same time we can see this.*/
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close() //After the connection is writting to the listener, it is closed.

	fmt.Fprintln(conn, "I dialed you.")
}

/* NOTE: YOU HAVE to run the read connection at the same time with this one and the TCP server
And you can keep running this and 02 will continually handle this connection writing to it!*/


/* Might need to run these with the telnet client open to write to */