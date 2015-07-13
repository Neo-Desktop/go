package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

const version float64 = 1.5

var SqlClient *sqlx.DB

func init() {
	db, err := sqlx.Open("mysql", "gouser:gopassword@tcp(172.16.0.2:3306)/goGrounds")

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	SqlClient = db
}

type User struct {
	ID          sql.NullInt64  `db:"id"`
	Username    sql.NullString `db:"username"`
	Password    sql.NullString `db:"password"`
	Name        sql.NullString `db:"name"`
	Address     sql.NullString `db:"address"`
	City        sql.NullString `db:"city"`
	State       sql.NullString `db:"state"`
	Zip         sql.NullString `db:"zip"`
	Country     sql.NullString `db:"country"`
	CreatedDate sql.NullString `db:"createdDate"`
	UpdatedDate sql.NullString `db:"updatedDate"`
}

func handler_sql(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Currently registered users:\r\n")

	users := []User{}
	err := SqlClient.Select(&users, "select * from users")

	if err != nil {
		log.Fatal(err)
	}

	for v := range users {

		id, err := users[v].ID.Value()
		if err != nil {
			log.Fatal(err)
		}

		name, err := users[v].Name.Value()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, "ID: %d Name: %s\r\n", id, name)
	}

	fmt.Fprintf(w, "\r\n")

	fmt.Fprintf(w, "Currently deployed version is: %f\r\n", version)
}

func handler_home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Currently deployed version is: %f\r\n", version)
}

func main() {
	defer SqlClient.Close()

	go http.HandleFunc("/", handler_home)
	go http.HandleFunc("/sql/", handler_sql)

	http.ListenAndServe(":8080", nil)
}
