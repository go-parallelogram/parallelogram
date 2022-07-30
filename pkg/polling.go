package gotelebot

type UpdateHandler interface {
	HandleUpdate(upd Update)
}

type HandlerFunc func(upd Update)

func (h HandlerFunc) HandleUpdate(upd Update) {
	h(upd)
}

func RunPolling(bot *Bot, handler UpdateHandler) error {
	var offset int64 = 0

	for {
		updates, err := bot.GetUpdates(GetUpdatesParams{
			Offset: offset,
		})
		if err != nil {
			return err
		}

		for _, update := range updates {
			offset = update.UpdateID + 1
			handler.HandleUpdate(update)
		}
	}
}
