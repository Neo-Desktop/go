package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

func handlerSQL(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling request from: ", r.RemoteAddr)

	fmt.Fprintln(w, "Currently registered users:")

	users := []User{}
	err := DB.Select(&users, "select * from users")

	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {

		fmt.Fprintf(w, "ID: %d\tName: %s\tCity: %s\tState: %s\r\n", user.ID.Int64, user.Name.String, user.City.String, user.State.String)
	}

	fmt.Fprintf(w, "\r\n")

	fmt.Fprintln(w, "Currently deployed version is:", version)
}

func handlerSQLJSON(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling request from: ", r.RemoteAddr)
	w.Header().Add("Content-Type", "text/json; charset=UTF-8")

	users := []User{}
	query := "select * from users"
	err := DB.Select(&users, query)

	if err != nil {
		log.Fatal(err)
	}

	val, err := json.MarshalIndent(users, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, string(val))

}

func handlerSQLXML(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling request from: ", r.RemoteAddr)

	users := []User{}
	err := DB.Select(&users, "select * from users")

	w.Header().Add("Content-Type", "text/xml; charset=UTF-8")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "%s", "<Users>")
	for _, user := range users {

		val, err := xml.MarshalIndent(user, "", "\t")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, "%s\r\n\r\n", string(val))
	}
	fmt.Fprintf(w, "%s", "</Users>")

}
