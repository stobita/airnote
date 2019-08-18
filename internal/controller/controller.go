package controller

import (
	"net/http"

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

type inputPortFactory func(o usecase.OutputPort) usecase.InputPort
type outputPortFactory func(w http.ResponseWriter) usecase.OutputPort

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
		if err := i.GetAllLinks(); err != nil {
			o.ResponseError(err)
		}
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
		if err := i.AddLink(usecase.InputData{
			URL:         json.URL,
			Description: json.Description,
		}); err != nil {
			o.ResponseError(err)
		}
	}
}
