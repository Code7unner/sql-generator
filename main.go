package main

import (
	_ "github.com/lib/pq"
	"log"
	db2 "sql-generator/db"
	"sql-generator/gen"
	"time"
)

func main()  {
	startGenerationTime := time.Now().Unix()

	db, err := db2.InitDatabase()
	errorHandler(err)
	defer db.DB.Close()

	err = db.CreateTablesQuery()
	errorHandler(err)

	categories, users, messages := gen.Generate()

	endGenerationTime := time.Now().Unix()

	log.Printf("Generation data: %d seconds" , endGenerationTime - startGenerationTime)

	startInsertingTime := time.Now().Unix()

	err = db.InsertUsers(users)
	errorHandler(err)

	err = db.InsertCategories(categories)
	errorHandler(err)

	err = db.InsertMessages(messages)
	errorHandler(err)

	endInsertingTime := time.Now().Unix()

	log.Printf("Inserting data: %d seconds", endInsertingTime - startInsertingTime)
}

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}