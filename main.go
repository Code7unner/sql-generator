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
	if err != nil {
		panic(err)
	}
	defer db.DB.Close()

	err = db.CreateTablesQuery()
	if err != nil {
		panic(err)
	}

	categories, users, messages := gen.Generate()

	endGenerationTime := time.Now().Unix()

	log.Printf("Generation data: %d seconds" , endGenerationTime - startGenerationTime)
}