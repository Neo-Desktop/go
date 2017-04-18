package main

import (
    "bufio"
    "fmt"
    "os"
    "time"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        now := time.Now()
        fmt.Println(now.Format(time.Stamp), scanner.Text())
    }
}
