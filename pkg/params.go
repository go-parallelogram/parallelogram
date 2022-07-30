package gotelebot

const (
	MessageUpdate            UpdateType = "message"
	EditedMessageUpdate      UpdateType = "edited_message"
	ChannelPostUpdate        UpdateType = "channel_post"
	EditedChannelPostUpdate  UpdateType = "edited_channel_post"
	InlineQueryUpdate        UpdateType = "inline_query"
	ChosenInlineResultUpdate UpdateType = "chosen_inline_result"
	CallbackQueryUpdate      UpdateType = "callback_query"
	ShippingQueryUpdate      UpdateType = "shipping_query"
	PreCheckoutQueryUpdate   UpdateType = "pre_checkout_query"
	PollUpdate               UpdateType = "poll"
	PollAnswerUpdate         UpdateType = "poll_answer"
	MyChatMemberUpdate       UpdateType = "my_chat_member"
	ChatMemberUpdate         UpdateType = "chat_member"
	ChatJoinRequestUpdate    UpdateType = "chat_join_request"
)

type UpdateType string

type GetUpdatesParams struct {
	Offset         int64        `json:"offset,omitempty"`
	Limit          int          `json:"limit,omitempty"`
	Timeout        int          `json:"timeout,omitempty"`
	AllowedUpdates []UpdateType `json:"allowed_updates,omitempty"`
}
