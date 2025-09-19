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

func showAllDataFromTables(db *sql.DB) {
	printTable := func(query string, name string, cols []string) {
		rows, err := db.Query(query)
		if err != nil {
			log.Fatalf("Error fetching data from %s: %v", name, err)
		}
		defer rows.Close()

		fmt.Printf("\n--- %s ---\n", name)

		for rows.Next() {
			// Create a slice of interface{} to hold values
			values := make([]interface{}, len(cols))
			valuePtrs := make([]interface{}, len(cols))
			for i := range cols {
				valuePtrs[i] = &values[i]
			}

			if err := rows.Scan(valuePtrs...); err != nil {
				log.Fatal(err)
			}

			// Print row as key=value pairs
			for i, col := range cols {
				fmt.Printf("%s=%v ", col, values[i])
			}
			fmt.Println()
		}
	}

	// Show Branch
	printTable("SELECT branch_name, branch_city, assets FROM branch", "Branch",
		[]string{"branch_name", "branch_city", "assets"})

	// Show Account
	printTable("SELECT accno, branch_name, balance FROM account", "Account",
		[]string{"accno", "branch_name", "balance"})

	// Show Customer
	printTable("SELECT customer_name, customer_street, customer_city FROM customer", "Customer",
		[]string{"customer_name", "customer_street", "customer_city"})

	// Show Depositor
	printTable("SELECT customer_name, accno FROM depositor", "Depositor",
		[]string{"customer_name", "accno"})
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
	fmt.Println("Data insertion started into specific tables")

	// query to insert data

	// insert data into Branch
	{
		fmt.Println("Data insertion started into Branch table....")

		branchDataInsertionQuery := `
		INSERT INTO branch (branch_name, branch_city, assets)
		VALUES
		('Main', 'New York', 1000000),
		('Downtown', 'Boston', 750000),
		('Uptown', 'Chicago', 500000);
	`

		if _, err = db.Exec(branchDataInsertionQuery); err != nil {
			log.Fatal("Branch insert failed:", err)
		} else {
			fmt.Println("Data insertion successful into Branch table....")
		}
	}

	// insert data into Account
	{
		fmt.Println("Data insertion started into Account table....")

		accountDataInsertionQuery := `
		INSERT INTO account (accno, branch_name, balance)
		VALUES
		(101, 'Main', 5000),
		(102, 'Main', 7000),
		(201, 'Downtown', 8000),
		(301, 'Uptown', 6000);
	`

		if _, err = db.Exec(accountDataInsertionQuery); err != nil {
			log.Fatal("Account insert failed:", err)
		} else {
			fmt.Println("Data insertion successful into Account table....")
		}
	}

	// insert data into Customer
	{
		fmt.Println("Data insertion started into Customer table....")

		customerDataInsertionQuery := `
		INSERT INTO customer (customer_name, customer_street, customer_city)
		VALUES
		('Alice', '5th Ave', 'New York'),
		('Bob', 'Main St', 'Boston'),
		('Charlie', 'Lake View', 'Chicago');
	`

		if _, err = db.Exec(customerDataInsertionQuery); err != nil {
			log.Fatal("Customer insert failed:", err)
		} else {
			fmt.Println("Data insertion successful into Customer table....")
		}
	}

	// insert data into Depositor
	{
		fmt.Println("Data insertion started into Depositor table....")

		depositorDataInsertionQuery := `
		INSERT INTO depositor (customer_name, accno)
		VALUES
		('Alice', 101),
		('Alice', 102),
		('Bob', 201),
		('Charlie', 301);
	`

		if _, err = db.Exec(depositorDataInsertionQuery); err != nil {
			log.Fatal("Depositor insert failed:", err)
		} else {
			fmt.Println("Data insertion successful into Depositor table....")
		}
	}

	fmt.Println("Data insertion  into specific tables completed")

	showAllDataFromTables(db)

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
