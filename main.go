package main

import (
	_ "github.com/lib/pq"
	"log"
	db2 "sql-generator/db"
	"sql-generator/gen"
	"time"
)

func main()  {
	db, err := db2.InitDatabase()
	errorHandler(err)
	defer db.DB.Close()

	log.Println("Successfully connected to database.\n")

	err = db.CreateTablesQuery()
	errorHandler(err)

	startGenerationTime := time.Now().Unix()

	g, err := gen.InitGenerator(db)
	errorHandler(err)

	g.Generate()

	endGenerationTime := time.Now().Unix()
	log.Printf("Generation data: %d seconds" , endGenerationTime - startGenerationTime)

	startInsertingTime := time.Now().Unix()

	err = g.InsertUsers()
	errorHandler(err)

	err = g.InsertCategories()
	errorHandler(err)

	err = g.InsertMessages()
	errorHandler(err)

	endInsertingTime := time.Now().Unix()
	log.Printf("Inserting data: %d seconds", endInsertingTime - startInsertingTime)
}

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}