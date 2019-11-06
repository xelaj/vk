package main

import (
	"github.com/ungerik/go-dry"
	"github.com/xelaj/vk"
	"github.com/xelaj/vk/longpoll"
	"github.com/xelaj/vk/methods"
	"github.com/xelaj/vk/types"
)

const (
	yourAppID         = 0
	yourAppSecret     = ""
	yourAppServiceKey = ""
	groupID           = 0
	groupToken        = ""
)

var (
	client *vk.Client
)

func main() {
	var err error

	if dry.FileExists("./config.yml") {
		client, err = vk.LoadClient("./config.yml")
		if err != nil {
			panic(err)
		}
	} else {
		client, err = vk.NewClient(yourAppID, yourAppSecret, yourAppServiceKey)
		if err != nil {
			panic(err)
		}

		err = client.AuthGroup(groupID, groupToken)
		if err != nil {
			panic(err)
		}
	}

	client.SaveTo("./config.yml")

	bot, err := longpoll.NewGroupBot(client, 184456315)
	if err != nil {
		panic(err)
	}

	bot.On("message_new", OnMessage)
	var x bool
	bot.Start(&x) // can be used as goroutine
}

func OnMessage(payload longpoll.Update) error {
	msg := payload.Object.(*types.Message)

	client.By(-184456315).MessagesSend(methods.MessagesSendRequest{
		ToId:    msg.FromId,
		Message: "ты написал: " + msg.Text,
	})
	return nil
}
