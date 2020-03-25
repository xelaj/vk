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
		ID:         6877441,
		SecureKey:  "Q22C8d5u15jzkMWSlnpj",
		ServiceKey: "83e5d90b83e5d90b83e5d90be5838d280a883e583e5d90bdf90b85aca7a73502e4bdd55",
		KeyStorage: keystorage.NewPrimitive("vk").Set("-183060742", "7370e31d1d7a6d7830bed6a62125e641d4039e499439763700388d2f98cd856187c408a1e2d3ad0791eaf"),
	})
	dry.PanicIfErr(err)

	bot, err := longpoll.NewGroupBot(client, "-183060742")
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
		return c.By("-183060742").MessagesSend(methods.MessagesSendRequest{
			ToId:    msg.FromId,
			Message: "привет! рад тебя видеть!",
		})
	}
	return nil
}
