package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	/* Below is part of the 'net' package. It returns a 'listener' and an err. It asks, 'what do you want to listen on'?
	We have TCP server. Then it asks for the port, (we have :8080, or local)*/
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close() //Need to defer to close it.

	/* Here we just loop forever until the connection is closed */
	for {
		/* A connection has a read/write type */
		connection, err := listener.Accept() //If somebody calls in, we accept
		if err != nil {
			log.Println(err)
			continue //This continue just continues the loop, even if there's an error.
		}

		/* Since a connection is of type read and write, we can write stuff out from the server. */
		io.WriteString(connection, "\nHello from TCP server\n")
		fmt.Fprintln(connection, "How is your day?")
		fmt.Fprintf(connection, "%v", "Well, I hope!")

		connection.Close() //Need to close the connection once finished.
	}
}
