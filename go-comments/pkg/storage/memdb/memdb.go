package memdb

import "go-comments/pkg/storage"

// The store
type Store struct{}

// Constructor for Store
func New() *Store {
	return new(Store)
}

// Gets news where n = number of entries
func (s *Store) Comments(articleId int) ([]storage.Comment, error) {
	var cs []storage.Comment
	for i := 0; i < len(comments); i++ {
		if comments[i].ArticleId == articleId {
			cs = append(cs, comments[i])
		}
	}

	return cs, nil
}

// Adds a comment to the database
func (s *Store) AddComment(p storage.Comment) error {
	comments = append(comments, p)
	return nil
}

var comments = []storage.Comment{
	{
		Id:        1,
		ParentId:  0,
		ArticleId: 1,
		Content:   "Go is a new language. Although it borrows ideas from existing languages, it has unusual properties that make effective Go programs different in character from programs written in its relatives. A straightforward translation of a C++ or Java program into Go is unlikely to produce a satisfactory result—Java programs are written in Java, not Go. On the other hand, thinking about the problem from a Go perspective could produce a successful but quite different program. In other words, to write Go well, it's important to understand its properties and idioms. It's also important to know the established conventions for programming in Go, such as naming, formatting, program construction, and so on, so that programs you write will be easy for other Go programmers to understand.",
		PubTime:   1627724610,
	},
	{
		Id:        1,
		ParentId:  0,
		ArticleId: 1,
		Content:   "The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine.",
		PubTime:   1627724609,
	},
}
