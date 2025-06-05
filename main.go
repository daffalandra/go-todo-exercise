package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/daffalandra/go-todo-exercise/config"
	"github.com/daffalandra/go-todo-exercise/cotrollers/homecontroller"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config.LoadEnv()
	fmt.Println("📦 DB URL:", config.GetDatabaseURL())

	dsn := config.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to DB: " + err.Error())
	}

	fmt.Println("✅ Connected to PostgreSQL", db)

	// Home Page
	http.HandleFunc("/", homecontroller.Welcome)

	// Categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	log.Println("Server Running on Port 8080")
	http.ListenAndServe(":8080", nil)
}
