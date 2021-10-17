package postgres

import (
	"GoNews/pkg/storage"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		connectionString string
	}
	tests := []struct {
		name    string
		args    args
		want    *Store
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.connectionString)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_NewsSingle(t *testing.T) {
	type args struct {
		articleId int
	}
	tests := []struct {
		name    string
		s       *Store
		args    args
		want    storage.Article
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.NewsSingle(tt.args.articleId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.NewsSingle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.NewsSingle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_News(t *testing.T) {
	type args struct {
		page     int
		pageSize int
	}
	tests := []struct {
		name    string
		s       *Store
		args    args
		want    storage.PaginatedResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.News(tt.args.page, tt.args.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.News() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.News() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_FilterNews(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		s       *Store
		args    args
		want    []storage.Article
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.FilterNews(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.FilterNews() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.FilterNews() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_AddNews(t *testing.T) {
	type args struct {
		posts []storage.Article
	}
	tests := []struct {
		name    string
		s       *Store
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.AddNews(tt.args.posts); (err != nil) != tt.wantErr {
				t.Errorf("Store.AddNews() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
