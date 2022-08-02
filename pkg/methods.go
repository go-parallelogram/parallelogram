package gotelebot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (b *Bot) GetUpdates(params GetUpdatesParams) ([]*Update, error) {
	return makeRequest[[]*Update](b, "getUpdates", params)
}

func (b *Bot) GetMe() (User, error) {
	return makeRequest[User](b, "getMe", nil)
}

func makeRequest[ResponseType any](b *Bot, method string, params any) (ResponseType, error) {
	const contentType = "application/json"

	var emptyResponse ResponseType

	requestBody, err := json.Marshal(params)
	if err != nil {
		return emptyResponse, err
	}

	resp, err := b.client.Post(
		fmt.Sprintf("%s/bot%s/%s", b.apiURL, b.token, method),
		contentType,
		bytes.NewReader(requestBody),
	)
	if err != nil {
		return emptyResponse, err
	}

	defer func() {
		if err = resp.Body.Close(); err != nil {
			log.Println("can't close response body while making request")
		}
	}()

	if resp.StatusCode == http.StatusOK {
		var successfulAPIResponse SuccessfulAPIResponse[ResponseType]
		if err := json.NewDecoder(resp.Body).Decode(&successfulAPIResponse); err != nil {
			return emptyResponse, err
		}

		return successfulAPIResponse.Result, nil
	}

	var failureAPIResponse FailureAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&failureAPIResponse); err != nil {
		return emptyResponse, err
	}

	// TODO add error handling on !apiResponse.Ok

	return emptyResponse, nil
}
