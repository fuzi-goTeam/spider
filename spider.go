package spider

import (
	"errors"
	"net/http"

	iconv "gopkg.in/iconv.v1"

	"github.com/PuerkitoBio/goquery"
)

// URL type represent url
type URL string

// Spider interface 定义了 爬虫需要的方法
type Spider interface {
	GetInformation(*goquery.Document) interface{}
	Next(interface{}) URL
}

// Spiderinfo struct represent spider information
type Spiderinfo struct {
	Encode  string
	Website URL
	Channal chan interface{}
	Spider  *struct{}
}

// New generate a Spiderinfo pointer
func New(website URL, encode string, concurrentNum int, spider *struct{}) (s *Spiderinfo, err error) {
	if website == "" {
		err = errors.New("Spider func new website can not be empty")
		return
	}
	if encode == "" {
		encode = "utf-8"
	}

	s.Website = website
	s.Encode = encode
	s.Channal = make(chan interface{}, concurrentNum)

	return
}

// GetHTML get html content
func (s *Spiderinfo) GetHTML() (doc *goquery.Document, err error) {
	resp, err := http.Get(string(s.Website))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	cd, err := iconv.Open(s.Encode, "utf-8")
	if err != nil {
		return
	}
	defer cd.Close()

	r := iconv.NewReader(cd, resp.Body, 0)
	doc, err = goquery.NewDocumentFromReader(r)

	return
}
