package vk

import (
	"fmt"
	"testing"
)

func TestAccountBan(t *testing.T) {
	client, err := LoadClient("./client.yaml")
	if err != nil {
		t.Error(err)
	}

	res, err := client.By(506331906).AccountGetAppPermissions(0)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res.Type())
	client.SaveTo("")

}
