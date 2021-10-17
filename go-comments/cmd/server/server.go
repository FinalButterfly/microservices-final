package main

import (
	"fmt"
	"go-comments/pkg/api"
	"go-comments/pkg/storage"
	"go-comments/pkg/storage/postgres"
	"log"
	"net/http"
	"os"
)

// Сервер GoNews.
type server struct {
	db  storage.Interface
	api *api.API
}

func main() {
	var s server

	pass := os.Getenv("postgresdbpass")
	cn := "postgresql://postgres:" + pass + "@localhost/gocomments"
	db, err := postgres.New(cn)
	if err != nil {
		log.Fatal(err)
		return
	}

	s.db = db

	s.api = api.New(s.db)

	if err != nil {
		fmt.Println(err)
		return
	}

	http.ListenAndServe(":5000", s.api.Router())
}
