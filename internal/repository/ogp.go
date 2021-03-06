package repository

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"golang.org/x/net/html"
)

func (r *repository) GetLinkTitle(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	log.Printf("fetch link: %s", url)
	res, err := r.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if !strings.HasPrefix(res.Header.Get("Content-Type"), "text/html") {
		return "", errors.New("Content type must be text/html")
	}
	title, err := findTitle(res.Body)
	if err != nil {
		return "", err
	}
	return title, nil
}

func findTitle(body io.Reader) (string, error) {
	t := html.NewTokenizer(body)
	loop := true
	depth := 0
	title := ""
	for loop {
		tt := t.Next()
		switch tt {
		case html.ErrorToken:
			loop = false
		case html.TextToken:
			if depth > 0 {
				title = string(t.Text())
				loop = false
			}
		case html.StartTagToken, html.EndTagToken:
			tn, _ := t.TagName()
			if string(tn) == "title" {
				if tt == html.StartTagToken {
					depth++
				} else {
					depth--
				}
			}
		}
	}
	return title, nil
}
