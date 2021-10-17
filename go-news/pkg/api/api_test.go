package api

import (
	"GoNews/pkg/storage/memdb"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

var url string = "https://habr.com/ru/rss/hub/go/all/?fl=ru"

func TestAPI_defaultHandler(t *testing.T) {

	db := memdb.New()
	api := New(db)
	req, err := http.NewRequest("GET", "/news/10", nil)
	if err != nil {
		t.Fatal(err)
	}

	//1
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.NewsHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	contains := `Думаю, многие из вас сталкивались с замысловатыми задачками`
	if strings.Contains(rr.Body.String(), contains) {
		t.Errorf("handler does not contain searched string: got %v want %v",
			rr.Body.String(), contains)
	}

	//2
	req, err = http.NewRequest("GET", "/news", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	contains = `Думаю, многие из вас сталкивались с замысловатыми задачками`
	if strings.Contains(rr.Body.String(), contains) {
		t.Errorf("handler does not contain searched string: got %v want %v",
			rr.Body.String(), contains)
	}
}

func TestAPI_DefaultHandler(t *testing.T) {
	db := memdb.New()
	api := New(db)
	req, err := http.NewRequest("GET", "/version", nil)
	if err != nil {
		t.Fatal(err)
	}

	//1
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.NewsHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	contains := `{
		"Api": "gonews",
		"Version": 1,
		"Message": "bitch"
	}`
	if strings.Contains(rr.Body.String(), contains) {
		t.Errorf("handler does not contain searched string: got %v want %v",
			rr.Body.String(), contains)
	}
}

func TestAPI_Router(t *testing.T) {
	db := memdb.New()
	api := New(db)
	tests := []struct {
		name string
		api  *API
		want *mux.Router
	}{
		{
			name: "default",
			api:  api,
			want: &mux.Router{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.api.Router(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("API.Router() = %v, want %v", got, tt.want)
			}
		})
	}
}
