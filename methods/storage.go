package methods

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/xelaj/vk/types"
)

func StorageGet(c types.Client, keys []string, userId int) (map[string]string, error) {
	params := map[string]interface{}{
		"keys": strings.Join(keys, ","),
	}
	if userId != 0 {
		params["user_id"] = userId
	}

	type kv struct {
		K string `json:"key"`
		V string `json:"value"`
	}

	res, err := c.RawMethod("storage.get", params, []kv{})
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)

	kvs := res.Object().(*[]kv)

	for _, item := range *kvs {
		if item.V == "" {
			continue
		}
		result[item.K] = item.V
	}

	return result, nil

}

func StorageGetKeys(c types.Client) {

}

func StorageSet(c types.Client, key string, value interface{}, userId int) error {
	params := map[string]interface{}{
		"key":   key,
		"value": fmt.Sprint(value),
	}
	if userId != 0 {
		params["user_id"] = userId
	}

	res, err := c.RawMethod("storage.set", params, 0)
	if err != nil {
		return err
	}

	if *(res.Object().(*int)) != 1 {
		return errors.New("strange data was returned (not 1)")
	}

	return nil
}
