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

var globalPanels []PanelGroup

// Defines a single panel
type Panel struct {
	Timeout int64
	Panel   string
}

type PanelGroup struct {
	V4Panel []Panel
	V6Panel []Panel
}

// Do some basic logging for any error
func checkErr(err error) bool {
	if err != nil {
		log.Fatalln("Error encountered:", err)
		return true
	}
	return false
}

// Load a specified file into the panels slice
func loadFile(fileName string, stringOut chan<- Panel) {
	fileh, err := os.Open(fileName)

	if checkErr(err) {
		return
	}

	scanner := bufio.NewScanner(fileh)

	thisPanel := Panel{
		Timeout: 0,
		Panel:   "START",
	}

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) >= 1 {
			if line[0] > '0' && line[0] <= '9' {
				if thisPanel.Panel != "START" {
					stringOut <- thisPanel
				}
				thisPanel = Panel{}
				uhold, _ := strconv.ParseInt(line, 10, 64)
				thisPanel.Timeout = uhold
				// fmt.Println("start a new panel, wait time", thisPanel.Timeout)
			} else {
				thisPanel.Panel = fmt.Sprintln(thisPanel.Panel, line)
			}
		} else {
			thisPanel.Panel = fmt.Sprintln(thisPanel.Panel, line)
		}

	}

	stringOut <- thisPanel

	close(stringOut)

}

// Generates a slice of panels from a filename
// calls loadFile()
func load(fileName string) []Panel {
	log.Println("Loading", fileName)
	panels := make(chan Panel, 100)

	go loadFile(fileName, panels)

	outPanel := make([]Panel, 0)

	for range panels {
		panel := <-panels
		outPanel = append(outPanel, panel)
	}

	return outPanel

}

func PlayPanel(con net.Conn, frames []Panel) {
	for _, frame := range frames {
		fmt.Fprintf(con, "\x1b[2J%s", frame.Panel)
		time.Sleep(time.Duration(frame.Timeout) * time.Second)
	}
}

func main() {
	fmt.Println("Go ASCII Service version", version)

	globalPanels = make([]PanelGroup, 0)

	sw1 := PanelGroup{}
	sem := make(chan string, 2)

	go func(fileName string) {
		log.Println("File:", fileName, "has begun loading")
		sw1.V4Panel = load(fileName)
		sem <- fileName
	}("swv4.txt")

	go func(fileName string) {
		log.Println("File:", fileName, "has begun loading")
		sw1.V6Panel = load(fileName)
		sem <- fileName
	}("swv6.txt")

	log.Println("File:", <-sem, "has finished loading")
	log.Println("File:", <-sem, "has finished loading")

	close(sem)

	globalPanels = append(globalPanels, sw1)

	sock, err := net.Listen("tcp", ":21")

	if checkErr(err) {
		return
	}

	log.Println("Now listening on port 21")

	for {
		conn, err := sock.Accept()
		if checkErr(err) {
			return
		}
		test, err := net.ResolveTCPAddr("tcp", sock.Addr().String())

		if test.IP.To4() != nil {
			log.Println("New v4 User: ", sock.Addr().String())
			go PlayPanel(conn, globalPanels[0].V4Panel)
		} else {
			log.Println("New v6 User: ", sock.Addr().String())
			go PlayPanel(conn, globalPanels[0].V6Panel)
		}
	}
}
