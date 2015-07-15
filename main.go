package main

import (
	"fmt"
	"net/http"
)

const version = "1.0.5"

func main() {
	fmt.Println("Go server version", version)

	DB = sqlInit()
	defer sqlClose()

	router := http.NewServeMux()

	router.HandleFunc("/", handlerHome)
	router.HandleFunc("/sql/", handlerSQL)
	router.HandleFunc("/sql/json", handlerSQLJSON)
	router.HandleFunc("/sql/xml", handlerSQLXML)

	http.ListenAndServe(":8081", router)

}
