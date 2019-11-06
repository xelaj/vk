package types

type PostSource struct {
	Type     string `json:"type"`
	Platform string `json:"platform"`
	Data     string `json:"data"`
	Url      string `json:"url"`
}
