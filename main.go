package main

import (
	"log"
	"net/http"
	"os"

	"go_final_project/pkg/api"
	"go_final_project/pkg/db"
)

func main() {
	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}

	if err := db.Init("scheduler.db"); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	api.Init()

	webDir := "./web"
	http.Handle("/", http.FileServer(http.Dir(webDir)))

	addr := ":" + port
	log.Printf("Server starting on http://localhost%s/", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
