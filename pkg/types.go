package gotelebot

import (
	"encoding/json"
)

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

type SuccessfulAPIResponse[ResponseType any] struct {
	Ok     bool         `json:"ok"`
	Result ResponseType `json:"result"`
}

type FailureAPIResponse struct {
	Ok bool `json:"ok"`

	Description string `json:"description,omitempty"`
	ErrorCode   int    `json:"error_code,omitempty"`
	Parameters  struct {
		MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`
		RetryAfter      int   `json:"retry_after,omitempty"`
	} `json:"parameters,omitempty"`
}

type Update struct {
	UpdateID           int64           `json:"update_id"`
	Message            json.RawMessage `json:"message,omitempty"`
	EditedMessage      json.RawMessage `json:"edited_message,omitempty"`
	ChannelPost        json.RawMessage `json:"channel_post,omitempty"`
	EditedChannelPost  json.RawMessage `json:"edited_channel_post,omitempty"`
	InlineQuery        json.RawMessage `json:"inline_query,omitempty"`
	ChosenInlineResult json.RawMessage `json:"chosen_inline_result,omitempty"`
	CallbackQuery      json.RawMessage `json:"callback_query,omitempty"`
	ShippingQuery      json.RawMessage `json:"shipping_query,omitempty"`
	PreCheckoutQuery   json.RawMessage `json:"pre_checkout_query,omitempty"`
	Poll               json.RawMessage `json:"poll,omitempty"`
	PollAnswer         json.RawMessage `json:"poll_answer,omitempty"`
	MyChatMember       json.RawMessage `json:"my_chat_member,omitempty"`
	ChatMember         json.RawMessage `json:"chat_member,omitempty"`
	ChatJoinRequest    json.RawMessage `json:"chat_join_request,omitempty"`
}

func (u *Update) Type() UpdateType {
	if u.Message != nil {
		return MessageUpdate
	}
	if u.EditedMessage != nil {
		return EditedMessageUpdate
	}
	if u.ChannelPost != nil {
		return ChannelPostUpdate
	}
	if u.EditedChannelPost != nil {
		return EditedChannelPostUpdate
	}
	if u.InlineQuery != nil {
		return InlineQueryUpdate
	}
	if u.ChosenInlineResult != nil {
		return ChosenInlineResultUpdate
	}
	if u.CallbackQuery != nil {
		return CallbackQueryUpdate
	}
	if u.ShippingQuery != nil {
		return ShippingQueryUpdate
	}
	if u.PreCheckoutQuery != nil {
		return PreCheckoutQueryUpdate
	}
	if u.Poll != nil {
		return PollUpdate
	}
	if u.PollAnswer != nil {
		return PollAnswerUpdate
	}
	if u.MyChatMember != nil {
		return MyChatMemberUpdate
	}
	if u.ChatMember != nil {
		return ChatMemberUpdate
	}
	if u.ChatJoinRequest != nil {
		return ChatJoinRequestUpdate
	}
	return ""
}

type User struct {
	ID int64 `json:"id"`
}
