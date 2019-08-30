package usecase

import "github.com/stobita/airnote/internal/domain/model"

type OutputPort interface {
	ResponseLink(o LinkOutputData) error
	ResponseLinks(o LinksOutputData) error
	ResponseLinkOriginal(o LinkOriginalOutputData) error

	ResponseTags(o TagsOutputData) error

	ResponseError(err error) error
	ResponseNoContent() error
}

// LinksOutputData is used by OutputPort
type LinksOutputData []*LinkOutputData

// LinkOutputData is used by OutputPort
type LinkOutputData struct {
	ID          int
	Title       string
	URL         string
	Description string
	Tags        []*TagOutputData
}

func makeLinksOutputData(links []*model.Link) LinksOutputData {
	var o LinksOutputData
	for _, v := range links {
		tagOutput := []*TagOutputData{}
		for _, v := range v.GetTags() {
			tagOutput = append(tagOutput, &TagOutputData{
				ID:   v.GetID(),
				Text: v.GetText(),
			})
		}
		o = append(o, &LinkOutputData{
			ID:          v.GetID(),
			Title:       v.GetTitle(),
			URL:         v.GetURL(),
			Description: v.GetDescription(),
			Tags:        tagOutput,
		})
	}
	return o
}

type TagsOutputData []*TagOutputData

type TagOutputData struct {
	ID   int
	Text string
}

type LinkOriginalOutputData struct {
	Title string
}
