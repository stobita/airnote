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
	URL string `json:"url"`
}

type listJSON struct {
	Items []interface{} `json:"items"`
}

func (p *presenter) ResponseLinks(o usecase.LinksOutputData) error {
	var j listJSON
	for _, v := range o {
		j.Items = append(j.Items, &linkJSON{
			URL: v.URL,
		})
	}
	return json.NewEncoder(p.writer).Encode(j)
}

func (p *presenter) ResponseLink(o usecase.LinkOutputData) error {
	j := linkJSON{
		URL: o.URL,
	}
	return json.NewEncoder(p.writer).Encode(j)
}
