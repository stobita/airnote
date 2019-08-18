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
	ID          int    `json:"id"`
	URL         string `json:"url"`
	Description string `json:"description"`
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
	var j listJSON
	for _, v := range o {
		j.Items = append(j.Items, &linkJSON{
			ID:          v.ID,
			URL:         v.URL,
			Description: v.Description,
		})
	}
	return json.NewEncoder(p.writer).Encode(j)
}

func (p *presenter) ResponseLink(o usecase.LinkOutputData) error {
	j := linkJSON{
		ID:          o.ID,
		URL:         o.URL,
		Description: o.Description,
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
