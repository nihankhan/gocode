package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	fmt.Println("Golang Mysql >>>")

	db, err := sql.Open("mysql", "root:nihankhan@tcp(127.0.0.1:3306)/Nihan")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO user VALUES(111, 'Nihan', 'Science', 5.00)")
	insert, err = db.Query("INSERT INTO user VALUES(113, 'Khan', 'Science', 4.00)")

	if err != nil {
		panic(err)
	}

	defer insert.Close()

	fmt.Println("Data inserted successfully...")

}

