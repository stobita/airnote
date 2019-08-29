package util

import (
	"encoding/json"

	"github.com/google/go-cmp/cmp"
)

type StringDiff struct {
	Inc []string
	Dec []string
}

func StringArrayDiff(before []string, after []string) StringDiff {
	return StringDiff{
		Inc: StringArraySub(after, before),
		Dec: StringArraySub(before, after),
	}
}

// StringArraySub return a - b
func StringArraySub(a []string, b []string) []string {
	r := []string{}
	m := make(map[string]bool)
	for _, v := range b {
		m[v] = true
	}
	for _, v := range a {
		if _, ok := m[v]; !ok {
			r = append(r, v)
		}
	}
	return r
}

func JSONStringEqual(i1, i2 string) (bool, error) {
	var json1, json2 interface{}
	if err := json.Unmarshal([]byte(i1), &json1); err != nil {
		return false, err
	}
	if err := json.Unmarshal([]byte(i2), &json2); err != nil {
		return false, err
	}
	return cmp.Equal(json1, json2), nil
}
