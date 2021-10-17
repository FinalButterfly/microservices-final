package polling

import (
	"GoNews/pkg/storage"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Gets thrown if the feeds array in Config is empty
var ErrorEmptyFeedsArray error = fmt.Errorf("links array is empty")

// XML Feed
type Feed struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

// XML Channel
type Channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	Items       []Item `xml:"item"`
}

// XML Item
type Item struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	Link        string `xml:"link"`
}

// Poller
type Poller struct {
	Interval time.Duration
	Feeds    []string
	db       storage.Interface
}

// Config for config.json
type Config struct {
	Feeds    []string `json:"rss"`
	Interval int      `json:"request_period"`
}

type CInterface interface {
	StartPolling()
}

var postsChan chan []storage.Article = make(chan []storage.Article)

// Creates a poller with given config
func NewPoller(c Config, db storage.Interface) (*Poller, error) {
	p := Poller{}
	var err error
	if c.Interval == 0 {
		p.Interval = time.Minute * 5
	} else {
		//check config.json
		p.Interval = time.Duration(c.Interval) * time.Minute
	}
	if len(c.Feeds) == 0 {
		err = ErrorEmptyFeedsArray
	}
	p.Feeds = c.Feeds
	p.db = db
	return &p, err
}

// Launches the RSS polling mechanism, writes to DB on success
func (p *Poller) StartPolling() error {
	if len(p.Feeds) == 0 {
		return ErrorEmptyFeedsArray
	}
	ticker := time.NewTicker(p.Interval)

	for {
		select {
		case _ = <-ticker.C:
			for i := 0; i < len(p.Feeds); i++ {
				go p.getPosts(p.Feeds[i])
			}
		case result := <-postsChan:
			err := p.db.AddNews(result)
			if err != nil {
				fmt.Println("Failed to add posts to database")
				return err
			}
		}
	}
}

func (p *Poller) getPosts(url string) {
	var posts []storage.Article
	response, err := http.Get(url)
	if err == nil {
		body, err := ioutil.ReadAll(response.Body)
		if err == nil {
			var f Feed
			err := xml.Unmarshal(body, &f)
			if err == nil {
				for _, item := range f.Channel.Items {
					var p storage.Article
					p.Content = item.Description
					p.Link = item.Link
					t, err := time.Parse("Fri, 23 Jul 2021 00:00:00 +0000", item.PubDate) //Не работает
					if err != nil {
						p.PubTime = 0
					} else {
						p.PubTime = t.UnixNano()
					}
					p.Title = item.Title
					posts = append(posts, p)
				}
				postsChan <- posts
			}
		}
	}
	return
}
