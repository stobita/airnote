package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/stobita/airnote/internal/controller"
	"github.com/stobita/airnote/internal/infrastructure"
	"github.com/stobita/airnote/internal/presenter"
	"github.com/stobita/airnote/internal/repository"
	"github.com/stobita/airnote/internal/usecase"
)

// Run api server
func Run() error {
	db, err := infrastructure.NewDBConn()
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()

	httpClient := http.DefaultClient
	esClient, err := infrastructure.NewESClient()
	if err != nil {
		return err
	}
	repo := repository.New(db, httpClient, esClient)

	controller := controller.New(
		func(o usecase.OutputPort) usecase.InputPort {
			return usecase.NewInteractor(repo, o)
		},
		func(w http.ResponseWriter) usecase.OutputPort {
			return presenter.New(w)
		},
	)

	engine, err := getEngine(controller)
	if err != nil {
		return err
	}
	return engine.Run()
}

func getEngine(controller *controller.Controller) (*gin.Engine, error) {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type"},
		AllowOrigins: []string{"http://localhost:3000"},
	}))

	v1 := r.Group("/api/v1")
	{
		v1.GET("/links", controller.GetLinks())
		v1.POST("/links", controller.PostLink())
		v1.PUT("/links/:id", controller.UpdateLink())
		v1.DELETE("/links/:id", controller.DeleteLink())

		v1.GET("/tags", controller.GetTags())
		v1.GET("/tags/:id/links", controller.GetTagLinks())

		// NOTE: when implement websocket, close this route
		v1.GET("/links/:id/original", controller.GetLinkOriginal())
	}
	return r, nil
}
