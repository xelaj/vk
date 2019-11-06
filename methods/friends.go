package methods

import (
	"strings"

	"github.com/pkg/errors"
	types "github.com/xelaj/vk/types"
)

type FriendsGetResponse struct {
	Count int          `json:"count"`
	Items []types.User `json:"items"`
}

type FriendsGetIdResponse struct {
	Count int   `json:"count"`
	Items []int `json:"items"`
}

// TODO: в документации в качестве параметра указан некий ref, о котором вообще нет никакой информации.
//       list_id тоже непонятно зачем нужен.
//       попытаться разобраться, и впихнуть
func FriendsGet(c types.Client, userId int, Order string, fields []string, nameCase string, friendsRequestCount int) (*FriendsGetResponse, error) {
	if len(fields) == 0 {
		return nil, errors.New("client.FriendsGet must have return fields. if you need only user id's, use client.FriendsGetId")
	}

	friends := new(FriendsGetResponse)

	params := map[string]interface{}{
		"user_id":   userId,
		"order":     Order,
		"count":     friendsRequestCount,
		"fields":    strings.Join(fields, ","),
		"name_case": nameCase,
	}
	offset := 0
	for {
		params["offset"] = offset
		res, err := c.RawMethod("friends.get", params, FriendsGetResponse{})
		if err != nil {
			return nil, err
		}

		d := res.Object().(*FriendsGetResponse)

		friends.Items = append(friends.Items, d.Items...)
		friends.Count += len(d.Items)

		if len(d.Items) != friendsRequestCount {
			break
		}
		offset += len(d.Items)
	}
	return friends, nil
}

func FriendsGetId(c types.Client, userId int, Order string, nameCase string, friendsRequestCount int) (*FriendsGetIdResponse, error) {
	friends := new(FriendsGetIdResponse)

	params := map[string]interface{}{
		"user_id":   userId,
		"order":     Order,
		"count":     friendsRequestCount,
		"name_case": nameCase,
	}
	offset := 0
	for {
		params["offset"] = offset
		res, err := c.RawMethod("friends.get", params, FriendsGetIdResponse{})
		if err != nil {
			return nil, err
		}

		d := res.Object().(*FriendsGetIdResponse)

		friends.Items = append(friends.Items, d.Items...)
		friends.Count += len(d.Items)

		if len(d.Items) != friendsRequestCount {
			break
		}
		offset += len(d.Items)
	}
	return friends, nil
}
