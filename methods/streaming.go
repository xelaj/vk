package methods

import types "github.com/xelaj/vk/types"

type StreamingGetServerUrlResponse struct {
	Endpoint string `json:"endpoint"`
	Key      string `json:"key"`
}

func StreamingGetServerUrl(c types.Client) (*StreamingGetServerUrlResponse, error) {
	res, err := c.RawMethod("streaming.getServerUrl", nil, StreamingGetServerUrlResponse{})
	return res.Object().(*StreamingGetServerUrlResponse), err
}

//func StreamingGetSettings(c types.Client) {
//
//}

//func StreamingGetStats(c types.Client) {
//
//}

//func StreamingGetStem(c types.Client) {
//
//}

//func StreamingSetSettings(c types.Client) {
//
//}
