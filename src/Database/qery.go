package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// this is a strut for all rows for coloumn
type Nihan struct {
	ID      int
	Name    string
	Section string
	GPA     float64
}

func main() {
	db, err := sql.Open("mysql", "root:nihankhan@tcp(127.0.0.1:3306)/Nihan")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	qry, err := db.Query("SELECT * FROM student")

	if err != nil {
		panic(err)
	}

	defer qry.Close()

	for qry.Next() {
		var nihan Nihan

		err := qry.Scan(&nihan.ID, &nihan.Name, &nihan.Section, &nihan.GPA)

		if err != nil {
			panic(err)
		}

		fmt.Printf("%v\n", nihan)
	}
}
