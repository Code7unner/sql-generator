package main

import (
	_ "github.com/lib/pq"
	db2 "sql-generator/db"
)

func main()  {
	db, err := db2.InitDatabase()
	if err != nil {
		panic(err)
	}
	defer db.DB.Close()

	err = db.CreateTablesQuery()
	if err != nil {
		panic(err)
	}
}