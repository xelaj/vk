package methods

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/xelaj/vk/types"
)

type GroupsGetResponse struct {
	Count int           `json:"count"`
	Items []types.Group `json:"items"`
}

func GroupsGet(c types.Client, userId int, filter, fields []string, groupsMaxCount int) (*GroupsGetResponse, error) {
	groups := new(GroupsGetResponse) // сюда складируются все ответы

	params := map[string]interface{}{
		"user_id":  userId,
		"extended": 1,
		"filter":   strings.Join(filter, ","),
		"fields":   strings.Join(fields, ","),
		"count":    groupsMaxCount,
	}

	offset := 0
	for {
		params["offset"] = offset
		response, err := c.RawMethod("groups.get", params, GroupsGetResponse{})
		if err != nil {
			return nil, err
		}

		d := response.Object().(*GroupsGetResponse)

		groups.Items = append(groups.Items, d.Items...)
		groups.Count += len(d.Items)

		if len(d.Items) != groupsMaxCount {
			break
		}
		offset += len(d.Items)
	}
	return groups, nil
}

type GroupsGetByIdResponse []types.Group

func GroupsGetById(c types.Client, groupIds []int, fields []string, groupsMaxCount int) (*GroupsGetByIdResponse, error) {
	idsstr := strings.Trim(strings.Replace(fmt.Sprint(groupIds), " ", ",", -1), "[]")

	params := map[string]interface{}{
		"group_ids": idsstr,
		"fields":    strings.Join(fields, ","),
	}

	res, err := c.RawMethod("groups.getbyid", params, GroupsGetByIdResponse{})
	if err != nil {
		return nil, err
	}

	return res.Object().(*GroupsGetByIdResponse), nil
}

type GroupsGetLongPollServerResponse struct {
	Key    string `json:"key"`
	Server string `json:"server"`
	Ts     string `json:"ts"` // хз почему string, но так в доках указано
}

// я не знаю как это работает, но ЗДЕСЬ groupId
// должен быть положительным. не знаю почему, но вот так.
func GroupsGetLongPollServer(c types.Client, groupId int) (*GroupsGetLongPollServerResponse, error) {
	params := map[string]interface{}{
		"group_id": groupId,
	}

	res, err := c.RawMethod("groups.getLongPollServer", params, GroupsGetLongPollServerResponse{})
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	return res.Object().(*GroupsGetLongPollServerResponse), nil
}

func GroupsGetTokenPermissions(c types.Client) (types.CommunityPermissions, error) {
	params := map[string]interface{}{}

	res, err := c.RawMethod("groups.getTokenPermissions", params, types.CommunityPermissions(0))
	if err != nil {
		return 0, errors.Wrap(err, "request failed")
	}

	return *(res.Object().(*types.CommunityPermissions)), nil

}
