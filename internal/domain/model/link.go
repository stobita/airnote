package model

type Link struct {
	id  int
	url string
}

type LinkInput struct {
	URL string
}

func NewLink(i LinkInput) *Link {
	return &Link{
		url: i.URL,
	}
}

func (l *Link) GetID() int {
	return l.id
}

func (l *Link) GetURL() string {
	return l.url
}

func (l *Link) SetID(id int) {
	l.id = id
}
