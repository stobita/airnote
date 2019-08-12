package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stobita/airnote/internal/controller"
	"github.com/stobita/airnote/internal/presenter"
	"github.com/stobita/airnote/internal/repository"
	"github.com/stobita/airnote/internal/usecase"
)

// Run api server
func Run() error {
	engine, err := getEngine()
	if err != nil {
		return err
	}
	return engine.Run()
}

func getEngine() (*gin.Engine, error) {
	r := gin.Default()

	db, err := repository.NewDBConn()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	repo := repository.New(db)

	controller := controller.New(func(w http.ResponseWriter) usecase.InputPort {
		return usecase.NewInteractor(repo, presenter.New(w))
	})

	v1 := r.Group("/api/v1")
	{
		v1.GET("/links", controller.GetLink())
		v1.POST("/links", controller.PostLink())

	}
	return r, nil
}
