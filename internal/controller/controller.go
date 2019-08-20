package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stobita/airnote/internal/usecase"
)

type controller struct {
	inputPortFactory
	outputPortFactory
}

type postLinkRequestBody struct {
	URL         string `json:"url"`
	Description string `json:"description"`
}

type putLinkRequestBody postLinkRequestBody

type inputPortFactory func(o usecase.OutputPort) usecase.InputPort
type outputPortFactory func(w http.ResponseWriter) usecase.OutputPort

// New create controller
func New(i inputPortFactory, o outputPortFactory) *controller {
	return &controller{
		i,
		o,
	}
}

func (c *controller) GetLink() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		o := c.outputPortFactory(ctx.Writer)
		i := c.inputPortFactory(o)
		i.GetAllLinks()
	}
}

func (c *controller) PostLink() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		o := c.outputPortFactory(ctx.Writer)
		i := c.inputPortFactory(o)
		var json postLinkRequestBody
		if err := ctx.BindJSON(&json); err != nil {
			o.ResponseError(err)
			return
		}
		i.AddLink(usecase.InputData{
			URL:         json.URL,
			Description: json.Description,
		})
	}
}

func (c *controller) UpdateLink() gin.HandlerFunc {
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
		i.UpdateLink(id, usecase.InputData{
			URL:         json.URL,
			Description: json.Description,
		})
	}
}

func (c *controller) DeleteLink() gin.HandlerFunc {
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
