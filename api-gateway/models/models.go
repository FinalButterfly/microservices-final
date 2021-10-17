package models

type Article struct {
	Id       int       `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	PubTime  int64     `json:"pubTime"`
	Link     string    `json:"link"`
	Comments []Comment `json:"comments"`
}

type ArticleShort struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	PubTime int64  `json:"pubTime"`
	Link    string `json:"link"`
}

type Comment struct {
	Id        int    `json:"id"`
	ParentId  int    `json:"parentId"`
	ArticleId int    `json:"articleId"`
	PubTime   int64  `json:"pubTime"`
	Content   string `json:"content"`
}

type PaginatedResult struct {
	Articles  []Article `json:"articles"`
	Page      int       `json:"page"`
	PageSize  int       `json:"pageSize"`
	PageCount int       `json:"pageCount"`
}
