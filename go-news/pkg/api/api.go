package api

import (
	"GoNews/pkg/storage"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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
	api.router.HandleFunc("/version", api.DefaultHandler).
		Methods(http.MethodGet, http.MethodOptions)
	api.router.HandleFunc("/news", api.NewsHandler).
		Methods(http.MethodGet, http.MethodOptions).
		Queries("page", "{page:[0-9]+}").
		Queries("pageSize", "{pageSize:[0-9]+}")
	api.router.HandleFunc("/news/filter", api.NewsFilterHandler).
		Methods(http.MethodGet).
		Queries("query", "{query:.+}")
	api.router.HandleFunc("/news", api.NewsSingleHandler).
		Methods(http.MethodGet).
		Queries("articleId", "{articleId:[0-9]+}")
	// api.router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./webapp"))))
}

func (api *API) Router() *mux.Router {
	return api.router
}

func (api *API) DefaultHandler(w http.ResponseWriter, r *http.Request) {
	message := Message{
		Api:     "gonews",
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

func (api *API) NewsSingleHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params)
	var articleId string
	if id, ok := params["articleId"]; ok {
		articleId = id
	}
	id, _ := strconv.Atoi(articleId)
	article, err := api.db.NewsSingle(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bytes, err := json.Marshal(article)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}

func (api *API) NewsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var page, pageSize string
	if tempPage, ok := params["page"]; ok {
		page = tempPage
	}
	if tempSize, ok := params["pageSize"]; ok {
		pageSize = tempSize
	}
	p, _ := strconv.Atoi(page)
	ps, _ := strconv.Atoi(pageSize)
	posts, err := api.db.News(p, ps)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	bytes, err := json.Marshal(posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}

func (api *API) NewsFilterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In news filter")
	params := mux.Vars(r)
	var query string
	if q, ok := params["query"]; !ok {
		http.Error(w, "Query parameter missing", http.StatusBadRequest)
	} else {
		query = q
	}
	fmt.Println(query)
	posts, err := api.db.FilterNews(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bytes, err := json.Marshal(posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}
