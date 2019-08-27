package usecase

// InputPort is usecase input port
type InputPort interface {
	GetAllLinks()
	AddLink(i LinkInputData)
	UpdateLink(id int, i LinkInputData)
	DeleteLink(id int)
}

// LinkInputData is used by InputPort
type LinkInputData struct {
	URL         string
	Description string
	Tags        []string
}

type LinkTagInputData struct {
	LinkID int
	Text   string
}
