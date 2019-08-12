package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stobita/airnote/internal/usecase"
)

type controller struct {
	inputPortFactory
}

type inputPortFactory func(w http.ResponseWriter) usecase.InputPort

func New(i inputPortFactory) *controller {
	return &controller{
		i,
	}
}

func (c *controller) GetLink() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		i := c.inputPortFactory(ctx.Writer)
		if err := i.GetAllLinks(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}

func (c *controller) PostLink() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		i := c.inputPortFactory(ctx.Writer)
		if err := i.AddLink(usecase.InputData{URL: "test"}); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}
