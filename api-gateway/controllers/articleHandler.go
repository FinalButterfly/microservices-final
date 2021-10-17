package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"sync"

	"api-gateway/models"
	u "api-gateway/utils"

	"github.com/gorilla/mux"
)

var commentsChan chan []models.Comment = make(chan []models.Comment, 500)
var errorsChan chan error = make(chan error, 30)
var wg sync.WaitGroup = sync.WaitGroup{}

type ArticleHandler struct {
	config models.Config
}

type IArticleHandler interface {
	GetArticleHandler(w http.ResponseWriter, r *http.Request)
	GetArticlesHandler(w http.ResponseWriter, r *http.Request)
	GetArticlesDetailedHandler(w http.ResponseWriter, r *http.Request)
	FilterArticlesHandler(w http.ResponseWriter, r *http.Request)
	GetCommentsHandler(w http.ResponseWriter, r *http.Request)
	AddCommentHandler(w http.ResponseWriter, r *http.Request)
}

func NewArticleHandler(c models.Config) *ArticleHandler {
	a := ArticleHandler{}
	a.config = c
	return &a
}

func (a *ArticleHandler) GetArticleHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	fmt.Println("In GetArticle")
	fmt.Println(params)

	var articleId string
	if p, ok := params["articleId"]; ok {
		articleId = p
	}

	var article models.Article

	fmt.Println(fmt.Sprintf("%s%s", a.getArticleEndpoint(), fmt.Sprintf("?articleId=%s", articleId)))
	resp, err := http.Get(fmt.Sprintf("%s%s", a.getArticleEndpoint(), fmt.Sprintf("?articleId=%s", articleId)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&article)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	wg.Add(1)
	go a.getCommentsByArticleId(article.Id)
	wg.Wait()

	close(commentsChan)
	if c, ok := <-commentsChan; ok {
		article.Comments = c
	}
	commentsChan = make(chan []models.Comment, 500)

	u.Respond(w, article)
	return
}

func (a *ArticleHandler) GetArticlesHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	fmt.Println("In GetArticles")
	fmt.Println(params)

	var page, pageSize string
	if p, ok := params["page"]; ok {
		page = p
	}
	if ps, ok := params["pageSize"]; ok {
		pageSize = ps
	}

	var result models.PaginatedResult

	fmt.Println(fmt.Sprintf("%s%s", a.getArticleEndpoint(), fmt.Sprintf("?page=%s&pageSize=%s", page, pageSize)))
	resp, err := http.Get(fmt.Sprintf("%s%s", a.getArticleEndpoint(), fmt.Sprintf("?page=%s&pageSize=%s", page, pageSize)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if resp.StatusCode == 200 || resp.StatusCode == 204 {
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		u.Respond(w, result)
		return
	} else {
		http.Error(w, resp.Status, resp.StatusCode)
	}
}

func (a *ArticleHandler) GetArticlesDetailedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In GetArticlesDetailed")
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

	articles := a.getArticlesDetailedInternal(p, ps)
	close(errorsChan)
	for e := range errorsChan {
		if e != nil {
			u.Respond(w, e)
			return
		}
	}
	errorsChan = make(chan error, 30)
	u.Respond(w, articles)
	return
}

func (a *ArticleHandler) FilterArticlesHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	fmt.Println("In FilterArticles")
	fmt.Println(params)
	var articles []models.ArticleShort
	var query string
	if q, ok := params["query"]; !ok {
		u.Error(w, u.Message(http.StatusBadRequest, "query parameter missing"), http.StatusBadRequest)
		return
	} else {
		query = q
	}

	fmt.Println(fmt.Sprintf("%s%s%s", a.getArticleEndpoint(), "/filter?query=", query))
	resp, err := http.Get(fmt.Sprintf("%s%s%s", a.getArticleEndpoint(), "/filter?query=", query))
	if err != nil {
		errorsChan <- err
		return
	}

	if resp.StatusCode == 200 || resp.StatusCode == 204 {
		err = json.NewDecoder(resp.Body).Decode(&articles)
		if err != nil {
			errorsChan <- err
			return
		}

		u.Respond(w, articles)
		return
	} else {
		http.Error(w, resp.Status, resp.StatusCode)
	}
}

