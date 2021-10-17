package postgres

import (
	"context"
	"go-comments/pkg/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

// The store which has the connection
type Store struct {
	db *pgxpool.Pool
}

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

// Gets comments where limit = number of entries
func (s *Store) Comments(articleId int) ([]storage.Comment, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT
			id,
			parentId,
			articleId,
			pubtime,
			content,
			profane
		FROM comments
		WHERE articleId = $1
		ORDER BY id
	`,
		articleId,
	)
	if err != nil {
		return nil, err
	}
	var comments []storage.Comment
	for rows.Next() {
		var t storage.Comment
		err = rows.Scan(
			&t.Id,
			&t.ParentId,
			&t.ArticleId,
			&t.PubTime,
			&t.Content,
			&t.Profane,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, t)
	}
	return comments, rows.Err()
}

// Adds a comment to the database
func (s *Store) AddComment(comment storage.Comment) error {
	var err error
	_, err = s.db.Exec(context.Background(),
		`
		INSERT into comments(
			parentId,
			articleId,
			content,
			pubtime,
			profane
		)
		values(
			$1,
			$2,
			$3,
			$4,
			$5
		) RETURNING id;
		`,
		comment.ParentId,
		comment.ArticleId,
		comment.Content,
		comment.PubTime,
		comment.Profane,
	)
	return err
}
