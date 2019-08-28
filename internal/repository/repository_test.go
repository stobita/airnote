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

func TestNew(t *testing.T) {
	t.Run("Success new repository", func(t *testing.T) {
		testDB, _ := testutils.GetTestDBConn()
		_ = repository.New(testDB)
	})
}

func TestRepository_SaveLink(t *testing.T) {
	testDB, truncate := testutils.GetTestDBConn()
	defer truncate()
	repo := repository.New(testDB)
	t.Run("Success save link", func(t *testing.T) {
		link, err := model.NewLink(model.LinkInput{URL: "test_url"})
		if err != nil {
			t.Fatal(err)
		}
		if err := repo.SaveLink(link); err != nil {
			t.Fatalf("save error: %s", err)
		}
		if link.GetID() == 0 {
			t.Errorf("link should be set id")
		}
	})
}

func TestRepository_SaveTag(t *testing.T) {
	testDB, truncate := testutils.GetTestDBConn()
	defer truncate()
	repo := repository.New(testDB)
	t.Run("Success save tag", func(t *testing.T) {
		tag, err := model.NewTag(model.TagInput{Text: "test_tag"})
		if err != nil {
			t.Fatal(err)
		}
		if err := repo.SaveTag(tag); err != nil {
			t.Fatalf("save error: %s", err)
		}
		if tag.GetID() == 0 {
			t.Errorf("tag should be set id")
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

func TestRepository_DeleteLink(t *testing.T) {
	testDB, truncate := testutils.GetTestDBConn()
	defer truncate()
	repo := repository.New(testDB)
	t.Run("Success delete link", func(t *testing.T) {
		link, err := model.NewLink(model.LinkInput{URL: "test_url"})
		if err != nil {
			t.Fatal(err)
		}
		if err := repo.SaveLink(link); err != nil {
			t.Fatal(err)
		}
		if err := repo.DeleteLink(link); err != nil {
			t.Fatalf("delete error: %s", err)
		}
		after, err := repo.GetLink(link.GetID())
		if err != nil {
			t.Fatal(err)
		}
		if after != nil {
			t.Fatalf("Want nil but get %v", after)
		}
	})
	t.Run("Success delete link with related tags", func(t *testing.T) {
		// TODO
	})
}

func TestRepository_GetTagByText(t *testing.T) {
	testDB, truncate := testutils.GetTestDBConn()
	defer truncate()
	repo := repository.New(testDB)
	t.Run("Success get link", func(t *testing.T) {
		testData := rdb.Tag{Text: "test"}
		if err := testData.Insert(context.Background(), testDB, boil.Infer()); err != nil {
			t.Fatal(err)
		}
		result, err := repo.GetTagByText("test")
		if err != nil {
			t.Fatalf("Failed get tag: %s", err)
		}
		if result.GetID() != testData.ID {
			t.Fatalf("Want %v but get %v", testData.ID, result.GetID())
		}
		if result.GetText() != testData.Text {
			t.Fatalf("Want %v but get %v", testData.Text, result.GetText())
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
		ctx := context.Background()
		for _, v := range testData {
			if err := v.Insert(ctx, testDB, boil.Infer()); err != nil {
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
