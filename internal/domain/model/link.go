package model

import "errors"

type Link struct {
	id          int
	url         string
	description string
	tags        []*Tag
}

type LinkInput struct {
	URL         string
	Description string
	Tags        []*Tag
}

func NewLink(i LinkInput) (*Link, error) {
	if i.URL == "" {
		return nil, errors.New("URL must set")
	}
	return &Link{
		url:         i.URL,
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

func (l *Link) GetDescription() string {
	return l.description
}

func (l *Link) GetTags() []*Tag {
	return l.tags
}

func (l *Link) SetID(id int) {
	l.id = id
}

func (l *Link) SetURL(url string) {
	l.url = url
}

func (l *Link) SetDescription(d string) {
	l.description = d
}

func (l *Link) SetTags(tags []*Tag) {
	l.tags = tags
}
