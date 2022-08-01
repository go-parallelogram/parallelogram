package gotelebot

type GetUpdatesParams struct {
	Offset         int64        `json:"offset,omitempty"`
	Limit          int          `json:"limit,omitempty"`
	Timeout        int          `json:"timeout,omitempty"`
	AllowedUpdates []UpdateType `json:"allowed_updates,omitempty"`
}
