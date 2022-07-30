package gotelebot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func (b *Bot) GetUpdates(params GetUpdatesParams) ([]*Update, error) {
	return MakeRequest[[]*Update](b, "getUpdates", params)
}

func (b *Bot) GetMe() (User, error) {
	return MakeRequest[User](b, "getMe", nil)
}

func MakeRequest[ResponseType any](b *Bot, method string, params any) (ResponseType, error) {
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

	var apiResponse APIResponse[ResponseType]
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return emptyResponse, err
	}

	// TODO add error handling on !apiResponse.Ok

	return apiResponse.Result, nil
}
