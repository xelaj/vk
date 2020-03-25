// this example explain, how to notify you, if somebody wrote you message with bad words

package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/ungerik/go-dry"

	"github.com/k0kubun/pp"
	"github.com/ololosha228/keystorage"
	"github.com/xelaj/vk"
	"github.com/xelaj/vk/longpoll"
	"github.com/xelaj/vk/longpoll/user"
)

func main() {
	client, err := vk.NewClient(vk.ClientConfig{
		ID:         clientID,
		SecureKey:  yourAppSecureKey,
		ServiceKey: yourAppAccessKey,
		KeyStorage: keystorage.NewPrimitive("vk").Set(userNameInStorage, userKey),
	})
	dry.PanicIfErr(err)

	bot, err := longpoll.NewUserBot(client, userNameInStorage)
	dry.PanicIfErr(err)

	bot.On("new_message", func(item *user.NewMessage) error {
		if ContainsBadWords(item.Text) {
			uidStr := item.Attachments["from"].(string)
			uid, _ := strconv.Atoi(uidStr)
			users, err := client.UsersGet([]int{uid}, nil)
			if err != nil {
				return err
			}
			user := (*users)[0]

			fmt.Println(user.FirstName, user.LastName, "wrote bad message!", item.Text)
		}
		return nil
	})

	errchan := bot.Start(context.Background())
	for {
		pp.Println(<-errchan)
	}
}

func ContainsBadWords(text string) bool {
	for _, badWord := range []string{
		"Ублюдок",
		"мать твою",
		"а ну иди сюда",
		"говно собачье",
		"решил ко мне лезть? Ты",
		"засранец вонючий",
		"мать твою",
	} {
		if strings.Contains(text, badWord) {
			return true
		}
	}
	return false
}
