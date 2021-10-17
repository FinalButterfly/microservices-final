package storage

//Comment structure
type Comment struct {
	Id        int    `json:"id"`
	ArticleId int    `json:"articleId"`
	ParentId  int    `json:"parentId"`
	PubTime   int64  `json:"pubTime"`
	Content   string `json:"content"`
	Profane   bool   `json:"-"` // Flag indicating if profane language was used
}

// Database methods
type Interface interface {
	Comments(articleId int) ([]Comment, error) // Get comments by article
	AddComment(comment Comment) error          // Adds a comment to the database
}
