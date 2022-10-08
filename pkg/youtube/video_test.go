package youtube

import (
	"testing"

	"github.com/NilsGeorg/ycli/internal/mock"
	"github.com/NilsGeorg/ycli/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestClient_VideoSearch(t *testing.T) {
	client := TestClient(mock.VideoSearchResponse)

	searchResponse, err := client.VideoSearch("Never gonna give you up", "", "")
	assert.NoError(t, err)

	jsonResponse := test.StructToJSON(t, searchResponse)
	assert.JSONEq(t, mock.VideoSearchResponse, jsonResponse)
}

func TestClient_Videos(t *testing.T) {
	client := TestClient(mock.VideoResponse)

	searchResponse, err := client.Videos([]string{"1", "2"})
	assert.NoError(t, err)

	jsonResponse := test.StructToJSON(t, searchResponse)

	assert.JSONEq(t, mock.VideoResponse, jsonResponse)
}

func TestClient_VideosByChart(t *testing.T) {
	client := TestClient(mock.VideoResponse)

	searchResponse, err := client.VideosByChart("", "")
	assert.NoError(t, err)

	jsonResponse := test.StructToJSON(t, searchResponse)

	assert.JSONEq(t, mock.VideoResponse, jsonResponse)
}

func TestClient_VideoGetCategoriesByRegion(t *testing.T) {
	client := TestClient(mock.VideoCategoryResponseByRegion)

	searchResponse, err := client.VideoGetCategories("", "1")
	assert.NoError(t, err)

	jsonResponse := test.StructToJSON(t, searchResponse)

	assert.JSONEq(t, mock.VideoCategoryResponseByRegion, jsonResponse)
}

func TestClient_VideoGetCategoriesByCategory(t *testing.T) {
	client := TestClient(mock.VideoCategoryResponseByCategory)

	searchResponse, err := client.VideoGetCategories("DE", "")
	assert.NoError(t, err)

	jsonResponse := test.StructToJSON(t, searchResponse)

	assert.JSONEq(t, mock.VideoCategoryResponseByCategory, jsonResponse)
}

func TestClient_VideoGetCategoriesBoth(t *testing.T) {
	client := TestClient("")

	searchResponse, err := client.VideoGetCategories("DE", "1")

	assert.Error(t, err)
	assert.Empty(t, searchResponse)
}
