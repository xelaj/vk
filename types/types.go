package types

import "github.com/xelaj/vk/response"

type Client interface {
	RawMethod(method string, params map[string]interface{}, storeTo interface{}) (*response.Basic, error)
	ByClient(int) Client
	DisableTempTokenDeleting()
	EnableTempTokenDeleting()
	ForceDeleteTempToken()
}

type AttachmentPlace struct {
	Id        int
	Title     string
	Latitude  int
	Longitude int
	Created   int
	Icon      string
	Checkins  int
	Updated   int
	Type      int
	Country   int
	City      int
	Address   string
}

// TODO: возможно, я чего-то не понял, но в разных объектах, разные типы геоданных.
//       как-нибудь потом разберусь
type MessagePlace struct {
	Type        string
	Coordinates struct {
		Latitude  float64
		Longitude float64
	}
	Place struct {
		Id        int
		Title     string
		Latitude  float64
		Longitude float64
		Created   int
		Icon      string
		Country   string
		City      string
	}
}
