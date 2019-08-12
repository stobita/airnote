package usecase

import (
	"github.com/stobita/airnote/internal/domain/model"
)

type interactor struct {
	repository repository
	outputPort outputPort
}

// InputPort is usecase input port
type InputPort interface {
	GetAllLinks() error
	AddLink(i InputData) error
}

// InputData is used by InputPort
type InputData struct {
	URL string
}

type outputPort interface {
	ResponseLinks(o LinksOutputData) error
	ResponseLink(o LinkOutputData) error
}

// LinksOutputData is used by OutputPort
type LinksOutputData []*LinkOutputData

// LinkOutputData is used by OutputPort
type LinkOutputData struct {
	URL string
}

type repository interface {
	SaveLink(input *model.Link) error
	GetLinks() ([]*model.Link, error)
}

// NewInteractor get interactor
func NewInteractor(r repository, o outputPort) *interactor {
	return &interactor{
		repository: r,
		outputPort: o,
	}
}

func (i *interactor) AddLink(input InputData) error {
	model := model.NewLink(model.LinkInput{URL: input.URL})
	if err := i.repository.SaveLink(model); err != nil {
		return err
	}
	o := LinkOutputData{URL: model.GetURL()}
	return i.outputPort.ResponseLink(o)
}

func (i *interactor) GetAllLinks() error {
	links, err := i.repository.GetLinks()
	if err != nil {
		return err
	}
	var o LinksOutputData
	for _, v := range links {
		o = append(o, &LinkOutputData{
			URL: v.GetURL(),
		})
	}
	return i.outputPort.ResponseLinks(o)
}
