package main

import (
	"context"

	"github.com/k0kubun/pp"
	"github.com/ololosha228/keystorage"
	"github.com/ungerik/go-dry"
	"github.com/xelaj/vk"
	"github.com/xelaj/vk/longpoll"
	"github.com/xelaj/vk/longpoll/group"
	"github.com/xelaj/vk/methods"
	"github.com/xelaj/vk/types"
)

func main() {
	client, err := vk.NewClient(vk.ClientConfig{
		ID:         clientID,
		SecureKey:  yourAppSecureKey,
		ServiceKey: yourAppAccessKey,
		KeyStorage: keystorage.NewPrimitive("vk").Set(groupNameInStorage, groupKey),
	})
	dry.PanicIfErr(err)

	bot, err := longpoll.NewGroupBot(client, groupNameInStorage)
	dry.PanicIfErr(err)

	bot.On("message_new", func(item *group.MessageNew) error {
		return proceedMessage(client, item.Message)
	})

	errchan := bot.Start(context.Background())
	for {
		pp.Println(<-errchan)
	}
}

func proceedMessage(c *vk.Client, msg *types.Message) error {
	if msg.Text == "привет" {
		return c.By(groupNameInStorage).MessagesSend(methods.MessagesSendRequest{
			ToId:    msg.FromId,
			Message: "привет! рад тебя видеть!",
		})
	}
	return nil
}
