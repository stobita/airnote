package controller_test

import (
	"bytes"
	"encoding/json"
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

func (f *fakeInputPort) UpdateLink(id int, i usecase.InputData) error {
	return nil
}

func (f *fakeInputPort) DeleteLink(id int) error {
	return nil
}

type fakeOutputPort struct{}

func (f *fakeOutputPort) ResponseLink(o usecase.LinkOutputData) error {
	return nil
}

func (f *fakeOutputPort) ResponseLinks(o usecase.LinksOutputData) error {
	return nil
}

func (f *fakeOutputPort) ResponseError(e error) error {
	return nil
}

func (f *fakeOutputPort) ResponseNoContent() error {
	return nil
}

func TestController_GetLink(t *testing.T) {
	w := httptest.NewRecorder()
	body, err := json.Marshal(controller.ExportPostLinkRequestBody{
		URL: "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("GET", "/", bytes.NewBuffer(body))
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
	req, err := http.NewRequest("POST", "/", nil)
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
