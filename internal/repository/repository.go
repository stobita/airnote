package repository

import (
	"context"
	"database/sql"

	"github.com/stobita/airnote/internal/domain/model"
	"github.com/stobita/airnote/internal/repository/rdb"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
)

type repository struct {
	db *sql.DB
}

// New return new repository
func New(db *sql.DB) *repository {
	return &repository{
		db,
	}
}

func (r *repository) GetLink(id int) (*model.Link, error) {
	link, err := rdb.FindLink(context.Background(), r.db, id)
	if err != nil {
		return nil, err
	}
	model, err := model.NewLink(model.LinkInput{
		URL:         link.URL,
		Description: link.Description.String,
	})
	model.SetID(link.ID)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (r *repository) GetLinks() ([]*model.Link, error) {
	links, err := rdb.Links().All(context.Background(), r.db)
	if err != nil {
		return nil, err
	}
	var result []*model.Link
	for _, v := range links {
		input := model.LinkInput{
			URL:         v.URL,
			Description: v.Description.String,
		}
		m, err := model.NewLink(input)
		if err != nil {
			return nil, err
		}
		m.SetID(v.ID)
		result = append(result, m)
	}
	return result, nil
}

func (r *repository) CreateLink(input *model.Link) error {
	link := rdb.Link{
		URL:         input.GetURL(),
		Description: null.StringFrom(input.GetDescription()),
	}
	if err := link.Insert(context.Background(), r.db, boil.Whitelist(
		"url",
		"description",
	)); err != nil {
		return err
	}
	input.SetID(link.ID)
	return nil
}

func (r *repository) UpdateLink(model *model.Link) error {
	link, err := rdb.FindLink(context.Background(), r.db, model.GetID())
	if err != nil {
		return err
	}
	link.URL = model.GetURL()
	link.Description = null.StringFrom(model.GetDescription())
	if _, err := link.Update(context.Background(), r.db, boil.Whitelist(
		"url",
		"description",
	)); err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteLink(model *model.Link) error {
	link, err := rdb.FindLink(context.Background(), r.db, model.GetID())
	if err != nil {
		return err
	}
	if _, err := link.Delete(context.Background(), r.db); err != nil {
		return err
	}
	model = nil
	return nil
}
