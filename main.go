package main

import (
	"database/sql"
	fmt "fmt"
	log "log"

	_ "github.com/go-sql-driver/mysql"
	uuid "github.com/google/uuid"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Data struct {
	ID   uuid.UUID
	Name string
}

func main() {
	connectionString := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", DBUser, DBPassword, DBName)
	db, err := sql.Open("mysql", connectionString)
	checkErr(err)
	defer db.Close()

	// ADD DATA to table
	result, err := db.Exec("INSERT INTO data (name) VALUES ( 'Mangmang Smangmang')")
	checkErr(err)
	lastInsertID, err := result.LastInsertId()
	checkErr(err)
	fmt.Println("Last Insert ID:", lastInsertID)
	rowsAffected, err := result.RowsAffected()
	checkErr(err)
	fmt.Println("Rows Affected:", rowsAffected)

	// RETRIEVE DATA from table

	rows, err := db.Query("SELECT * FROM data")
	checkErr(err)
	for rows.Next() {
		var data Data
		err = rows.Scan(&data.ID, &data.Name)
		checkErr(err)
		fmt.Println(data.ID, data.Name)
	}
}
