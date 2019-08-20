package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stobita/airnote/internal/controller"
	"github.com/stobita/airnote/internal/usecase"
)

func TestController_GetLink(t *testing.T) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Error create request")
	}
	t.Run("Success", func(t *testing.T) {
		inputPort := &fakeInputPort{}
		outputPort := &fakeOutputPort{}
		c := controller.New(
			func(o usecase.OutputPort) usecase.InputPort {
				return inputPort
			},
			func(w http.ResponseWriter) usecase.OutputPort {
				return outputPort
			},
		)
		r := gin.Default()
		r.GET("/", c.GetLink())
		r.ServeHTTP(w, req)

		if !inputPort.getAllLinksSuccess {
			t.Error("Failed call input port")
		}
	})
	t.Run("Error at input port", func(t *testing.T) {
		inputPort := &fakeInputPort{
			getAllLinksError: true,
		}
		outputPort := &fakeOutputPort{}
		c := controller.New(
			func(o usecase.OutputPort) usecase.InputPort {
				return inputPort
			},
			func(w http.ResponseWriter) usecase.OutputPort {
				return outputPort
			},
		)
		r := gin.Default()
		r.GET("/", c.GetLink())
		r.ServeHTTP(w, req)

		if inputPort.getAllLinksSuccess {
			t.Error("Want error but get success")
		}
	})
}

func TestController_PostLink(t *testing.T) {
	w := httptest.NewRecorder()
	body, err := json.Marshal(controller.ExportPostLinkRequestBody{
		URL: "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Error create request")
	}
	t.Run("Success", func(t *testing.T) {
		inputPort := &fakeInputPort{}
		outputPort := &fakeOutputPort{}
		c := controller.New(
			func(o usecase.OutputPort) usecase.InputPort {
				return inputPort
			},
			func(w http.ResponseWriter) usecase.OutputPort {
				return outputPort
			},
		)
		r := gin.Default()
		r.POST("/", c.PostLink())
		r.ServeHTTP(w, req)
		if !inputPort.getAddLinkSuccess {
			t.Error("Failed call input port")
		}
	})
	t.Run("Error at input port", func(t *testing.T) {
		inputPort := &fakeInputPort{
			getAddLinkError: true,
		}
		outputPort := &fakeOutputPort{}
		c := controller.New(
			func(o usecase.OutputPort) usecase.InputPort {
				return inputPort
			},
			func(w http.ResponseWriter) usecase.OutputPort {
				return outputPort
			},
		)
		r := gin.Default()
		r.POST("/", c.PostLink())
		r.ServeHTTP(w, req)

		if inputPort.getAddLinkSuccess {
			t.Error("Want error but get success")
		}
	})

}
