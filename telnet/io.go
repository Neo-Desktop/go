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
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error reading file", err)
	}
}
