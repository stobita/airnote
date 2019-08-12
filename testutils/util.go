package testutils

import (
	"encoding/json"

	"github.com/google/go-cmp/cmp"
)

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
