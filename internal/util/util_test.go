package util_test

import (
	"testing"

	"github.com/stobita/airnote/internal/util"
)

func TestStringArrayDiff(t *testing.T) {
	t.Run("Success equal array", func(t *testing.T) {
		before := []string{"test1", "test2", "test3"}
		after := []string{"test1", "test2", "test3"}
		result := util.StringArrayDiff(before, after)
		if len(result.Inc) != 0 {
			t.Errorf("Inc want 0 but get %v", result.Inc)
		}
		if len(result.Dec) != 0 {
			t.Errorf("Dec want 0 but get %v", result.Dec)
		}
	})
	t.Run("Success increment array", func(t *testing.T) {
		before := []string{"test1", "test2", "test3"}
		after := []string{"test1", "test2", "test3", "test4"}
		result := util.StringArrayDiff(before, after)
		if len(result.Inc) != 1 {
			t.Errorf("Inc want 1 but get %v", result.Inc)
		}
		if result.Inc[0] != "test4" {
			t.Errorf("Inc[0] want test4 but get %v", result.Inc[0])
		}
		if len(result.Dec) != 0 {
			t.Errorf("Dec want 0 but get %v", result.Dec)
		}
	})
	t.Run("Success decrement array", func(t *testing.T) {
		before := []string{"test1", "test2", "test3"}
		after := []string{"test1", "test3"}
		result := util.StringArrayDiff(before, after)
		if len(result.Dec) != 1 {
			t.Errorf("Dec want 1 but get %v", result.Dec)
		}
		if result.Dec[0] != "test2" {
			t.Errorf("Dec[0] want test2 but get %v", result.Dec[0])
		}
		if len(result.Inc) != 0 {
			t.Errorf("Dec want 0 but get %v", result.Dec)
		}
	})
	t.Run("Success increment and decrement array", func(t *testing.T) {
		before := []string{"test1", "test2", "test3"}
		after := []string{"test1", "test3", "test4"}
		result := util.StringArrayDiff(before, after)
		if len(result.Inc) != 1 {
			t.Errorf("Inc want 1 but get %v", result.Inc)
		}
		if result.Inc[0] != "test4" {
			t.Errorf("Inc[0] want test4 but get %v", result.Inc[0])
		}
		if len(result.Dec) != 1 {
			t.Errorf("Dec want 1 but get %v", result.Dec)
		}
		if result.Dec[0] != "test2" {
			t.Errorf("Dec[0] want test2 but get %v", result.Dec[0])
		}
	})
}
