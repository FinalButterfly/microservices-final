package storage

// Публикация, получаемая из RSS.
type Article struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	PubTime int64  `json:"pubTime"`
	Link    string `json:"link"`
}

// Database methods
type Interface interface {
	NewsSingle(articleId int) (Article, error)        // Get a single article by articleId
	News(page, pageSize int) (PaginatedResult, error) // Gets all news in a paginated result
	FilterNews(query string) ([]Article, error)       // Filter news by query in title
	AddNews(posts []Article) error                    // Adds multiple posts to the database
}

type PaginatedResult struct {
	Articles  []Article `json:"articles"`
	Page      int       `json:"page"`
	PageSize  int       `json:"pageSize"`
	PageCount int       `json:"pageCount"`
}
