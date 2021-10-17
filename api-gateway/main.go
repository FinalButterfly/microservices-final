package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	"api-gateway/controllers"
	"api-gateway/models"
)

var maxConfigFileSize int = 1024
var c models.Config
var l *log.Logger = log.Default()
var fallbackConfig models.Config = models.Config{
	NewsServerCount:                1,
	CommentsServerCount:            1,
	NewsServerEndpointTemplate:     "http://localhost:400x",
	CommentsServerEndpointTemplate: "http://localhost:500x",
}

func init() {
	rand.Seed(time.Now().UnixNano())
	file, err := os.Open("config.json")

	bytes, err := ioutil.ReadAll(file)
	if err == nil {
		err = json.Unmarshal(bytes, &c)
		return
	}
	c = fallbackConfig
	l.Println("Using fallback config")
}

func main() {
	router := mux.NewRouter()
	// router.Use(app.JwtAuthentication)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	a := controllers.NewArticleHandler(c)

	router.HandleFunc("/news", a.GetArticleHandler).
		Methods(http.MethodGet).
		Queries("articleId", "{articleId:[0-9]+}")

	router.HandleFunc("/news/full", a.GetArticlesDetailedHandler).
		Methods("GET").
		Queries("page", "{page:[0-9]+}", "pageSize", "{pageSize:[0-9]+}")

	router.HandleFunc("/news/filter", a.FilterArticlesHandler).
		Methods("GET").
		Queries("query", "{query:.+}")

	router.HandleFunc("/news", a.GetArticlesHandler).
		Methods("GET").
		Queries("page", "{page:[0-9]+}", "pageSize", "{pageSize:[0-9]+}")

	router.HandleFunc("/news/comments", a.GetCommentsHandler).
		Methods("GET").
		Queries("articleId", "{articleId:.+}")

	router.HandleFunc("/news/comments/add", a.AddCommentHandler).
		Methods("POST")

	router.HandleFunc("/", controllers.DefaultHandler).
		Methods("GET")

	fmt.Println("Running on port", port)

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		fmt.Print(err)
	}
}
