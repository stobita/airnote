package model

import "errors"

type Link struct {
	id          int
	url         string
	description string
}

type LinkInput struct {
	URL         string
	Description string
}

func NewLink(i LinkInput) (*Link, error) {
	if i.URL == "" {
		return nil, errors.New("URL must set")
	}
	return &Link{
		url:         i.URL,
		description: i.Description,
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

func (l *Link) SetID(id int) {
	l.id = id
}
