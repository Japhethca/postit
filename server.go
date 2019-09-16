package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	router := gin.Default()
	dbConnStr := ""
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		panic(fmt.Sprint("Failed to open db connection. ", err))
	}

	svc := service{
		router:      router,
		db:          db,
		controllers: &controllers{db},
	}
	svc.InitRoutes()
	http.ListenAndServe(":8001", &svc)
}
