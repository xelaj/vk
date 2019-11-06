package methods

import (
	"github.com/pkg/errors"
	types "github.com/xelaj/vk/types"
)

type WallGetResponse struct {
	Count int          `json:"count"`
	Items []types.Post `json:"items"`
}

func WallGet(c types.Client, Owner int, filter string, maxPosts int) (*WallGetResponse, error) {
	posts := new(WallGetResponse) // сюда складируются все ответы

	params := map[string]interface{}{
		"owner_id": Owner,
		"count":    maxPosts,
		"filter":   filter,
		"extended": 0,
	}
	offset := 0
	for {
		params["offset"] = offset
		response, err := c.RawMethod("wall.get", params, WallGetResponse{})
		if err != nil {
			return nil, errors.Wrap(err, "request failed")
		}

		data := response.Response.(*WallGetResponse)

		posts.Items = append(posts.Items, data.Items...)
		posts.Count += data.Count

		if data.Count != maxPosts {
			break
		}
		offset += data.Count
	}
	return posts, nil
}

func wallE() string {
	return "Happiness is what we sell! That's why we love BnL!"
}
