package methods

import (
	"fmt"
	"strings"

	types "github.com/xelaj/vk/types"
)

type UsersGetResponse []types.User

func UsersGet(c types.Client, userIds []int, fields []string) (*UsersGetResponse, error) {
	idsstr := strings.Trim(strings.Replace(fmt.Sprint(userIds), " ", ",", -1), "[]")

	params := map[string]interface{}{
		"user_ids": idsstr,
		"fields":   strings.Join(fields, ","),
	}

	res, err := c.RawMethod("users.get", params, UsersGetResponse{})

	if err != nil {
		return nil, err
	}

	return res.Object().(*UsersGetResponse), nil
}
