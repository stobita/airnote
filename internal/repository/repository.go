package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/pkg/errors"
	"github.com/stobita/airnote/internal/domain/model"
	"github.com/stobita/airnote/internal/repository/rdb"
	"github.com/stobita/airnote/internal/util"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
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
	ctx := context.Background()
	link, err := rdb.FindLink(ctx, r.db, id)
	if err != nil {
		return nil, err
	}
	m, err := model.NewLink(model.LinkInput{
		URL:         link.URL,
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
			rdb.LinkRels.LinksTags,
			qm.Load(rdb.LinksTagRels.Tag),
		),
	).All(ctx, r.db)
	if err != nil {
		return nil, err
	}
	var result []*model.Link
	for _, v := range links {
		var tags []*model.Tag
		for _, v := range v.R.LinksTags {
			tag, err := model.NewTag(model.TagInput{
				Text: v.R.Tag.Text,
			})
			if err != nil {
				return nil, err
			}
			tags = append(tags, tag)
		}
		input := model.LinkInput{
			URL:         v.URL,
			Description: v.Description.String,
			Tags:        tags,
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

func (r *repository) SaveLink(input *model.Link) error {
	ctx := context.Background()
	if len(input.GetTags()) > 0 {
		if err := r.findOrSaveTags(input.GetTags()); err != nil {
			return err
		}
	}
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

func (r *repository) findOrSaveTags(input []*model.Tag) error {
	ctx := context.Background()
	textList := []interface{}{}
	for _, v := range input {
		textList = append(textList, v.GetText())
	}

	saved, err := rdb.Tags(qm.WhereIn("text in ?", textList...)).All(ctx, r.db)
	if err != nil {
		return errors.Wrap(err, "Failed get saved tags")
	}

	savedTags := []*model.Tag{}

	for _, v := range saved {
		tag, err := model.NewTag(model.TagInput{
			Text: v.Text,
		})
		if err != nil {
			return err
		}
		tag.SetID(v.ID)
		savedTags = append(savedTags, tag)
	}

	// t - saved
	var inputText []string
	for _, v := range input {
		inputText = append(inputText, v.GetText())
	}
	var savedText []string
	for _, v := range saved {
		savedText = append(savedText, v.Text)
	}
	unsavedText := util.StringArraySub(inputText, savedText)
	unsaved := []*model.Tag{}
	for _, v := range unsavedText {
		model, err := model.NewTag(model.TagInput{
			Text: v,
		})
		if err != nil {
			return errors.Wrap(err, "NewTag error")
		}
		unsaved = append(unsaved, model)
	}
	if err != nil {
		return errors.Wrap(err, "getUnSavedTags error")
	}
	if len(unsaved) == 0 {
		return nil
	}
	// NOTE: if slow, use bulk insert by sql
	for _, v := range unsaved {
		tag := rdb.Tag{
			Text: v.GetText(),
		}
		if err := tag.Insert(ctx, r.db, boil.Whitelist("text")); err != nil {
			return err
		}
		v.SetID(tag.ID)
		input = append(input, v)
	}
	for _, v := range input {
		log.Printf("tag: %#v", v)
	}
	return nil
}

func (r *repository) UpdateLink(input *model.Link) error {
	ctx := context.Background()
	if len(input.GetTags()) > 0 {
		if err := r.findOrSaveTags(input.GetTags()); err != nil {
			return err
		}
	}

	link, err := rdb.FindLink(ctx, r.db, input.GetID())
	if err != nil {
		return err
	}
	link.URL = input.GetURL()
	link.Description = null.StringFrom(input.GetDescription())
	if _, err := link.Update(ctx, r.db, boil.Whitelist(
		"url",
		"description",
	)); err != nil {
		return err
	}

	var beforeText []string
	for _, v := range link.R.LinksTags {
		beforeText = append(beforeText, v.R.Tag.Text)
	}

	var afterText []string
	for _, v := range input.GetTags() {
		afterText = append(afterText, v.GetText())
	}

	tagDiff := util.StringArrayDiff(beforeText, afterText)

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
		for _, vv := range link.R.LinksTags {
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
	if err := link.AddLinksTags(ctx, r.db, true, rels...); err != nil {
		return err
	}

	// remove tag
	if _, err := link.LinksTags(qm.WhereIn("tag_id in ?", removeTagIDs...)).DeleteAll(ctx, r.db); err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteLink(model *model.Link) error {
	ctx := context.Background()
	link, err := rdb.FindLink(ctx, r.db, model.GetID())
	if err != nil {
		return err
	}
	if _, err := link.Delete(ctx, r.db); err != nil {
		return err
	}
	model = nil
	return nil
}
