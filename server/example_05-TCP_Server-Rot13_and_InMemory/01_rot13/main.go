package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	/* Start by getting our listener */
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()
	/* This is our listener loop forever... */
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		/* When a connection comes in, we handle it... */
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		/* Take the strings coming in from the connection first. Then turn it into a slice of bytes, then passs it to the rot13 function*/
		ln := strings.ToLower(scanner.Text())
		bs := []byte(ln)
		r := rot13(bs)

		fmt.Fprintf(conn, "%s - %s\n\n", ln, r)
	}
}

func rot13(bs []byte) []byte {
	/* This just moves each byte to a spot 13 characters around */
	var r13 = make([]byte, len(bs))
	for i, v := range bs {
		// ascii 97 - 122
		if v <= 109 {
			r13[i] = v + 13
		} else {
			r13[i] = v - 13
		}
	}
	return r13
}

/* Might need to run these with the telnet client open to write to */