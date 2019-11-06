package types

type Message struct {
	Id           int
	Date         int
	PeerId       int
	FromId       int    `json:"from_id"`
	Text         string `json:"text"`
	RandomId     int
	Ref          string
	RefSource    string
	Attachments  []MessageAttachment
	Important    bool
	Geo          MessagePlace
	Payload      string
	FwdMessages  []interface{} // документация просто не говорит, что там, есть предположение,
	ReplyMessage interface{}   // что там будут тоже объекты Message? но я не уверен
	Action       struct {
		Type     string
		MemberId int
		Text     string
		email    string
		Photo    struct {
			Photo50  string
			Photo100 string
			Photo200 string
		}
	}
}
