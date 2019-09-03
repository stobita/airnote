package repository

import (
	"context"
	"strconv"

	"github.com/stobita/airnote/internal/domain/model"
)

type linkDocument struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

const linkIndex = "link"

func (r *repository) SaveLinkDocument(input *model.Link) error {
	ctx := context.Background()
	doc := linkDocument{
		Title:       input.GetTitle(),
		Description: input.GetDescription(),
	}
	_, err := r.esClient.Index().Index(linkIndex).
		Id(strconv.Itoa(input.GetID())).
		BodyJson(doc).
		Do(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateLinkDocument(input *model.Link) error {
	ctx := context.Background()
	doc := linkDocument{
		Title:       input.GetTitle(),
		Description: input.GetDescription(),
	}
	_, err := r.esClient.Update().Index(linkIndex).
		Id(strconv.Itoa(input.GetID())).
		Doc(doc).
		Do(ctx)
	if err != nil {
		return err
	}
	return nil

}

func (r *repository) DeleteLinkDocument(model *model.Link) error {
	ctx := context.Background()
	_, err := r.esClient.Delete().Index(linkIndex).
		Id(strconv.Itoa(model.GetID())).
		Do(ctx)
	if err != nil {
		return err
	}
	return nil
}
