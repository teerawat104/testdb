package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type UserData struct {
	Id              int
	CitizenId       string
	FirstName       string
	LastName        string
	BirthYear       int
	FirstnameFather string
	LastnameFather  string
	FirstnameMother string
	LastnameMother  string
	SoldierId       int
	AddressId       int
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")

	if err != nil {
		fmt.Println("connect fail")
	}
	fmt.Println("connect success")

	defer db.Close()
	read(db)

}

func read(db *sql.DB) []UserData {

	results, _ := db.Query("SELECT * FROM user") // result จะ return ออกมาเป็น struc เราต้องทำ struc มารองรับ

	var userDataList []UserData
	for results.Next() {
		var userData UserData

		err := results.Scan(
			&userData.Id,
			&userData.CitizenId,
			&userData.FirstName,
			&userData.LastName,
			&userData.BirthYear,
			&userData.FirstnameFather,
			&userData.LastnameFather,
			&userData.FirstnameMother,
			&userData.LastnameMother,
			&userData.SoldierId,
			&userData.AddressId,
		)

		if err != nil {
			panic(err.Error())
		}

		userDataList = append(userDataList, userData)
		fmt.Println(userDataList)

	}
	return userDataList

}
