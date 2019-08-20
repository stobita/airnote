package repository_test

import (
	"context"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stobita/airnote/internal/domain/model"
	"github.com/stobita/airnote/internal/repository"
	"github.com/stobita/airnote/internal/repository/rdb"
	"github.com/stobita/airnote/testutils"
	"github.com/volatiletech/sqlboiler/boil"
)

func TestMain(m *testing.M) {
	result, err := testutils.StorageTestRunner(m)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	os.Exit(result)
}

func TestRepository_CreateLink(t *testing.T) {
	testDB, truncate := testutils.GetTestDBConn()
	defer truncate()
	repo := repository.New(testDB)
	t.Run("Success save link", func(t *testing.T) {
		link, err := model.NewLink(model.LinkInput{URL: "test_url"})
		if err != nil {
			t.Fatal(err)
		}
		if err := repo.CreateLink(link); err != nil {
			t.Fatalf("create error: %s", err)
		}
		if link.GetID() == 0 {
			t.Errorf("link should be set id")
		}
	})
}

func TestRepository_GetLink(t *testing.T) {
	testDB, truncate := testutils.GetTestDBConn()
	defer truncate()
	repo := repository.New(testDB)
	t.Run("Success get link", func(t *testing.T) {
		testData := rdb.Link{URL: "test"}
		if err := testData.Insert(context.Background(), testDB, boil.Infer()); err != nil {
			t.Fatal(err)
		}
		result, err := repo.GetLink(testData.ID)
		if err != nil {
			t.Fatalf("Failed get link: %s", err)
		}
		if result.GetID() != testData.ID {
			t.Fatalf("Want %v but get %v", testData.ID, result.GetID())
		}
		if result.GetURL() != testData.URL {
			t.Fatalf("Want %v but get %v", testData.URL, result.GetURL())
		}
	})
}

func TestRepository_GetLinks(t *testing.T) {
	testDB, truncate := testutils.GetTestDBConn()
	defer truncate()
	repo := repository.New(testDB)
	t.Run("Success get all links", func(t *testing.T) {
		// set test data
		testData := []rdb.Link{
			rdb.Link{URL: "test_1"},
			rdb.Link{URL: "test_2"},
			rdb.Link{URL: "test_3"},
		}
		for _, v := range testData {
			if err := v.Insert(context.Background(), testDB, boil.Infer()); err != nil {
				t.Fatalf("error set test data: %s", err)
			}
		}

		result, err := repo.GetLinks()
		if err != nil {
			t.Fatalf("Failed get links: %s", err)
		}

		if len(result) != 3 {
			t.Errorf("want 3 but get: %v", len(result))
		}

		for _, v := range result {
			if v.GetID() == 0 {
				t.Errorf("link should be set id")
			}
		}
	})

	t.Run("No record", func(t *testing.T) {
		truncate()
		result, err := repo.GetLinks()
		if err != nil {
			t.Fatalf("Failed get links: %s", err)
		}
		if len(result) != 0 {
			t.Errorf("want 0 but get: %v", len(result))
		}
	})

}