func (a *ArticleHandler) GetCommentsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	fmt.Println("In GetComments")

	var articleId string
	if id, ok := params["articleId"]; !ok {
		u.Error(w, u.Message(http.StatusBadRequest, "article parameter missing"), http.StatusBadRequest)
		return
	} else {
		articleId = id
	}

	var comments []models.Comment

	fmt.Println(fmt.Sprintf("%s%s", a.getCommentsEndpoint(), "?articleId="+articleId))
	resp, err := http.Get(fmt.Sprintf("%s%s", a.getCommentsEndpoint(), "?articleId="+articleId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if resp.StatusCode == 200 || resp.StatusCode == 204 {
		err = json.NewDecoder(resp.Body).Decode(&comments)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		u.Respond(w, comments)
		return
	} else {
		http.Error(w, resp.Status, resp.StatusCode)
	}
}

func (a *ArticleHandler) AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In AddComment")
	comment := &models.Comment{}
	err := json.NewDecoder(r.Body).Decode(comment)
	if err != nil {
		u.Respond(w, u.Message(http.StatusBadRequest, err.Error()))
		return
	}

	jsonBody, err := json.Marshal(comment)
	if err != nil {
		u.Respond(w, u.Message(http.StatusInternalServerError, err.Error()))
	}

	fmt.Println(fmt.Sprintf("%s/add", a.getCommentsEndpoint()))
	fmt.Println(comment)

	resp, err := http.Post(fmt.Sprintf("%s/add", a.getCommentsEndpoint()), "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		http.Error(w, resp.Status, resp.StatusCode)
		return
	}
	fmt.Println(resp.Status)
	if resp.Status != "200 OK" {
		http.Error(w, resp.Status, resp.StatusCode)
		return
	}

	u.Respond(w, resp)
	return
}

// Внутренний механизм для определения endpoint новостного сервера
// По умолчанию пока отправляет 0
func (a *ArticleHandler) getArticleEndpoint() string {
	endpoint := a.config.NewsServerEndpointTemplate
	p := rand.Intn(1)
	if p == 1 {
		p = 0
	}

	return fmt.Sprintf(endpoint, p)
}

// Внутренний механизм для определения endpoint сервера комментариев
// По умолчанию пока отправляет 0
func (a *ArticleHandler) getCommentsEndpoint() string {
	endpoint := a.config.CommentsServerEndpointTemplate
	p := rand.Intn(1)
	if p == 1 {
		p = 0
	}

	return fmt.Sprintf(endpoint, p)
}

func (a *ArticleHandler) getArticlesDetailedInternal(page, pageSize int) (result models.PaginatedResult) {
	resp, err := http.Get(fmt.Sprintf("%s%s", a.getArticleEndpoint(), fmt.Sprintf("?page=%d&pageSize=%d", page, pageSize)))
	if err != nil {
		errorsChan <- err
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		errorsChan <- err
		return
	}

	wg.Add(len(result.Articles))
	for i := 0; i < len(result.Articles); i++ {
		go a.getCommentsByArticleId(result.Articles[i].Id)
	}
	wg.Wait()
	close(commentsChan)
	for c := range commentsChan {
		if c != nil && len(c) > 0 {
			for i := 0; i < len(result.Articles); i++ {
				fmt.Println("c", c)
				if c[0] != (models.Comment{}) {
					if c[0].ArticleId == result.Articles[i].Id {
						result.Articles[i].Comments = c
					}
				}
			}
		}
	}
	commentsChan = make(chan []models.Comment, 500)

	return
}

func (a *ArticleHandler) getCommentsByArticleId(id int) {
	defer wg.Done()
	var comments []models.Comment
	resp, err := http.Get(fmt.Sprintf("%s?articleId=%d", a.getCommentsEndpoint(), id))
	if err != nil {
		errorsChan <- err
	}

	err = json.NewDecoder(resp.Body).Decode(&comments)
	if err != nil {
		errorsChan <- err
	}

	commentsChan <- comments
}
