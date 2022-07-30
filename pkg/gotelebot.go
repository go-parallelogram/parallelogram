package gotelebot

import (
	"net/http"
)

const TelegramAPIServer = "https://api.telegram.org"

type Bot struct {
	apiURL string
	token  string

	client *http.Client
}

func NewBot(apiURL, token string) (*Bot, error) {
	return &Bot{
		token:  token,
		apiURL: apiURL,

		client: &http.Client{},
	}, nil
}
