package postgres

import (
	"GoNews/pkg/storage"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

// The store which has the connection
type Store struct {
	db *pgxpool.Pool
}

var ErrPageCountExceeded = fmt.Errorf("page count exceeded")

// Constructor for Store
func New(connectionString string) (*Store, error) {
	db, err := pgxpool.Connect(context.Background(), connectionString)
	if err != nil {
		return nil, err
	}
	s := Store{
		db: db,
	}
	return &s, nil
}

func (s *Store) cleanDb() error {
	tx, err := s.db.Begin(context.Background())
	if err != nil {
		return err
	}

	_, err = tx.Exec(context.Background(),
		`
			begin;
			drop table if exists posts;
			
			create table posts(
				id bigserial primary key,
				title text not null,
				content text not null,
				pubtime bigint not null,
				link text not null,
				UNIQUE (title, link)
			);
			
			commit;
		`)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}

	tx.Commit(context.Background())

	return nil
}

// Gets a single piece of news by articleId
func (s *Store) NewsSingle(articleId int) (storage.Article, error) {
	fmt.Println("In news single")
	var article storage.Article
	row := s.db.QueryRow(context.Background(), `
		SELECT
			id,
			title,
			content,
			pubtime,
			link
		FROM posts
		WHERE id = $1
	`,
		articleId,
	)
	if err := row.Scan(
		&article.Id,
		&article.Title,
		&article.Content,
		&article.PubTime,
		&article.Link,
	); err != nil {
		return storage.Article{}, err
	}
	return article, nil
}

// Gets news where limit = number of entries
func (s *Store) News(page, pageSize int) (storage.PaginatedResult, error) {
	fmt.Println("In news")
	var result storage.PaginatedResult
	result.Page = page
	result.PageSize = pageSize
	rows, err := s.db.Query(context.Background(), `
		SELECT
			id,
			title,
			content,
			pubtime,
			link
		FROM posts
		ORDER BY id
		LIMIT $1
		OFFSET $2
	`,
		pageSize,
		page,
	)
	if err != nil {
		return storage.PaginatedResult{}, err
	}
	var posts []storage.Article
	for rows.Next() {
		var t storage.Article
		err = rows.Scan(
			&t.Id,
			&t.Title,
			&t.Content,
			&t.PubTime,
			&t.Link,
		)
		if err != nil {
			return storage.PaginatedResult{}, err
		}
		posts = append(posts, t)
	}
	result.Articles = posts
	var rowCount int
	r := s.db.QueryRow(context.Background(), `
	SELECT COUNT(*) from posts
	`)
	r.Scan(&rowCount)
	result.PageCount = rowCount / pageSize
	if page > rowCount/pageSize {
		return storage.PaginatedResult{}, ErrPageCountExceeded
	}
	if r := rowCount % pageSize; r > 0 { //handle pagination remainder
		if page == rowCount/pageSize {
			result.Articles = result.Articles[:r]
		}
	}
	return result, rows.Err()
}

// Finds news with specified string in the title
func (s *Store) FilterNews(query string) ([]storage.Article, error) {
	fmt.Println("In filter news")
	rows, err := s.db.Query(context.Background(), `
		SELECT 
			id,
			title,
			content,
			pubtime,
			link
		FROM posts
		WHERE LOWER(title) LIKE LOWER('%`+query+`%')
		ORDER BY id
	`,
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var posts []storage.Article
	for rows.Next() {
		var t storage.Article
		err = rows.Scan(
			&t.Id,
			&t.Title,
			&t.Content,
			&t.PubTime,
			&t.Link,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, t)

	}
	return posts, rows.Err()
}

// Adds multiple posts to the database
func (s *Store) AddNews(posts []storage.Article) error {
	fmt.Println("In add news")
	var err error
	tx, err := s.db.Begin(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, post := range posts {
		_, err = tx.Exec(context.Background(), `
		INSERT into posts(
			title,
			content,
			pubtime,
			link
		)
		values(
			$1,
			$2,
			$3,
			$4
		)
		ON CONFLICT DO NOTHING
		RETURNING id;
		`,
			post.Title,
			post.Content,
			post.PubTime,
			post.Link,
		)

		if err == nil {
			tx.Commit(context.Background())
		}
	}
	return err
}
