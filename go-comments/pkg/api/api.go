package api

import (
	"encoding/json"
	"fmt"
	"go-comments/pkg/storage"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

var ErrorInvalidArticleId = fmt.Errorf("invalid article id")
var ErrorContentFieldEmpty = fmt.Errorf("content field is empty")
var ErrorPubtimeFieldEmpty = fmt.Errorf("pubtime field is empty")

var profanes = []string{"qwerty", "йцукен", "zxvbnm"}

// Default message handler
type Message struct {
	Api     string
	Version float32
	Message string
}

// API
type API struct {
	db     storage.Interface
	router *mux.Router
}

// API constructor
func New(db storage.Interface) *API {
	api := API{
		db: db,
	}
	api.router = mux.NewRouter()
	api.endpoints()
	return &api
}

func (api *API) endpoints() {
	api.router.HandleFunc("/", api.DefaultHandler).
		Methods(http.MethodGet, http.MethodOptions)
	api.router.HandleFunc("/comments", api.CommentsHandler).
		Methods(http.MethodGet).
		Queries("articleId", "{articleId:[0-9]+}")
	api.router.HandleFunc("/comments/add", api.AddCommentHandler).
		Methods(http.MethodPost)
}

func (api *API) Router() *mux.Router {
	return api.router
}

func (api *API) DefaultHandler(w http.ResponseWriter, r *http.Request) {
	message := Message{
		Api:     "go-comments",
		Version: 1.0,
		Message: "bitch",
	}
	bytes, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}

func (api *API) CommentsHandler(w http.ResponseWriter, r *http.Request) {
	qs := mux.Vars(r)
	fmt.Println(qs)
	var articleId int
	if id := qs["articleId"]; id != "" {
		articleId, _ = strconv.Atoi(id)
	}
	if articleId == 0 {
		http.Error(w, ErrorInvalidArticleId.Error(), http.StatusBadRequest)
		return
	}
	comments, err := api.db.Comments(articleId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for i := 0; i < len(comments); i++ {
		if comments[i].Profane {
			comments[i] = comments[len(comments)-1]
			comments = comments[:len(comments)-1]
		}
	}
	if len(comments) == 1 {
		if comments[0].Profane {
			comments = []storage.Comment{}
		}
	}
	bytes, err := json.Marshal(comments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}

func (api *API) AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	var comment storage.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if comment.Content == "" {
		http.Error(w, ErrorContentFieldEmpty.Error(), http.StatusBadRequest)
		return
	}
	if comment.PubTime == 0 {
		http.Error(w, ErrorPubtimeFieldEmpty.Error(), http.StatusBadRequest)
		return
	}
	//instead of introducing a separate service for this functionality, do this check here
	checkForProfanities(&comment)
	err = api.db.AddComment(comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func checkForProfanities(c *storage.Comment) {
	for i := 0; i < len(profanes); i++ {
		if strings.Contains(c.Content, profanes[i]) {
			c.Profane = true
		}
	}
}
