package repository

import (
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

func (r *repository) GetLinkTitle(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
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

func (r *repository) SaveLinkTitle(title string, linkID int) error {
	if err := r.redisClient.Set(strconv.Itoa(linkID), title, 0).Err(); err != nil {
		return err
	}
	return nil
}
