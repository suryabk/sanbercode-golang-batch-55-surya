package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Employee struct {
	ID        int
	Full_name string
	Email     string
	Age       int
	Division  string
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "surya27bk"
	dbname   = "db-go-sql"
)

var (
	db  *sql.DB
	err error
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")

	createEmployee()
}

func createEmployee() {
	var employee = Employee{}

	sqlStatement := `
	INSERT INTO employees (full_name, email, age, division)
	VALUES ($1, $2, $3, $4)
	RETURNING *
	`

	err = db.QueryRow(sqlStatement, "Surya Kusuma", "surya@gmail.com", 24, "IT").Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Employee Data : %+v\n", employee)
}
