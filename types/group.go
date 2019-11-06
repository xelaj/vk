package types

type Group struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	ScreenName   string `json:"screen_name"`
	IsClosed     int    `json:"is_closed"`
	Deactivated  string `json:"deactivated"`
	IsAdmin      int    `json:"is_admin"`
	AdminLevel   int    `json:"admin_level"`
	IsMember     int    `json:"is_member"`
	IsAdvertiser int    `json:"is_advertiser"`
	InvitedBy    int    `json:"invited_by"`
	Type         string `json:"type"`
	Photo50      string `json:"photo_50"`
	Photo100     string `json:"photo_100"`
	Photo200     string `json:"photo_200"`
}
