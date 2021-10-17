package memdb

import "GoNews/pkg/storage"

// The store
type Store struct{}

// Constructor for Store
func New() *Store {
	return new(Store)
}

// Gets news where n = number of entries
func (s *Store) News(limit int) ([]storage.Article, error) {
	return posts[0:limit], nil
}

// Adds multiple posts to the database
func (s *Store) AddNews(p []storage.Article) error {
	posts = append(posts, p...)
	return nil
}

var posts = []storage.Article{
	{
		Id:      1,
		Title:   "Effective Go",
		Content: "Go is a new language. Although it borrows ideas from existing languages, it has unusual properties that make effective Go programs different in character from programs written in its relatives. A straightforward translation of a C++ or Java program into Go is unlikely to produce a satisfactory result—Java programs are written in Java, not Go. On the other hand, thinking about the problem from a Go perspective could produce a successful but quite different program. In other words, to write Go well, it's important to understand its properties and idioms. It's also important to know the established conventions for programming in Go, such as naming, formatting, program construction, and so on, so that programs you write will be easy for other Go programmers to understand.",
		PubTime: 1627724610,
		Link:    "https://google.com/imghp",
	},
	{
		Id:      2,
		Title:   "The Go Memory Model",
		Content: "The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine.",
		PubTime: 1627724609,
		Link:    "https://google.com/imghp",
	},
	{
		Id:      2,
		Title:   "The Go Memory Model",
		Content: "The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine.",
		PubTime: 1627724609,
		Link:    "https://google.com/imghp",
	},
	{
		Id:      2,
		Title:   "The Go Memory Model",
		Content: "The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine.",
		PubTime: 1627724609,
		Link:    "https://google.com/imghp",
	},
	{
		Id:      2,
		Title:   "The Go Memory Model",
		Content: "The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine.",
		PubTime: 1627724609,
		Link:    "https://google.com/imghp",
	},
	{
		Id:      2,
		Title:   "The Go Memory Model",
		Content: "The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine.",
		PubTime: 1627724609,
		Link:    "https://google.com/imghp",
	},
	{
		Id:      2,
		Title:   "The Go Memory Model",
		Content: "The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine.",
		PubTime: 1627724609,
		Link:    "https://google.com/imghp",
	},
	{
		Id:      2,
		Title:   "The Go Memory Model",
		Content: "The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine.",
		PubTime: 1627724609,
		Link:    "https://google.com/imghp",
	},
	{
		Id:      2,
		Title:   "The Go Memory Model",
		Content: "The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine.",
		PubTime: 1627724609,
		Link:    "https://google.com/imghp",
	},
	{
		Id:      2,
		Title:   "The Go Memory Model",
		Content: "The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine.",
		PubTime: 1627724609,
		Link:    "https://google.com/imghp",
	},
	{
		Id:      2,
		Title:   "The Go Memory Model",
		Content: "The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine.",
		PubTime: 1627724609,
		Link:    "https://google.com/imghp",
	},
	{
		Id:      2,
		Title:   "The Go Memory Model",
		Content: "The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine.",
		PubTime: 1627724609,
		Link:    "https://google.com/imghp",
	},
	{
		Id:      2,
		Title:   "The Go Memory Model",
		Content: "The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine.",
		PubTime: 1627724609,
		Link:    "https://google.com/imghp",
	},
	{
		Id:      2,
		Title:   "The Go Memory Model",
		Content: "The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine.",
		PubTime: 1627724609,
		Link:    "https://google.com/imghp",
	},
	{
		Id:      2,
		Title:   "The Go Memory Model",
		Content: "The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine.",
		PubTime: 1627724609,
		Link:    "https://google.com/imghp",
	},
	{
		Id:      2,
		Title:   "The Go Memory Model",
		Content: "The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine.",
		PubTime: 1627724609,
		Link:    "https://google.com/imghp",
	},
	{
		Id:      2,
		Title:   "The Go Memory Model",
		Content: "The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine.",
		PubTime: 1627724609,
		Link:    "https://google.com/imghp",
	},
	{
		Id:      2,
		Title:   "The Go Memory Model",
		Content: "The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine.",
		PubTime: 1627724609,
		Link:    "https://google.com/imghp",
	},
	{
		Id:      2,
		Title:   "The Go Memory Model",
		Content: "The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine.",
		PubTime: 1627724609,
		Link:    "https://google.com/imghp",
	},
	{
		Id:      2,
		Title:   "The Go Memory Model",
		Content: "The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine.",
		PubTime: 1627724609,
		Link:    "https://google.com/imghp",
	},
	{
		Id:      2,
		Title:   "The Go Memory Model",
		Content: "The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine.",
		PubTime: 1627724609,
		Link:    "https://google.com/imghp",
	},
	{
		Id:      2,
		Title:   "The Go Memory Model",
		Content: "The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine.",
		PubTime: 1627724609,
		Link:    "https://google.com/imghp",
	},
}
