package types

type Post struct {
	Id           int        `json:"id"`
	OwnerId      int        `json:"owner_id"`
	FromId       int        `json:"from_id"`
	CreatedBy    int        `json:"created_by"`
	Date         int        `json:"date"`
	Text         string     `json:"text"`
	ReplyOwnerId int        `json:"reply_owner_id"`
	ReplyPostId  int        `json:"reply_post_id"`
	FriendsOnly  int        `json:"friends_only"`
	PostType     string     `json:"post_type"`
	PostSource   PostSource `json:"post_source"`
	SignerId     int        `json:"signer_id"`
	CopyHistory  int        `json:"copy_history"` // документация дает информацию, о том что это массив. поскольку конкретная механика этого параметра не указана, здесь возможны ошибки
	CanPin       int        `json:"can_pin"`
	CanDelete    int        `json:"can_delete"`
	CanEdit      int        `json:"can_edit"`
	IsPinned     int        `json:"is_pinned"`
	MarkedAsAds  int        `json:"marked_as_ads"`
	IsFovorite   bool       `json:"is_favorite"`
	Comments     struct {
		Count         int  `json:"count"`
		CanPost       int  `json:"can_post"`
		GroupsCanPost int  `json:"groups_can_post"`
		CanClose      bool `json:"can_close"`
		CanOpen       bool `json:"can_open"`
	} `json:"comments"`
	Likes struct {
		Count      int `json:"count"`
		UserLikes  int `json:"user_likes"`
		CanLike    int `json:"can_like"`
		CanPublish int `json:"can_publish"`
	} `json:"likes"`
	Reposts struct {
		Count        int `json:"count"`
		UserReposted int `json:"user_reposted"`
	} `json:"reposts"`
	Views struct {
		Count int `json:"count"`
	} `json:"views"`
	// Attachments []Attachment `json:"attachments"` TODO: NOT IMPLEMENTED YET!!!
	Geo struct {
		Type        string
		Coordinates string
		place       AttachmentPlace
	} `json:"geo"`
	PostponedId int `json:"postponed_id"` // https://vk.com/dev/groups_events
}
