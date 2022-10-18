package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"test-school/repository"
	"test-school/restapi"
	"test-school/services"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/school_db")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	repo := repository.NewRepository(db)
	service := services.NewService(repo)
	router := restapi.NewAPI(service)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8090"
	}

	log.Println("Starting the HTTP server on port ", port)

	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}

}
