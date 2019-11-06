package methods

import (
	"github.com/pkg/errors"
	types "github.com/xelaj/vk/types"
)

const (
	accountGetBannedMaxCount = 10
)

func AccountBan(c types.Client, ownerId int) error {
	params := map[string]interface{}{
		"owner_id": ownerId,
	}
	res, err := c.RawMethod("account.ban", params, 0)
	if err != nil {
		return err
	}
	if *(res.Object().(*int)) != 1 {
		return errors.New("server returns unexpected data")
	}
	return nil
}

type AccountGetBannedResponse struct {
	Count    int           `json:"count"`
	Items    []int         `json:"items"`
	Profiles []*types.User `json:"profiles"`
}

// TODO: написать тесты, я не очень уверен, что оно работает правильно
func AccountGetBanned(c types.Client, offset, count int) (*AccountGetBannedResponse, error) {
	if count < 0 {
		count = 2<<31 - 1
	}

	result := new(AccountGetBannedResponse)
	c.DisableTempTokenDeleting()
	for ; offset < count; offset += accountGetBannedMaxCount {
		params := map[string]interface{}{
			"offset": offset,
			"count":  minimum(count, accountGetBannedMaxCount),
		}

		res, err := c.RawMethod("account.getBanned", params, AccountGetBannedResponse{})
		if err != nil {
			return nil, err
		}

		t := res.Object().(*AccountGetBannedResponse)
		if t.Count < offset {
			break
		}

		result.Count = t.Count
		result.Items = append(result.Items, t.Items...)
		result.Profiles = append(result.Profiles, t.Profiles...)
	}
	c.EnableTempTokenDeleting()
	c.ForceDeleteTempToken()

	return result, nil
}

func AccountGetAppPermissions(c types.Client, ownerId int) (types.Permissions, error) {
	params := map[string]interface{}{}
	if ownerId != 0 {
		params["owner_id"] = ownerId
	}

	res, err := c.RawMethod("account.getAppPermissions", params, types.UserPermissions(0))
	if err != nil {
		return nil, err
	}
	return types.Permissions(*(res.Object().(*types.UserPermissions))), nil

}

func AccountUnban(c types.Client, ownerId int) error {
	params := map[string]interface{}{
		"owner_id": ownerId,
	}
	res, err := c.RawMethod("account.unban", params, 0)
	if err != nil {
		return err
	}
	if *(res.Object().(*int)) != 1 {
		return errors.New("server returns unexpected data")
	}
	return nil
}

/*













 */

func minimum(x, y int) int {
	if x < y {
		return x
	}
	return y
}
