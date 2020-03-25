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
		ID:         6877441,
		SecureKey:  "Q22C8d5u15jzkMWSlnpj",
		ServiceKey: "83e5d90b83e5d90b83e5d90be5838d280a883e583e5d90bdf90b85aca7a73502e4bdd55",
		KeyStorage: keystorage.NewPrimitive("vk").Set("506331906", "0c057f2f4fde05e3fc06b9e62a0d057805118af63392f7426d99b441e46f62cece2f743aa78669c6b1b6f"),
	})
	dry.PanicIfErr(err)

	bot, err := longpoll.NewUserBot(client, "506331906")
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
