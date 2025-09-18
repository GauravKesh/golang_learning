package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func handleConnection(res http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(res, "working")

}

func executeSQLOperations() {
	db, err := sql.Open("mysql", "username:password@(127.0.0.1:3306)/bank?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	// to drop all the existing tables
	{

		queryDelete :=
			`
			DROP TABLE IF EXISTS depositor,customer,account,branch
		;`

		if _, err := db.Exec(queryDelete); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Tables deletion Done")
		}
	}

	fmt.Println("Tables Creation started ")

	//  to create a branch table
	{
		query := `
			CREATE TABLE branch (
			branch_name VARCHAR(20) NOT NULL,
			branch_city VARCHAR(20) NOT NULL,
			assets INT,
			PRIMARY KEY(branch_name)
		);`

		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Tables Creation Done")
		}
	}
	// Account table query
	{
		query := `
			CREATE TABLE account (
			accno INT,
			branch_name VARCHAR(20),
			balance INT,
			PRIMARY KEY(accno),
			FOREIGN KEY(branch_name) REFERENCES branch(branch_name) ON DELETE CASCADE
		);`
		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Tables Creation Done")
		}
	}

	// Customer table creation
	{
		query := `
			CREATE TABLE customer (
			customer_name VARCHAR(20),
			customer_street VARCHAR(120),
			customer_city VARCHAR(20),
			PRIMARY KEY(customer_name)
		);`
		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Tables Creation Done")
		}
	}

	// Depositor table creation
	{
		query := `
			CREATE TABLE depositor (
			customer_name VARCHAR(20),
			accno INT,
			PRIMARY KEY (customer_name, accno),
			FOREIGN KEY (customer_name) REFERENCES customer(customer_name) ON DELETE CASCADE,
			FOREIGN KEY (accno) REFERENCES account(accno) ON DELETE CASCADE
		);`
		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Tables Creation Done")
		}
	}
	fmt.Println("Tables Creation Done.. ")
	fmt.Println(err)
}

func main() {
	r := mux.NewRouter()
	PORT := "8080"
	executeSQLOperations()

	api := r.PathPrefix("/api").Subrouter()

	sql := api.PathPrefix("/sql").Subrouter()

	sql.HandleFunc("/status", handleConnection)

	fmt.Print("Server started")

	if err := http.ListenAndServe(":"+PORT, r); err != nil {
		fmt.Println("Error starting server:", err)
	}

}
