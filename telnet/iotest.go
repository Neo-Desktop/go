package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := "swv4.txt"
	fmt.Println("Loading", filename)

	fileh, _ := os.Open(filename)
	scanner := bufio.NewScanner(fileh)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) >= 1 {
			if line[0] > '0' && line[0] <= '9' {
				fmt.Println("found a new panel, wait time", line)
			}
		}
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error reading file", err)
	}
}
