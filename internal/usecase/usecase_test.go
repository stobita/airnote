package usecase_test

import (
	"errors"
	"testing"

	"github.com/stobita/airnote/internal/domain/model"
	"github.com/stobita/airnote/internal/usecase"
)

func TestInteractor_AddLink(t *testing.T) {
	repository := &fakeRepository{}
	presenter := &fakePresenter{}
	interactor := usecase.NewInteractor(repository, presenter)
	t.Run("Success", func(t *testing.T) {
		input := usecase.LinkInputData{
			URL:         "http://localhost",
			Description: "test link",
			Tags:        []string{"test1", "test2", "test3"},
		}
		interactor.AddLink(input)
	})
}

type fakeRepository struct {
	doneGetLink,
	doneGetLinks,
	doneSaveLink,
	doneUpdateLink,
	doneDeleteLink,
	doneGetTagByText,
	doneSaveTag bool

	errorGetLink,
	errorGetLinks,
	errorSaveLink,
	errorUpdateLink,
	errorDeleteLink,
	errorGetTagByText,
	errorSaveTag bool
}

func (r *fakeRepository) GetLink(id int) (*model.Link, error) {
	if r.errorGetLink {
		return nil, errors.New("fake error")
	}
	r.doneGetLink = true
	return nil, nil
}
func (r *fakeRepository) GetLinks() ([]*model.Link, error) {
	if r.errorGetLinks {
		return nil, errors.New("fake error")
	}
	r.doneGetLinks = true
	return nil, nil
}

func (r *fakeRepository) SaveLink(input *model.Link) error {
	if r.errorSaveLink {
		return errors.New("fake error")
	}
	r.doneSaveLink = true
	return nil
}
func (r *fakeRepository) UpdateLink(*model.Link) error {
	if r.errorUpdateLink {
		return errors.New("fake error")
	}
	r.doneUpdateLink = true
	return nil
}
func (r *fakeRepository) DeleteLink(*model.Link) error {
	if r.errorDeleteLink {
		return errors.New("fake error")
	}
	r.doneDeleteLink = true
	return nil
}
func (r *fakeRepository) GetTagByText(text string) (*model.Tag, error) {
	if r.errorGetTagByText {
		return nil, errors.New("fake error")
	}
	r.doneGetTagByText = true
	return nil, nil
}
func (r *fakeRepository) SaveTag(input *model.Tag) error {
	if r.errorSaveTag {
		return errors.New("fake error")
	}
	r.doneSaveTag = true
	return nil
}

type fakePresenter struct{}

func (r *fakePresenter) ResponseLink(o usecase.LinkOutputData) error   { return nil }
func (r *fakePresenter) ResponseLinks(o usecase.LinksOutputData) error { return nil }
func (r *fakePresenter) ResponseError(err error) error                 { return nil }
func (r *fakePresenter) ResponseNoContent() error                      { return nil }
