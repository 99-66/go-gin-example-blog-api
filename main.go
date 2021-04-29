package main

import (
	"github.com/99-66/go-gin-example-blog-api/repositories"
	"github.com/99-66/go-gin-example-blog-api/server"
	"log"
)

func main() {
	repositories.Init()
	db := repositories.Connections.DB

	sqlDB, ok := db.DB()
	if ok != nil {
		defer sqlDB.Close()
	}

	log.Fatal(server.RunAPI("0.0.0.0:8000"))
}
