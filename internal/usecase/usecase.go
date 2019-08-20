package usecase

import (
	"github.com/stobita/airnote/internal/domain/model"
)

type interactor struct {
	repository repository
	outputPort OutputPort
}

// InputPort is usecase input port
type InputPort interface {
	GetAllLinks()
	AddLink(i InputData)
	UpdateLink(id int, i InputData)
	DeleteLink(id int)
}

// InputData is used by InputPort
type InputData struct {
	URL         string
	Description string
}

type OutputPort interface {
	ResponseLink(o LinkOutputData) error
	ResponseLinks(o LinksOutputData) error
	ResponseError(err error) error
	ResponseNoContent() error
}

// LinksOutputData is used by OutputPort
type LinksOutputData []*LinkOutputData

// LinkOutputData is used by OutputPort
type LinkOutputData struct {
	ID          int
	URL         string
	Description string
}

type repository interface {
	GetLink(id int) (*model.Link, error)
	GetLinks() ([]*model.Link, error)
	CreateLink(input *model.Link) error
	UpdateLink(*model.Link) error
	DeleteLink(*model.Link) error
}

// NewInteractor get interactor
func NewInteractor(r repository, o OutputPort) *interactor {
	return &interactor{
		repository: r,
		outputPort: o,
	}
}

func (i *interactor) AddLink(input InputData) {
	model, err := model.NewLink(model.LinkInput{
		URL:         input.URL,
		Description: input.Description,
	})
	if err != nil {
		i.outputPort.ResponseError(err)
		return
	}
	if err := i.repository.CreateLink(model); err != nil {
		i.outputPort.ResponseError(err)
		return
	}
	o := LinkOutputData{
		ID:          model.GetID(),
		URL:         model.GetURL(),
		Description: model.GetDescription(),
	}
	if err := i.outputPort.ResponseLink(o); err != nil {
		i.outputPort.ResponseError(err)
		return
	}
}

func (i *interactor) GetAllLinks() {
	links, err := i.repository.GetLinks()
	if err != nil {
		i.outputPort.ResponseError(err)
		return
	}
	var o LinksOutputData
	for _, v := range links {
		o = append(o, &LinkOutputData{
			ID:          v.GetID(),
			URL:         v.GetURL(),
			Description: v.GetDescription(),
		})
	}
	if err := i.outputPort.ResponseLinks(o); err != nil {
		i.outputPort.ResponseError(err)
		return
	}
}

func (i *interactor) UpdateLink(id int, input InputData) {
	model, err := i.repository.GetLink(id)
	if err != nil {
		i.outputPort.ResponseError(err)
		return
	}
	model.SetURL(input.URL)
	model.SetDescription(input.Description)
	if err := i.repository.UpdateLink(model); err != nil {
		i.outputPort.ResponseError(err)
		return
	}
	o := LinkOutputData{
		ID:          model.GetID(),
		URL:         model.GetURL(),
		Description: model.GetDescription(),
	}
	if err := i.outputPort.ResponseLink(o); err != nil {
		i.outputPort.ResponseError(err)
		return
	}
}

func (i *interactor) DeleteLink(id int) {
	model, err := i.repository.GetLink(id)
	if err != nil {
		i.outputPort.ResponseError(err)
		return
	}
	if err := i.repository.DeleteLink(model); err != nil {
		i.outputPort.ResponseError(err)
		return
	}
	if err := i.outputPort.ResponseNoContent(); err != nil {
		i.outputPort.ResponseError(err)
		return
	}
}
