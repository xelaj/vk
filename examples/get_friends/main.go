package main

import (
	"fmt"

	"github.com/ololosha228/keystorage"
	"github.com/ungerik/go-dry"
	"github.com/xelaj/vk"
)

func main() {
	client, err := vk.NewClient(vk.ClientConfig{
		ID:         clientID,
		SecureKey:  yourAppSecureKey,
		ServiceKey: yourAppAccessKey,
		KeyStorage: keystorage.NewPrimitive("vk").Set(userNameInStorage, userKey),
	})
	dry.PanicIfErr(err)

	resp, err := client.FriendsGet(userId, "", []string{"online"}, "")
	dry.PanicIfErr(err)

	fmt.Println("Friends are online:")
	for _, item := range resp.Items {
		if item.Online == 1 {
			fmt.Println(item.FirstName, item.LastName)
		}
	}
}
