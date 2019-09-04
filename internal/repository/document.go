package repository

import (
	"context"
	"strconv"

	"github.com/olivere/elastic/v7"
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

func (r *repository) SearchLinks(word string) ([]int, error) {
	ctx := context.Background()
	query := elastic.NewMultiMatchQuery(word, "title", "description")
	result, err := r.esClient.Search().
		Index(linkIndex).
		Query(query).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	var ids []int
	if result.Hits.TotalHits.Value > 0 {
		for _, hit := range result.Hits.Hits {
			id, err := strconv.Atoi(hit.Id)
			if err != nil {
				return nil, err
			}
			ids = append(ids, id)
		}
	}
	return ids, nil
}
