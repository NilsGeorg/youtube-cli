package youtube

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRequest(t *testing.T) {
	var request = NewRequest()

	assert.NotNil(t, request.Params)
	assert.Len(t, request.Params, 0)
}

func TestRequest_AddParam(t *testing.T) {
	for _, paramList := range getTestParams() {
		var name = fmt.Sprintf("Size-%d", len(paramList))

		t.Run(name, func(t *testing.T) {
			var request = NewRequest()

			for key, val := range paramList {
				request.AddParam(key, val)
			}

			assert.Len(t, request.Params, len(paramList))
			assert.Equal(t, paramList, request.Params)
		})
	}
}

func TestRequest_GetAsQuery(t *testing.T) {
	for _, paramList := range getTestParams() {
		var name = fmt.Sprintf("Size-%d", len(paramList))

		t.Run(name, func(t *testing.T) {
			var request = NewRequest()

			query := url.Values{}
			for key, val := range paramList {
				request.AddParam(key, val)
				query.Add(key, val)
			}

			assert.Equal(t, query.Encode(), request.GetAsQuery())
		})
	}
}

func getTestParams() []map[string]string {
	var paramTest = []map[string]string{}
	var amountTestRuns = 10

	for i := 1; i <= amountTestRuns; i++ {
		var params = map[string]string{}
		for j := 1; j <= i; j++ {
			params[fmt.Sprintf("%d-key-%d", i, j)] = fmt.Sprintf("%d-value-%d", i, j)
		}

		paramTest = append(paramTest, params)
	}

	return paramTest
}
