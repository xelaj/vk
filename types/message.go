package types

type Message struct {
	Id           int                  `json:"id"`
	Date         int                  `json:"date"`
	PeerId       int                  `json:"peer_id"`
	FromId       int                  `json:"from_id"`
	Text         string               `json:"text"`
	RandomId     int                  `json:"random_id"`
	Ref          string               `json:"ref"`
	RefSource    string               `json:"ref_source"`
	Attachments  []*MessageAttachment `json:"attachments"`
	Important    bool                 `json:"important"`
	Geo          *MessagePlace        `json:"geo"`
	Payload      string               `json:"payload"`
	FwdMessages  []interface{}        `json:"fwd_messages"`  // документация просто не говорит, что там, есть предположение,
	ReplyMessage interface{}          `json:"reply_message"` // что там будут тоже объекты Message? но я не уверен
	Action       struct {
		Type     string `json:"type"`
		MemberId int    `json:"member_id"`
		Text     string `json:"text"`
		Email    string `json:"email"`
		Photo    struct {
			Photo50  string `json:"photo_50"`
			Photo100 string `json:"photo_100"`
			Photo200 string `json:"photo_200"`
		} `json:"photo"`
	} `json:"action"`
}
