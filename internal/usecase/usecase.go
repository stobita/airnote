package usecase

import (
	"log"

	"github.com/stobita/airnote/internal/domain/model"
)

type interactor struct {
	repository repository
	outputPort OutputPort
}

type repository interface {
	linkRepository
}

type linkRepository interface {
	GetLink(id int) (*model.Link, error)
	GetLinks() ([]*model.Link, error)
	SaveLink(input *model.Link) error
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

func (i *interactor) AddLink(input LinkInputData) {
	tags := []*model.Tag{}
	for _, v := range input.Tags {
		tag, err := model.NewTag(model.TagInput{
			Text: v,
		})
		if err != nil {
			log.Print(err)
			i.outputPort.ResponseError(err)
			return
		}
		tags = append(tags, tag)
	}
	link, err := model.NewLink(model.LinkInput{
		URL:         input.URL,
		Description: input.Description,
		Tags:        tags,
	})
	if err != nil {
		log.Printf("NewLink error: %s", err)
		i.outputPort.ResponseError(err)
		return
	}
	if err := i.repository.SaveLink(link); err != nil {
		log.Printf("SaveLink error: %s", err)
		i.outputPort.ResponseError(err)
		return
	}
	tagOutput := []*TagOutputData{}
	for _, v := range link.GetTags() {
		tagOutput = append(tagOutput, &TagOutputData{
			ID:   v.GetID(),
			Text: v.GetText(),
		})
	}
	o := LinkOutputData{
		ID:          link.GetID(),
		URL:         link.GetURL(),
		Description: link.GetDescription(),
		Tags:        tagOutput,
	}
	if err := i.outputPort.ResponseLink(o); err != nil {
		log.Printf("ResponseLink error: %s", err)
		i.outputPort.ResponseError(err)
		return
	}
}

func (i *interactor) GetAllLinks() {
	links, err := i.repository.GetLinks()
	if err != nil {
		log.Printf("GetLinks error: %s", err)
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
		log.Print(err)
		i.outputPort.ResponseError(err)
		return
	}
}

func (i *interactor) UpdateLink(id int, input LinkInputData) {
	link, err := i.repository.GetLink(id)
	if err != nil {
		log.Printf("GetLink error: %s", err)
		i.outputPort.ResponseError(err)
		return
	}
	link.SetURL(input.URL)
	link.SetDescription(input.Description)
	tags := []*model.Tag{}
	for _, v := range input.Tags {
		tag, err := model.NewTag(model.TagInput{
			Text: v,
		})
		if err != nil {
			log.Print(err)
			i.outputPort.ResponseError(err)
			return
		}
		tags = append(tags, tag)
	}
	link.SetTags(tags)
	if err := i.repository.UpdateLink(link); err != nil {
		log.Print(err)
		i.outputPort.ResponseError(err)
		return
	}
	tagOutput := []*TagOutputData{}
	for _, v := range link.GetTags() {
		tagOutput = append(tagOutput, &TagOutputData{
			ID:   v.GetID(),
			Text: v.GetText(),
		})
	}
	o := LinkOutputData{
		ID:          link.GetID(),
		URL:         link.GetURL(),
		Description: link.GetDescription(),
		Tags:        tagOutput,
	}
	if err := i.outputPort.ResponseLink(o); err != nil {
		log.Print(err)
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
