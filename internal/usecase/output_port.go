package usecase

type OutputPort interface {
	ResponseLink(o LinkOutputData) error
	ResponseLinks(o LinksOutputData) error
	ResponseTag(o TagOutputData) error

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
	Tags        []*TagOutputData
}

type TagOutputData struct {
	ID   int
	Text string
}
