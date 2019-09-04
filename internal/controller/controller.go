package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stobita/airnote/internal/usecase"
)

type Controller struct {
	inputPortFactory
	outputPortFactory
}

type postLinkRequestBody struct {
	URL         string   `json:"url"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

type putLinkRequestBody postLinkRequestBody

type postLinkTagRequestBody struct {
	LinkID int    `json:"linkId"`
	Text   string `json:"text"`
}

type inputPortFactory func(o usecase.OutputPort) usecase.InputPort
type outputPortFactory func(w http.ResponseWriter) usecase.OutputPort

// New create controller
func New(i inputPortFactory, o outputPortFactory) *Controller {
	return &Controller{
		i,
		o,
	}
}

func (c *Controller) GetLinks() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		o := c.outputPortFactory(ctx.Writer)
		i := c.inputPortFactory(o)
		word := ctx.Query("word")
		if word == "" {
			i.GetAllLinks()
		} else {
			i.SearchLinks(word)
		}
	}
}

func (c *Controller) PostLink() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		o := c.outputPortFactory(ctx.Writer)
		i := c.inputPortFactory(o)
		var json postLinkRequestBody
		if err := ctx.BindJSON(&json); err != nil {
			o.ResponseError(err)
			return
		}
		i.AddLink(usecase.LinkInputData{
			URL:         json.URL,
			Description: json.Description,
			Tags:        json.Tags,
		})
	}
}

func (c *Controller) UpdateLink() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		o := c.outputPortFactory(ctx.Writer)
		i := c.inputPortFactory(o)
		var json putLinkRequestBody
		if err := ctx.BindJSON(&json); err != nil {
			o.ResponseError(err)
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			o.ResponseError(err)
			return
		}
		i.UpdateLink(id, usecase.LinkInputData{
			URL:         json.URL,
			Description: json.Description,
			Tags:        json.Tags,
		})
	}
}

func (c *Controller) DeleteLink() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		o := c.outputPortFactory(ctx.Writer)
		i := c.inputPortFactory(o)
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			o.ResponseError(err)
			return
		}
		i.DeleteLink(id)
	}
}

func (c *Controller) GetLinkOriginal() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		o := c.outputPortFactory(ctx.Writer)
		i := c.inputPortFactory(o)
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			o.ResponseError(err)
			return
		}
		i.GetLinkOriginal(id)
	}
}

func (c *Controller) GetTags() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		o := c.outputPortFactory(ctx.Writer)
		i := c.inputPortFactory(o)
		i.GetAllTags()
	}
}

func (c *Controller) GetTagLinks() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		o := c.outputPortFactory(ctx.Writer)
		i := c.inputPortFactory(o)
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			o.ResponseError(err)
			return
		}
		i.GetTaggedLinks(id)

	}
}
