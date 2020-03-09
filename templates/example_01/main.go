package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	//Below is an example of using 'string concatitination to make different files from Go
	//This is kind of a poopy way to go...
	name := "Joseph Keller"
	//We could also take arguments as well
	otherName := os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	htmlString := fmt.Sprint(`
			<!DOCTYPE html>
			<html lang="en">
			<head>
			<meta charset="UTF-8">
			<title>Hello World!</title>
			</head>
			<body>
			<h1>` +
		name +
		`</h1>
		<h2>
		` +
		otherName +
		`</h2>
			</body>
			</html>
				`)
	//We can then use the os.Create package from the Standard Go library to create and serve these html file.
	newFile, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer newFile.Close()
	//This will copy the contents of our concatinated string into that html file.
	io.Copy(newFile, strings.NewReader(htmlString))
}
