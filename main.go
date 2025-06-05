package main

import (
	"log"
	"net/http"

	"github.com/daffalandra/go-todo-exercise/config"
	"github.com/daffalandra/go-todo-exercise/controllers/categorycontroller"
	"github.com/daffalandra/go-todo-exercise/controllers/homecontroller"
)

func main() {
	// Load environment variables
	if err := config.LoadEnv(); err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}

	// Initialize database connection
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer func() {
		if err := config.CloseDB(db); err != nil {
			log.Printf("Failed to close DB: %v", err)
		}
	}()

	log.Println("✅ Connected to PostgreSQL")

	// Home Page
	http.HandleFunc("/", homecontroller.Welcome)

	// Categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	log.Println("Server Running on Port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
