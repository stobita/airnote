package repository

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/stobita/airnote/internal/domain/model"
	"github.com/stobita/airnote/internal/repository/rdb"
	"github.com/stobita/airnote/internal/util"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func (r *repository) GetLink(id int) (*model.Link, error) {
	ctx := context.Background()
	// link, err := rdb.FindLink(ctx, r.db, id)
	link, err := rdb.Links(
		rdb.LinkWhere.ID.EQ(id),
		qm.Load(
			qm.Rels(
				rdb.LinkRels.LinksTags,
				rdb.LinksTagRels.Tag,
			),
		),
		qm.Load(rdb.LinkRels.LinkOriginal),
	).One(ctx, r.db)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	title := ""
	if link.R.LinkOriginal != nil {
		title = link.R.LinkOriginal.Title.String
	}
	m, err := model.NewLink(model.LinkInput{
		URL:         link.URL,
		Title:       title,
		Description: link.Description.String,
	})
	m.SetID(link.ID)
	if err != nil {
		return nil, err
	}
	if len(link.R.LinksTags) == 0 {
		return m, nil
	}
	tags := []*model.Tag{}
	for _, v := range link.R.LinksTags {
		tag, err := model.NewTag(model.TagInput{
			Text: v.R.Tag.Text,
		})
		if err != nil {
			return nil, err
		}
		tag.SetID(v.R.Tag.ID)
		tags = append(tags, tag)
	}
	return m, nil
}

func (r *repository) GetLinks() ([]*model.Link, error) {
	ctx := context.Background()
	links, err := rdb.Links(
		qm.Load(
			qm.Rels(
				rdb.LinkRels.LinksTags,
				rdb.LinksTagRels.Tag,
			),
		),
		qm.Load(
			rdb.LinkRels.LinkOriginal,
		),
	).All(ctx, r.db)
	if err != nil {
		return nil, errors.Wrap(err, "get All error")
	}
	var result []*model.Link
	for _, link := range links {
		var tags []*model.Tag
		for _, v := range link.R.LinksTags {
			tag, err := model.NewTag(model.TagInput{
				Text: v.R.Tag.Text,
			})
			if err != nil {
				return nil, err
			}
			tag.SetID(v.R.Tag.ID)
			tags = append(tags, tag)
		}

		title := ""
		if link.R.LinkOriginal != nil {
			title = link.R.LinkOriginal.Title.String
		}

		input := model.LinkInput{
			URL:         link.URL,
			Title:       title,
			Description: link.Description.String,
			Tags:        tags,
		}
		m, err := model.NewLink(input)
		if err != nil {
			return nil, err
		}
		m.SetID(link.ID)
		result = append(result, m)
	}
	return result, nil
}

func (r *repository) SaveLink(input *model.Link) error {
	ctx := context.Background()
	link := rdb.Link{
		URL:         input.GetURL(),
		Description: null.StringFrom(input.GetDescription()),
	}
	if err := link.Insert(ctx, r.db, boil.Whitelist(
		"url",
		"description",
	)); err != nil {
		return errors.Wrap(err, "Link.Insert error")
	}
	input.SetID(link.ID)

	if len(input.GetTags()) == 0 {
		return nil
	}

	rels := []*rdb.LinksTag{}
	for _, v := range input.GetTags() {
		i := &rdb.LinksTag{
			TagID: v.GetID(),
		}
		rels = append(rels, i)
	}
	if err := link.AddLinksTags(ctx, r.db, true, rels...); err != nil {
		return errors.Wrap(err, "link.AddLinksTags error")
	}

	return nil
}

func (r *repository) SaveTag(tag *model.Tag) error {
	ctx := context.Background()
	dbTag := rdb.Tag{
		Text: tag.GetText(),
	}
	if err := dbTag.Insert(ctx, r.db, boil.Whitelist("text")); err != nil {
		return errors.Wrap(err, "Insert error")
	}
	tag.SetID(dbTag.ID)
	return nil
}

func (r *repository) GetTagByText(text string) (*model.Tag, error) {
	ctx := context.Background()
	tag, err := rdb.Tags(rdb.TagWhere.Text.EQ(text)).One(ctx, r.db)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "Failed One")
	}
	model, err := model.NewTag(model.TagInput{Text: tag.Text})
	if err != nil {
		return nil, errors.Wrap(err, "NewTag error")
	}
	model.SetID(tag.ID)
	return model, nil
}

func (r *repository) UpdateLink(input *model.Link) error {
	ctx := context.Background()
	// dbLink, err := rdb.FindLink(ctx, r.db, input.GetID())
	dbLink, err := rdb.Links(rdb.LinkWhere.ID.EQ(input.GetID()), qm.Load(
		qm.Rels(
			rdb.LinkRels.LinksTags,
			rdb.LinksTagRels.Tag,
		),
	)).One(ctx, r.db)
	if err != nil {
		return err
	}
	dbLink.URL = input.GetURL()
	dbLink.Description = null.StringFrom(input.GetDescription())
	if _, err := dbLink.Update(ctx, r.db, boil.Whitelist(
		"url",
		"description",
	)); err != nil {
		return err
	}

	var beforeTagText []string
	for _, v := range dbLink.R.LinksTags {
		beforeTagText = append(beforeTagText, v.R.Tag.Text)
	}

	var afterTagText []string
	for _, v := range input.GetTags() {
		afterTagText = append(afterTagText, v.GetText())
	}

	tagDiff := util.StringArrayDiff(beforeTagText, afterTagText)

	addTagIDs := []interface{}{}
	for _, v := range tagDiff.Inc {
		for _, vv := range input.GetTags() {
			if v == vv.GetText() {
				addTagIDs = append(addTagIDs, vv.GetID())
			}
		}
	}
	removeTagIDs := []interface{}{}
	for _, v := range tagDiff.Dec {
		for _, vv := range dbLink.R.LinksTags {
			if v == vv.R.Tag.Text {
				removeTagIDs = append(removeTagIDs, vv.R.Tag.ID)
			}
		}
	}

	// tag add
	rels := []*rdb.LinksTag{}
	for _, v := range addTagIDs {
		i := &rdb.LinksTag{
			TagID: v.(int),
		}
		rels = append(rels, i)
	}
	if err := dbLink.AddLinksTags(ctx, r.db, true, rels...); err != nil {
		return err
	}

	// remove tag
	if _, err := dbLink.LinksTags(qm.WhereIn("tag_id in ?", removeTagIDs...)).DeleteAll(ctx, r.db); err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteLink(model *model.Link) error {
	ctx := context.Background()
	link, err := rdb.Links(rdb.LinkWhere.ID.EQ(model.GetID()), qm.Load(
		qm.Rels(
			rdb.LinkRels.LinksTags,
			rdb.LinksTagRels.Tag,
		),
	)).One(ctx, r.db)
	if err != nil {
		return err
	}
	if _, err := link.R.LinksTags.DeleteAll(ctx, r.db); err != nil {
		return err
	}
	if _, err := link.Delete(ctx, r.db); err != nil {
		return err
	}
	model = nil
	return nil
}
