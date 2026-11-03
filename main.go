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

	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		dbFile = "scheduler.db"
	}

	if err := db.Init(dbFile); err != nil {
		log.Printf("ошибка инициализации базы данных: %v", err)
		return
	}
	defer db.Close()

	api.Init()

	webDir := "./web"
	http.Handle("/", http.FileServer(http.Dir(webDir)))

	addr := ":" + port
	log.Printf("сервер запускается на http://localhost%s/", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Printf("ошибка запуска сервера: %v", err)
	}
}
