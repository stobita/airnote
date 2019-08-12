package presenter_test

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/stobita/airnote/internal/presenter"
	"github.com/stobita/airnote/internal/usecase"
	"github.com/stobita/airnote/testutils"
)

func TestPresenter_ResponseLinks(t *testing.T) {
	t.Run("Success response link list", func(t *testing.T) {
		w := httptest.NewRecorder()
		p := presenter.New(w)
		o := usecase.LinksOutputData{
			&usecase.LinkOutputData{
				URL: "test",
			},
		}
		if err := p.ResponseLinks(o); err != nil {
			t.Fatalf("Failed response link list: %s", err)
		}
		result := w.Result()
		body, _ := ioutil.ReadAll(result.Body)
		expect := `
			{
				"items": [
					{
						"url": "test"
					}
				]
			}
		`

		if match, err := testutils.JSONStringEqual(string(body), expect); err != nil {
			t.Errorf("Invalid result: %s", body)
		} else if !match {
			t.Errorf("want %s but get %s", expect, body)
		}
	})
}

func TestPresenter_ResponseLink(t *testing.T) {
	t.Run("Success response link", func(t *testing.T) {
		w := httptest.NewRecorder()
		p := presenter.New(w)
		o := usecase.LinkOutputData{
			URL: "test",
		}
		if err := p.ResponseLink(o); err != nil {
			t.Fatalf("Failed response link list: %s", err)
		}
		result := w.Result()
		body, _ := ioutil.ReadAll(result.Body)
		expect := `
			{
				"url": "test"
			}
		`

		if match, err := testutils.JSONStringEqual(string(body), expect); err != nil {
			t.Errorf("Invalid result: %s", body)
		} else if !match {
			t.Errorf("want %s but get %s", expect, body)
		}
	})
}
