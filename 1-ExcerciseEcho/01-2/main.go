package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Hello	")
	db, err := sql.Open("mysql", "root:dadagiri6@tcp(localhost:3306)/IDandPass")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	rows, err := db.Query("INSERT INTO information (username, password) VALUES ('Arakyn523232323', 'dadagiri323232324326')")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	fmt.Println("Sucesffiyl logged in and inserted")
	add(2, 3)
}

// function to add numbers and print
func add(a int, b int) {
	fmt.Println(a + b)
}
// function that converts human years into dog years
func dogYears(humanYears int) int {
	return humanYears * 7
}
