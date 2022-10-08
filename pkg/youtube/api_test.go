package youtube

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/NilsGeorg/ycli/internal/mock"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func Test_Get(t *testing.T) {
	var path = "/path"
	var key = "paramKey"
	var value = "paramValue"

	var request = NewRequest().
		AddParam(key, value)

	jsonString, videoCategoryResponse := getJSONAndStruct()

	defer gock.Off()
	gock.
		New("https://" + TestHTTPClient.YoutubeAPIURL).
		Get(TestHTTPClient.YoutubeAPIPath + path).
		Reply(200).
		JSON(jsonString)

	response, err := NewAPI[VideoCategoryResource](TestHTTPClient).Get(path, request)

	assert.True(t, gock.IsDone())
	assert.NoError(t, err)
	assert.Equal(t, videoCategoryResponse, response)
}

func Test_CreateRequest(t *testing.T) {
	var path = "/path"
	var key = "paramKey"
	var value = "paramValue"

	request := NewRequest().
		AddParam("paramKey", "paramValue").
		AddParam("key", "value")

	var apiImpl = NewAPI[VideoCategoryResource](TestHTTPClient)

	httpRequest, err := apiImpl.createRequest(path, request)

	assert.NoError(t, err)
	assert.Equal(t, "GET", httpRequest.Method)
	assert.Equal(t, TestHTTPClient.YoutubeAPIPath+path, httpRequest.URL.Path)
	assert.Equal(t, TestHTTPClient.YoutubeAPIURL, httpRequest.URL.Host)
	assert.Equal(t, fmt.Sprintf("key=%s&%s=%s", mock.AuthenticationToken, key, value), httpRequest.URL.RawQuery)
}

func Test_ResponseToStruct(t *testing.T) {
	jsonString, videoCategoryResponse := getJSONAndStruct()
	httpResponse := &http.Response{Body: io.NopCloser(bytes.NewBufferString(jsonString))}

	var apiImpl = NewAPI[VideoCategoryResource](TestHTTPClient)

	response, err := apiImpl.responseToStruct(httpResponse)

	assert.NoError(t, err)
	assert.Equal(t, videoCategoryResponse, response)
}

func getJSONAndStruct() (string, *Response[VideoCategoryResource]) {
	var kind = "youtube#videoCategoryListResponse"
	var etag = "s6RguhiCzdsBQFsS4YzslvqtQtI"
	var itemKind = "youtube#videoCategory"
	var itemEtag = "s6RguhiCzdsBQFsS4YzslvqtQtI"
	var itemID = "1"
	var snippetTitle = "Film & Animation"
	var snippetAssignable = true
	var snippetChannelID = "UCBR8-60-B28hp2BmDPdntcQ"

	var videoCategoryResponse = &Response[VideoCategoryResource]{
		Kind: kind,
		Etag: etag,
		Items: []*VideoCategoryResource{
			{
				Kind: itemKind,
				Etag: itemEtag,
				ID:   itemID,
				Snippet: struct {
					ChannelID  string `json:"channelId"`
					Title      string `json:"title"`
					Assignable bool   `json:"assignable"`
				}{
					Title:      snippetTitle,
					Assignable: snippetAssignable,
					ChannelID:  snippetChannelID,
				},
			},
		},
	}

	jsonString := fmt.Sprintf(`{
    "kind": "%s",
    "etag": "%s",
    "items": [
        {
            "kind": "%s",
            "etag": "%s",
            "id": "%s",
            "snippet": {
                "title": "%s",
                "assignable": %t,
                "channelId": "%s"
            }
        }
    ]
}`, kind, etag, itemKind, itemEtag, itemID, snippetTitle, snippetAssignable, snippetChannelID)

	return jsonString, videoCategoryResponse
}
