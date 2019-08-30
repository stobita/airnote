package usecase

import (
	"errors"
	"log"

	"github.com/stobita/airnote/internal/domain/model"
)

type interactor struct {
	repository repository
	outputPort OutputPort
}

type repository interface {
	linkRepository
	tagRepository
	ogpRepository
}

type linkRepository interface {
	GetLink(id int) (*model.Link, error)
	GetLinks() ([]*model.Link, error)
	GetLinksByTagID(tagID int) ([]*model.Link, error)
	SaveLink(input *model.Link) error
	UpdateLink(*model.Link) error
	DeleteLink(*model.Link) error
}

type tagRepository interface {
	GetTag(id int) (*model.Tag, error)
	GetTags() ([]*model.Tag, error)
	GetTagByText(text string) (*model.Tag, error)
	SaveTag(input *model.Tag) error
}

type ogpRepository interface {
	GetLinkTitle(url string) (string, error)
	SaveLinkTitle(title string, linkID int) error
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
		tag, err := i.repository.GetTagByText(v)
		if err != nil {
			log.Print(err)
			i.outputPort.ResponseError(err)
			return
		}
		if tag == nil {
			tag, err = model.NewTag(model.TagInput{Text: v})
			if err != nil {
				log.Print(err)
				i.outputPort.ResponseError(err)
				return
			}
			if err := i.repository.SaveTag(tag); err != nil {
				log.Print(err)
				i.outputPort.ResponseError(err)
				return
			}
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

	o := makeLinksOutputData(links)

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

	// TODO: fix duplicate AddLink
	tags := []*model.Tag{}
	for _, v := range input.Tags {
		tag, err := i.repository.GetTagByText(v)
		if err != nil {
			log.Print(err)
			i.outputPort.ResponseError(err)
			return
		}
		if tag == nil {
			tag, err = model.NewTag(model.TagInput{Text: v})
			if err != nil {
				log.Print(err)
				i.outputPort.ResponseError(err)
				return
			}
			if err := i.repository.SaveTag(tag); err != nil {
				log.Print(err)
				i.outputPort.ResponseError(err)
				return
			}
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
	link, err := i.repository.GetLink(id)
	if err != nil {
		i.outputPort.ResponseError(err)
		return
	}
	if err := i.repository.DeleteLink(link); err != nil {
		i.outputPort.ResponseError(err)
		return
	}
	if err := i.outputPort.ResponseNoContent(); err != nil {
		i.outputPort.ResponseError(err)
		return
	}
}

func (i *interactor) GetLinkOriginal(id int) {
	link, err := i.repository.GetLink(id)
	if err != nil {
		log.Printf("GetLink error: %s", err)
		i.outputPort.ResponseError(err)
		return
	}
	title, err := i.repository.GetLinkTitle(link.GetURL())
	if err != nil {
		log.Printf("GetLinkTitle error: %s", err)
		i.outputPort.ResponseError(err)
		return
	}
	if err := i.repository.SaveLinkTitle(title, link.GetID()); err != nil {
		log.Printf("GetLinkTitle error: %s", err)
		i.outputPort.ResponseError(err)
		return
	}
	o := LinkOriginalOutputData{
		Title: title,
	}
	if err := i.outputPort.ResponseLinkOriginal(o); err != nil {
		log.Print(err)
		i.outputPort.ResponseError(err)
		return
	}

}

func (i *interactor) GetTaggedLinks(tagID int) {
	tag, err := i.repository.GetTag(tagID)
	if err != nil {
		log.Printf("GetTag error: %s", err)
		i.outputPort.ResponseError(err)
		return
	}
	if tag == nil {
		i.outputPort.ResponseError(errors.New("Invalid tag id"))
		return
	}
	links, err := i.repository.GetLinksByTagID(tag.GetID())
	if err != nil {
		log.Printf("GetLinksByTagID error: %s", err)
		i.outputPort.ResponseError(err)
		return
	}
	o := makeLinksOutputData(links)
	if err := i.outputPort.ResponseLinks(o); err != nil {
		log.Print(err)
		i.outputPort.ResponseError(err)
		return
	}
}

func (i *interactor) GetAllTags() {
	tags, err := i.repository.GetTags()
	if err != nil {
		log.Printf("GetTags error: %s", err)
		i.outputPort.ResponseError(err)
		return
	}
	var o TagsOutputData
	for _, v := range tags {
		o = append(o, &TagOutputData{
			ID:   v.GetID(),
			Text: v.GetText(),
		})
	}
	if err := i.outputPort.ResponseTags(o); err != nil {
		log.Print(err)
		i.outputPort.ResponseError(err)
		return
	}
}
