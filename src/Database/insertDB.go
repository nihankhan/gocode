package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = "nihankhan"
	hostname = "127.0.0.1:3306"
	dbname   = "Nihan"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}

func main() {
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	user := []struct {
		ID      int
		Name    string
		Section string
		GPA     float32
	}{
		{22, "Badsha", "Science", 4.50},
		{23, "Bijoy", "Science", 4.50},
		{24, "Monir", "Commerce", 4.60},
		{25, "Ashiq", "Arts", 4.00},
	}

	stmt, err := db.Prepare("INSERT INTO Nihan.student (ID, Name, Section, GPA) VALUES (?, ?, ?, ?)")

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	for _, users := range user {
		_, err := stmt.Exec(users.ID, users.Name, users.Section, users.GPA)

		if err != nil {
			panic(err)
		}
	}

}
