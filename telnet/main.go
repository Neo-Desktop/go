package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

// http://semver.org/
const version = "0.2.0"

// Defines a single Frame
type Frame struct {
	Timeout int64
	Frame   string
}

// Load a specified file into the Frames slice
func parseFile(fileName string, stringOut chan<- Frame) {
	fileh, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("Error encountered:", err)
		return
	}

	scanner := bufio.NewScanner(fileh)

	nextFrame := Frame{
		Timeout: 0,
		Frame:   "START",
	}

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) >= 1 {
			if line[0] > '0' && line[0] <= '9' {
				if nextFrame.Frame != "START" {
					stringOut <- nextFrame
				}
				nextFrame = Frame{}
				uhold, _ := strconv.ParseInt(line, 10, 64)
				nextFrame.Timeout = uhold
			} else {
				nextFrame.Frame = fmt.Sprintln(nextFrame.Frame, line, "\r")
			}
		} else {
			nextFrame.Frame = fmt.Sprintln(nextFrame.Frame, line, "\r")
		}

	}

	stringOut <- nextFrame

	close(stringOut)

}

// Generates a slice of Frames from a filename
// calls parseFile()
func loadFile(fileName string) []Frame {
	log.Println("Loading", fileName)
	Frames := make(chan Frame, 100)

	go parseFile(fileName, Frames)

	outFrame := make([]Frame, 0)

	for range Frames {
		Frame := <-Frames
		outFrame = append(outFrame, Frame)
	}

	return outFrame

}

// Sends frames over the wire at 15 fps
func PlayFrame(con net.Conn, frames []Frame) {
	for _, frame := range frames {
		fmt.Fprintf(con, "\x1b[2J\x1b[H%s", frame.Frame)
		time.Sleep(time.Duration(frame.Timeout) * time.Second / 15)
	}
}

func main() {
	fmt.Println("Go ASCII Service version", version)

	v4Frames := make([]Frame, 0)
	v6Frames := make([]Frame, 0)

	sem := make(chan string, 2)

	go func(fileName string) {
		log.Println("File:", fileName, "has begun loading")
		v4Frames = loadFile(fileName)
		sem <- fileName
	}("swv4.txt")

	go func(fileName string) {
		log.Println("File:", fileName, "has begun loading")
		v6Frames = loadFile(fileName)
		sem <- fileName
	}("swv6.txt")

	log.Println("File:", <-sem, "has finished loading")
	log.Println("File:", <-sem, "has finished loading")

	close(sem)

	sock, err := net.Listen("tcp", ":2323")

	if err != nil {
		log.Fatalln("Error encountered:", err)
		return
	}

	log.Println("Now listening on port 23")

	for {
		conn, err := sock.Accept()
		if err != nil {
			log.Fatalln("Error encountered:", err)
			break
		}

		test, err := net.ResolveTCPAddr("tcp", sock.Addr().String())
		if err != nil {
			log.Fatalln("Error encountered:", err)
			break
		}

		if test.IP.To4() != nil {
			log.Println("New v4 User: ", sock.Addr().String())
			go PlayFrame(conn, v4Frames)
		} else {
			log.Println("New v6 User: ", sock.Addr().String())
			go PlayFrame(conn, v6Frames)
		}
	}
}
