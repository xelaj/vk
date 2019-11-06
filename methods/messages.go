package methods

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	kb "github.com/xelaj/vk/keyboard"
	"github.com/xelaj/vk/response"
	types "github.com/xelaj/vk/types"
)

type MessagesSendRequest struct {
	ToId        int
	FromId      int // используется для client.By()
	Message     string
	Attachments []string // пока не знаю как сделать, но было бы правильно передавать таки медиа объект вместо его id
	Keyboard    *kb.Keyboard
}

type MessagesSendResponse int

// TODO: пока сомневаюсь, что инициализация объекта ответа будет правильной, нужны тесты
func MessagesSend(c types.Client, m MessagesSendRequest) (*MessagesSendResponse, error) {
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
	// <<< pre init params

	var (
		res *response.Basic
		err error
	)

	if m.FromId != 0 {
		res, err = c.ByClient(m.FromId).RawMethod("messages.send", params, MessagesSendResponse(0))
	} else {
		res, err = c.RawMethod("messages.send", params, MessagesSendResponse(0))
	}

	if err != nil {
		return nil, err
	}
	return res.Object().(*MessagesSendResponse), nil
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
