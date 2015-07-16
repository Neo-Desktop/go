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
const version = "0.1.0"

var globalFrames []FrameGroup

// Defines a single Frame
type Frame struct {
	Timeout int64
	Frame   string
}

type FrameGroup struct {
	V4Frame []Frame
	V6Frame []Frame
}

// Do some basic logging for any error
func checkErr(err error) bool {
	if err != nil {
		log.Fatalln("Error encountered:", err)
		return true
	}
	return false
}

// Load a specified file into the Frames slice
func loadFile(fileName string, stringOut chan<- Frame) {
	fileh, err := os.Open(fileName)

	if checkErr(err) {
		return
	}

	scanner := bufio.NewScanner(fileh)

	thisFrame := Frame{
		Timeout: 0,
		Frame:   "START",
	}

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) >= 1 {
			if line[0] > '0' && line[0] <= '9' {
				if thisFrame.Frame != "START" {
					stringOut <- thisFrame
				}
				thisFrame = Frame{}
				uhold, _ := strconv.ParseInt(line, 10, 64)
				thisFrame.Timeout = uhold
			} else {
				thisFrame.Frame = fmt.Sprintln(thisFrame.Frame, line, "\r")
			}
		} else {
			thisFrame.Frame = fmt.Sprintln(thisFrame.Frame, line, "\r")
		}

	}

	stringOut <- thisFrame

	close(stringOut)

}

// Generates a slice of Frames from a filename
// calls loadFile()
func load(fileName string) []Frame {
	log.Println("Loading", fileName)
	Frames := make(chan Frame, 100)

	go loadFile(fileName, Frames)

	outFrame := make([]Frame, 0)

	for range Frames {
		Frame := <-Frames
		outFrame = append(outFrame, Frame)
	}

	return outFrame

}

func PlayFrame(con net.Conn, frames []Frame) {
	for _, frame := range frames {
		fmt.Fprintf(con, "\x1b[2J\x1b[H%s", frame.Frame)
		time.Sleep(time.Duration(frame.Timeout) * time.Second / 15)
	}
}

func main() {
	fmt.Println("Go ASCII Service version", version)

	globalFrames = make([]FrameGroup, 0)

	sw1 := FrameGroup{}
	sem := make(chan string, 2)

	go func(fileName string) {
		log.Println("File:", fileName, "has begun loading")
		sw1.V4Frame = load(fileName)
		sem <- fileName
	}("swv4.txt")

	go func(fileName string) {
		log.Println("File:", fileName, "has begun loading")
		sw1.V6Frame = load(fileName)
		sem <- fileName
	}("swv6.txt")

	log.Println("File:", <-sem, "has finished loading")
	log.Println("File:", <-sem, "has finished loading")

	close(sem)

	globalFrames = append(globalFrames, sw1)

	sock, err := net.Listen("tcp", ":23")

	if checkErr(err) {
		return
	}

	log.Println("Now listening on port 23")

	for {
		conn, err := sock.Accept()
		if checkErr(err) {
			return
		}
		test, err := net.ResolveTCPAddr("tcp", sock.Addr().String())

		if test.IP.To4() != nil {
			log.Println("New v4 User: ", sock.Addr().String())
			go PlayFrame(conn, globalFrames[0].V4Frame)
		} else {
			log.Println("New v6 User: ", sock.Addr().String())
			go PlayFrame(conn, globalFrames[0].V6Frame)
		}
	}
}
