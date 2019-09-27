package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/japhethca/postit-api/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to read .env configuration file.")
	}

	dbConnStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatal(fmt.Sprint("Failed to open db connection. ", err))
	}

	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./web/build", true)))
	svc := service.New(router, db)
	svc.Init()
	http.ListenAndServe(":8001", svc)
}
