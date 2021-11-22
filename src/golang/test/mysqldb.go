package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:nihankhan@tcp(127.0.0.1:3306)/Nihan")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO user VALUES(1111, 'Robin', 'Science', 5.00)")

	if err != nil {
		panic(err)
	}

	defer insert.Close()

	fmt.Println("Data inster into database successfully!!!")
}
