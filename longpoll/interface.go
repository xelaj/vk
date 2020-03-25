package longpoll

import (
	"context"

	"github.com/xelaj/vk"
	"github.com/xelaj/vk/longpoll/group"
	"github.com/xelaj/vk/longpoll/user"
)

type Bot interface {
	On(event string, function interface{})
	Start(context.Context) chan error
}

func NewUserBot(client *vk.Client, u string) (Bot, error) {
	return user.New(client, u)
}

func NewGroupBot(client *vk.Client, g string) (Bot, error) {
	return group.New(client, g)
}
