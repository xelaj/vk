package types

import (
	"time"

	"github.com/xelaj/vk/response"
)

type Conversation struct {
	Peer struct {
		ID      int    `json:"id"`
		Type    string `json:"type"`
		LocalID int    `json:"local_id"` // TODO: не понятно как это работает, наверное нужно какую-то специфическую обертку делать
	} `json:"peer"`
	InRead       int  `json:"in_read"`
	OutRead      int  `json:"out_read"`
	UnreadCount  int  `json:"unread_count"`
	Important    bool `json:"important"`
	Unanswered   bool `json:"unanswered"`
	PushSettings struct {
		DisabledUntil   *time.Time `json:"disable_until" unmarshal:"parse_time"`
		DisabledForever bool       `json:"disable_forever"`
		Muted           bool       `json:"no_sound"`
	} `json:"push_settings"`
	CanWrite struct {
		Allowed bool               `json:"allowed"`
		Reason  response.ErrorCode `json:"reason"`
	} `json:"can_write"`
	ChatSettings struct {
		MembersCount  int      `json:"members_count"`
		Title         string   `json:"title"`
		PinnedMessage *Message `json:"pinned_message"`
		State         string   `json:"state"`
		Photo         struct {
			Photo50  string `json:"photo_50"`
			Photo100 string `json:"photo_100"`
			Photo200 string `json:"photo_200"`
		} `json:"photo"`
		ActiveIDs      []int `json:"active_ids"`
		IsGroupChannel bool  `json:"is_group_channel"`
	} `json:"chat_settings"`
}
