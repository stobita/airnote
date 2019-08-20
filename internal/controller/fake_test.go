package controller_test

import "github.com/stobita/airnote/internal/usecase"

type fakeInputPort struct {
	getAllLinksSuccess bool
	getAllLinksError   bool
	getAddLinkSuccess  bool
	getAddLinkError    bool
}

func (f *fakeInputPort) GetAllLinks() {
	if f.getAllLinksError {
		return
	}
	f.getAllLinksSuccess = true
	return
}
func (f *fakeInputPort) AddLink(i usecase.InputData) {
	if f.getAddLinkError {
		return
	}
	f.getAddLinkSuccess = true
	return
}

func (f *fakeInputPort) UpdateLink(id int, i usecase.InputData) {
	return
}

func (f *fakeInputPort) DeleteLink(id int) {
	return
}

type fakeOutputPort struct{}

func (f *fakeOutputPort) ResponseLink(o usecase.LinkOutputData) error {
	return nil
}

func (f *fakeOutputPort) ResponseLinks(o usecase.LinksOutputData) error {
	return nil
}

func (f *fakeOutputPort) ResponseError(e error) error {
	return nil
}

func (f *fakeOutputPort) ResponseNoContent() error {
	return nil
}
