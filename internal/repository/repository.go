package repository

import (
	"context"
	"database/sql"

	"github.com/stobita/airnote/internal/domain/model"
	"github.com/stobita/airnote/internal/repository/rdb"
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

func (r *repository) SaveLink(input *model.Link) error {
	link := rdb.Link{
		URL: input.GetURL(),
	}
	if err := link.Insert(context.Background(), r.db, boil.Whitelist("url")); err != nil {
		return err
	}
	input.SetID(link.ID)
	return nil
}

func (r *repository) GetLinks() ([]*model.Link, error) {

	links, err := rdb.Links().All(context.Background(), r.db)
	if err != nil {
		return nil, err
	}
	var result []*model.Link
	for _, v := range links {
		input := model.LinkInput{
			URL: v.URL,
		}
		m := model.NewLink(input)
		m.SetID(v.ID)
		result = append(result, m)
	}
	return result, nil
}
