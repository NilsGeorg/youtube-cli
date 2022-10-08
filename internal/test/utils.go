package test

import (
	"encoding/json"
	"testing"
)

func StructToJSON(t *testing.T, response interface{}) string {
	t.Helper()

	jsonByte, err := json.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}

	return string(jsonByte)
}
