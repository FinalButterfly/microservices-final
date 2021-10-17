package main

import (
	"GoNews/pkg/api"
	"GoNews/pkg/polling"
	"GoNews/pkg/storage"
	"GoNews/pkg/storage/postgres"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var maxConfigFileSize int = 1024
var c polling.Config
var fallbackConfig polling.Config = polling.Config{
	Feeds: []string{
		"https://habr.com/ru/rss/hub/go/all/?fl=ru",
		"https://habr.com/ru/rss/best/daily/?fl=ru",
		"https://cprss.s3.amazonaws.com/golangweekly.com.xml",
	},
	Interval: 5,
}

func init() {
	file, err := os.Open("config.json")

	bytes, err := ioutil.ReadAll(file)
	if err == nil {
		err = json.Unmarshal(bytes, &c)
		return
	}
	c = fallbackConfig
	fmt.Println("Using fallback config")
}

// Сервер GoNews.
type server struct {
	db  storage.Interface
	api *api.API
}

func main() {
	var s server

	pass := os.Getenv("postgresdbpass")
	cn := "postgresql://postgres:" + pass + "@localhost/gonews"
	db, err := postgres.New(cn)
	if err != nil {
		log.Fatal(err)
		return
	}

	s.db = db

	s.api = api.New(s.db)

	poller, err := polling.NewPoller(c, s.db)

	if err != nil {
		fmt.Println(err)
		fmt.Println("articles will not be fetched, poller is offline")
		return
	}

	go poller.StartPolling()

	http.ListenAndServe(":4000", s.api.Router())
}
