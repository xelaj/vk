package methods

import (
	types "github.com/xelaj/vk/types"
)

type AppsGetResponse struct {
	Count int         `json:"count"`
	Items []types.App `json:"items"`
}

func AppsGet(c types.Client, additional map[string]interface{}) (*AppsGetResponse, error) {
	resp, err := c.RawMethod("apps.get", additional, AppsGetResponse{})
	if err != nil {
		return nil, err
	}
	return resp.Object().(*AppsGetResponse), nil
}
