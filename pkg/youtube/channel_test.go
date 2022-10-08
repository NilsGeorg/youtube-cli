package youtube

import (
	"testing"

	"github.com/NilsGeorg/ycli/internal/mock"
	"github.com/NilsGeorg/ycli/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestClient_Channels(t *testing.T) {
	client := TestClient(mock.ChannelResponse)

	response, err := client.Channels([]string{"UCuAXFkgsw1L7xaCfnd5JJOw"})
	assert.NoError(t, err)

	jsonResponse := test.StructToJSON(t, response)
	println(jsonResponse)

	assert.JSONEq(t, mock.ChannelResponse, jsonResponse)
}

func TestClient_ChannelByUsername(t *testing.T) {
	client := TestClient(mock.ChannelResponse)

	response, err := client.ChannelByUsername("RickAstleyOfficial")
	assert.NoError(t, err)

	jsonResponse := test.StructToJSON(t, response)

	assert.JSONEq(t, mock.ChannelResponse, jsonResponse)
}
