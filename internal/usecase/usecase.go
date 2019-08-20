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
	GetAllLinks() error
	AddLink(i InputData) error
	UpdateLink(id int, i InputData) error
	DeleteLink(id int) error
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

func (i *interactor) AddLink(input InputData) error {
	model, err := model.NewLink(model.LinkInput{
		URL:         input.URL,
		Description: input.Description,
	})
	if err != nil {
		return err
	}
	if err := i.repository.CreateLink(model); err != nil {
		return err
	}
	o := LinkOutputData{
		ID:          model.GetID(),
		URL:         model.GetURL(),
		Description: model.GetDescription(),
	}
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
			ID:          v.GetID(),
			URL:         v.GetURL(),
			Description: v.GetDescription(),
		})
	}
	return i.outputPort.ResponseLinks(o)
}

func (i *interactor) UpdateLink(id int, input InputData) error {
	model, err := i.repository.GetLink(id)
	if err != nil {
		return err
	}
	model.SetURL(input.URL)
	model.SetDescription(input.Description)
	if err := i.repository.UpdateLink(model); err != nil {
		return err
	}
	o := LinkOutputData{
		ID:          model.GetID(),
		URL:         model.GetURL(),
		Description: model.GetDescription(),
	}
	return i.outputPort.ResponseLink(o)
}

func (i *interactor) DeleteLink(id int) error {
	model, err := i.repository.GetLink(id)
	if err != nil {
		return err
	}
	if err := i.repository.DeleteLink(model); err != nil {
		return err
	}
	return i.outputPort.ResponseNoContent()
}
