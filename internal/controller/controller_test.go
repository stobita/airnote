package controller_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stobita/airnote/internal/controller"
	"github.com/stobita/airnote/internal/usecase"
)

type fakeInputPort struct {
	getAllLinksSuccess bool
	getAllLinksError   bool
	getAddLinkSuccess  bool
	getAddLinkError    bool
}

func (f *fakeInputPort) GetAllLinks() error {
	if f.getAllLinksError {
		return errors.New("Fake error")
	}
	f.getAllLinksSuccess = true
	return nil
}
func (f *fakeInputPort) AddLink(i usecase.InputData) error {
	if f.getAddLinkError {
		return errors.New("Fake error")
	}
	f.getAddLinkSuccess = true
	return nil
}

func TestController_GetLink(t *testing.T) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Error create request")
	}
	t.Run("Success", func(t *testing.T) {
		inputPort := &fakeInputPort{}
		c := controller.New(func(w http.ResponseWriter) usecase.InputPort {
			return inputPort
		})
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
		c := controller.New(func(w http.ResponseWriter) usecase.InputPort {
			return inputPort
		})
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
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Error create request")
	}
	t.Run("Success", func(t *testing.T) {
		inputPort := &fakeInputPort{}
		c := controller.New(func(w http.ResponseWriter) usecase.InputPort {
			return inputPort
		})
		r := gin.Default()
		r.GET("/", c.PostLink())
		r.ServeHTTP(w, req)

		if !inputPort.getAddLinkSuccess {
			t.Error("Failed call input port")
		}
	})
	t.Run("Error at input port", func(t *testing.T) {
		inputPort := &fakeInputPort{
			getAddLinkError: true,
		}
		c := controller.New(func(w http.ResponseWriter) usecase.InputPort {
			return inputPort
		})
		r := gin.Default()
		r.GET("/", c.PostLink())
		r.ServeHTTP(w, req)

		if inputPort.getAddLinkSuccess {
			t.Error("Want error but get success")
		}
	})

}
