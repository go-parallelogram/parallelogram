package gotelebot

import (
	"context"
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

type APIResponse[ResultType any] struct {
	Ok          bool       `json:"ok"`
	Result      ResultType `json:"result,omitempty"`
	Description string     `json:"description,omitempty"`
	ErrorCode   int        `json:"error_code,omitempty"`
	Parameters  struct {
		MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`
		RetryAfter      int   `json:"retry_after,omitempty"`
	} `json:"parameters,omitempty"`
}

type Update struct {
	ctx context.Context

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

func (u *Update) Context() context.Context {
	if u.ctx != nil {
		return u.ctx
	}
	return context.Background()
}

func (u *Update) WithContext(ctx context.Context) *Update {
	if ctx == nil {
		return u
	}

	newUpdate := new(Update)
	*newUpdate = *u

	newUpdate.ctx = ctx

	return newUpdate
}

type botCtxKey struct{}

func (u *Update) WithBot(bot *Bot) *Update {
	if savedBot, ok := u.Context().Value(botCtxKey{}).(*Bot); ok {
		if savedBot == bot {
			return u
		}
	}

	return u.WithContext(context.WithValue(u.Context(), botCtxKey{}, bot))
}

func (u *Update) Bot() *Bot {
	if bot := u.Context().Value(botCtxKey{}); bot != nil {
		return bot.(*Bot)
	}
	return nil
}

type User struct {
	ID int64 `json:"id"`
}
