package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"unicode/utf8"
)

// http://semver.org/
const version = "0.1.0"

// Defines a single panel
type Panel struct {
	timeout int64
	panel   string
}

// Do some basic logging for any error
func checkErr(error err) bool {
	if err != nil {
		log.Fatalln("Error encountered:", err)
		return true
	}
	return false
}

// Load a specified file into the panels slice
func loadFile(string fileName, queue chan *stringOut) {
	file, err := os.Open(filename)
	if checkErr(err) {
		return
	}

	stringOut <- "END"

}

func load(string fileName) {
	log.Println("Loading", fileName)
	lines := make(chan string, 100)

	go loadFile(fileName, lines)

	for {
		line := <-lines
		if line != "END" {
			runeValue, width := utf8.DecodeRuneInString(line[0:])

		} else {
			break
		}
	}

}

func main() {
	fmt.Println("Go ASCII Service version", version)

	go load("sw1.txt")

	sock, err := net.Listen("tcp", "21")

	if checkErr(err) {
		return
	}

	for {

	}
}
