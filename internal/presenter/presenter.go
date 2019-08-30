package presenter

import (
	"encoding/json"
	"net/http"

	"github.com/stobita/airnote/internal/usecase"
)

type presenter struct {
	writer http.ResponseWriter
}

// New create presenter
func New(w http.ResponseWriter) *presenter {
	return &presenter{w}
}

type linkJSON struct {
	ID          int       `json:"id"`
	URL         string    `json:"url"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Tags        []tagJSON `json:"tags"`
}

type tagJSON struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type linkOriginalJSON struct {
	Title string `json:"title"`
}

type listJSON struct {
	Items []interface{} `json:"items"`
}

type errorJSON struct {
	Error Error `json:"errors"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (p *presenter) ResponseLinks(o usecase.LinksOutputData) error {
	j := listJSON{Items: []interface{}{}}
	for _, v := range o {
		tagListJSON := []tagJSON{}
		for _, v := range v.Tags {
			tagListJSON = append(tagListJSON, tagJSON{
				ID:   v.ID,
				Text: v.Text,
			})
		}
		j.Items = append(j.Items, &linkJSON{
			ID:          v.ID,
			URL:         v.URL,
			Title:       v.Title,
			Description: v.Description,
			Tags:        tagListJSON,
		})
	}
	return json.NewEncoder(p.writer).Encode(j)
}

func (p *presenter) ResponseLink(o usecase.LinkOutputData) error {
	tags := []tagJSON{}
	for _, v := range o.Tags {
		tags = append(tags, tagJSON{
			Text: v.Text,
		})
	}
	j := linkJSON{
		ID:          o.ID,
		URL:         o.URL,
		Title:       o.Title,
		Description: o.Description,
		Tags:        tags,
	}
	return json.NewEncoder(p.writer).Encode(j)
}

func (p *presenter) ResponseLinkOriginal(o usecase.LinkOriginalOutputData) error {
	j := linkOriginalJSON{
		Title: o.Title,
	}
	return json.NewEncoder(p.writer).Encode(j)
}

func (p *presenter) ResponseError(e error) error {
	j := errorJSON{
		Error: Error{
			Message: e.Error(),
		},
	}
	return json.NewEncoder(p.writer).Encode(j)
}

func (p *presenter) ResponseNoContent() error {
	// TODO: implement
	return nil
}
