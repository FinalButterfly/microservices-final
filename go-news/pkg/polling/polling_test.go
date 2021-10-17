package polling

import (
	"GoNews/pkg/storage"
	"GoNews/pkg/storage/memdb"
	"reflect"
	"testing"
	"time"
)

var feeds []string = []string{
	"https://habr.com/ru/rss/hub/go/all/?fl=ru",
	"https://habr.com/ru/rss/best/daily/?fl=ru",
	"https://cprss.s3.amazonaws.com/golangweekly.com.xml",
}

func TestNewPoller(t *testing.T) {
	db := memdb.New()
	type args struct {
		c  Config
		db storage.Interface
	}
	tests := []struct {
		name    string
		args    args
		want    *Poller
		wantErr bool
	}{
		{
			name: "default",
			args: args{
				c: Config{
					Feeds:    feeds,
					Interval: 1,
				},
				db: db,
			},
			want: &Poller{
				Interval: time.Minute,
				Feeds:    feeds,
				db:       db,
			},
		},
		{
			name: "interval = 0",
			args: args{
				c: Config{
					Feeds:    feeds,
					Interval: 0,
				},
				db: db,
			},
			want: &Poller{
				Interval: time.Minute * 5,
				Feeds:    feeds,
				db:       db,
			},
		},
		{
			name: "len(feeds) == 0",
			args: args{
				c: Config{
					Feeds:    []string{},
					Interval: 5,
				},
				db: db,
			},
			want: &Poller{
				Interval: time.Minute * 5,
				Feeds:    []string{},
				db:       db,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPoller(tt.args.c, tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPoller() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPoller() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Каким образом тестировать горутины в тесте, которые находятся в цикле?
func TestPoller_StartPolling(t *testing.T) {
	db := memdb.New()
	tests := []struct {
		name    string
		p       *Poller
		wantErr bool
	}{
		{
			name: "default",
			p: &Poller{
				Interval: time.Minute,
				Feeds:    feeds,
				db:       db,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.StartPolling(); (err != nil) != tt.wantErr {
				t.Errorf("Poller.StartPolling() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPoller_getPosts(t *testing.T) {
	db := memdb.New()
	type args struct {
		url string
	}
	tests := []struct {
		name string
		p    *Poller
		args args
	}{
		{
			name: "default",
			p: &Poller{
				Interval: time.Minute,
				Feeds:    feeds,
				db:       db,
			},
			args: args{
				url: feeds[0],
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.getPosts(tt.args.url)
		})
	}
}
