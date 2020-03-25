package methods

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	kb "github.com/xelaj/vk/keyboard"
	types "github.com/xelaj/vk/types"
)

type MessagesSendRequest struct {
	ToId        int
	Message     string
	Attachments []string // пока не знаю как сделать, но было бы правильно передавать таки медиа объект вместо его id
	Keyboard    *kb.Keyboard
}

// TODO: пока сомневаюсь, что инициализация объекта ответа будет правильной, нужны тесты
func MessagesSend(c types.Client, m MessagesSendRequest) error {
	rand.Seed(time.Now().Unix())
	// >>> pre init params
	params := map[string]interface{}{
		"user_id":   m.ToId,
		"message":   m.Message,
		"random_id": rand.Uint32(),
	}
	if len(m.Attachments) != 0 {
		params["attachments"] = strings.Join(m.Attachments, ",")
	}

	if m.Keyboard != nil {
		params["keyboard"] = m.Keyboard.String()
	}

	_, err := c.RawMethod("messages.send", params, 0)

	return err
}

type MessagesGetByIdResponse struct {
	Count int             `json:"count"`
	Items []types.Message `json:"items"`
}

func MessagesGetById(c types.Client, messageIds []int, previewLen, groupId int) (*MessagesGetByIdResponse, error) {
	idsstr := strings.Trim(strings.Replace(fmt.Sprint(messageIds), " ", ",", -1), "[]")
	params := map[string]interface{}{
		"message_ids":    idsstr,
		"preview_length": previewLen,
		"group_id":       groupId,
	}
	res, err := c.RawMethod("messages.getById", params, MessagesGetByIdResponse{})
	if err != nil {
		return nil, err
	}
	return res.Object().(*MessagesGetByIdResponse), nil
}

type MessagesGetConversationsRequest struct {
	Offset       int
	Count        int    // default 20
	Filter       string // all unread important unanswered message_request business_notify
	Extended     bool
	StartMessage int
	Fields       []string
	GroupID      int
	MajorSortID  int
}

type MessagesGetConversationsResponse struct {
	Count int `json:"count"`
	Items []struct {
		Conversation *types.Conversation `json:"conversation"`
		LastMessage  *types.Message      `json:"last_message"`
	} `json:"items"`
	UnreadCount int            `json:"unread_count"`
	Profiles    []*types.User  `json:"profiles"`
	Groups      []*types.Group `json:"groups"`
}

func MessagesGetConversations(c types.Client, r MessagesGetConversationsRequest) (*MessagesGetConversationsResponse, error) {
	params := map[string]interface{}{}
	if r.Offset != 0 {
		params["offset"] = r.Offset
	}
	if r.Count != 0 {
		params["count"] = r.Count
	}
	if r.Filter != "" {
		params["filter"] = r.Filter
	}
	if r.Extended {
		params["extended"] = true
	}
	if r.StartMessage != 0 {
		params["start_message_id"] = r.StartMessage
	}
	if len(r.Fields) != 0 {
		params["fields"] = strings.Join(r.Fields, ",")
	}
	if r.GroupID != 0 {
		params["group_id"] = r.GroupID
	}
	if r.MajorSortID != 0 {
		params["major_sort_id"] = r.MajorSortID
	}

	res, err := c.RawMethod("messages.getConversations", params, MessagesGetConversationsResponse{})
	if err != nil {
		return nil, err
	}
	return res.Object().(*MessagesGetConversationsResponse), nil
}

type MessagesGetLongPollServerResponse struct {
	Key    string `json:"key"`
	Server string `json:"server"`
	Ts     int    `json:"ts"` // хз почему string, но так в доках указано
	Pts    int    `json:"pts"`
}

func MessagesGetLongPollServer(c types.Client, needPts bool, groupID int) (*MessagesGetLongPollServerResponse, error) {
	params := map[string]interface{}{
		"lp_version": 3,
	}
	if needPts {
		params["need_pts"] = true
	}
	if groupID != 0 {
		params["group_id"] = groupID
	}
	res, err := c.RawMethod("messages.getLongPollServer", params, MessagesGetLongPollServerResponse{})
	if err != nil {
		return nil, err
	}
	return res.Object().(*MessagesGetLongPollServerResponse), nil
}
