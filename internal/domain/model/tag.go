package model

import "errors"

// Tag model of tag
type Tag struct {
	id   int
	text string
}

// TagInput input of tag model
type TagInput struct {
	Text string
}

// NewTag return new tag model
func NewTag(i TagInput) (*Tag, error) {
	if i.Text == "" {
		return nil, errors.New("Tag.text must set")
	}
	return &Tag{
		text: i.Text,
	}, nil
}

// GetID return tag id
func (t *Tag) GetID() int {
	return t.id
}

// GetText return tag text
func (t *Tag) GetText() string {
	return t.text
}

// SetID set id to tag model
func (t *Tag) SetID(id int) {
	t.id = id
}
