package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerHome(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling request from: ", r.RemoteAddr)
	fmt.Fprintln(w, "Currently deployed version is:", version)
}
