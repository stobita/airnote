package model

import (
	"fmt"
	"net/url"

	"github.com/pkg/errors"
)

type Link struct {
	id          int
	url         string
	title       string
	description string
	tags        []*Tag
}

type LinkInput struct {
	URL         string
	Title       string
	Description string
	Tags        []*Tag
}

func NewLink(i LinkInput) (*Link, error) {
	if err := validateURL(i.URL); err != nil {
		return nil, err
	}

	return &Link{
		url:         i.URL,
		title:       i.Title,
		description: i.Description,
		tags:        i.Tags,
	}, nil
}

func (l *Link) GetID() int {
	return l.id
}

func (l *Link) GetURL() string {
	return l.url
}

func (l *Link) GetTitle() string {
	return l.title
}

func (l *Link) GetDescription() string {
	return l.description
}

func (l *Link) GetTags() []*Tag {
	return l.tags
}

func (l *Link) SetID(id int) {
	l.id = id
}

func (l *Link) SetURL(url string) error {
	if err := validateURL(url); err != nil {
		return err
	}
	l.url = url
	return nil
}

func (l *Link) SetTitle(title string) {
	l.title = title
}

func (l *Link) SetDescription(d string) {
	l.description = d
}

func (l *Link) SetTags(tags []*Tag) {
	l.tags = tags
}

func validateURL(u string) error {
	if u == "" {
		return errors.New("URL must set")
	}
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Invalid url: %s", u))
	}
	return nil
}
