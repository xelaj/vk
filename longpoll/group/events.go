package group

import (
	"reflect"

	"github.com/xelaj/vk/types"
)

type MessageNew struct {
	Message    *types.Message    `json:"message"`
	ClientInfo *types.ClientInfo `json:"client_info"`
}
type MessageReply = types.Message
type MessageEdit = types.Message

type MessageAllow struct {
	UserId int    `json:"user_id"`
	Key    string `json:"key"`
}

type MessageDeny struct {
	UserId int `json:"user_id"`
}

type PhotoNew = types.Photo

type PhotoCommentNew = types.Comment
type PhotoCommentEdit = types.Comment
type PhotoCommentRestore = types.Comment

type PhotoCommentDelete struct {
	OwnerId   int `json:"owner_id"`
	Id        int `json:"id"`
	UserId    int `json:"user_id"`
	DeleterId int `json:"deleter_id"`
	PhotoId   int `json:"photo_id"`
}

type AudioNew = types.Audio
type VideoNew = types.Video
type VideoCommentNew = types.Comment
type VideoCommentEdit = types.Comment
type VideoCommentRestore = types.Comment

type VideoCommentDelete struct {
	OwnerId   int `json:"owner_id"`
	Id        int `json:"id"`
	UserId    int `json:"user_id"`
	DeleterId int `json:"deleter_id"`
	ViedoId   int `json:"video_id"`
}

type WallPostNew = types.Post
type WallRepost = types.Post

type WallReplyNew = types.Comment
type WallReplyEdit = types.Comment
type WallReplyRestore = types.Comment

type WallReplyDelete struct {
	OwnerId   int `json:"owner_id"`
	Id        int `json:"id"`
	UserId    int `json:"user_id"`
	DeleterId int `json:"deleter_id"`
	PostId    int `json:"post_id"`
}

type BoardPostNew = types.Post
type BoardPostEdit = types.Post
type BoardPostRestore = types.Post

type BoardPostDelete struct {
	OwnerId int `json:"topic_owner_id"`
	Id      int `json:"id"`
	TopicId int `json:"topic_id"`
}

type MarketCommentNew = types.Comment
type MarketCommentEdit = types.Comment
type MarketCommentRestore = types.Comment

type MarketCommentDelete struct {
	OwnerId   int `json:"owner_id"`
	Id        int `json:"id"`
	UserId    int `json:"user_id"`
	DeleterId int `json:"deleter_id"`
	ItemId    int `json:"item_id"`
}

type GroupLeave struct {
	UserId int `json:"user_id"`
	Self   int `json:"self"`
}

type GroupJoin struct {
	UserId   int    `json:"user_id"`
	JoinType string `json:"join_type"`
}

type UserBlock struct {
	AdminId     int    `json:"admin_id"`
	UserId      int    `json:"user_id"`
	UnblockDate int    `json:"unblock_date"`
	Reason      int    `json:"reason"`
	Comment     string `json:"comment"`
}

type UserUnblock struct {
	AdminId   int `json:"admin_id"`
	UserId    int `json:"user_id"`
	ByEndDate int `json:"by_end_date"`
}

type PollVoteNew struct {
	OwnerId  int `json:"owner_id"`
	PollId   int `json:"poll_id"`
	OptionId int `json:"option_id"`
	UserId   int `json:"user_id"`
}

type GroupOfficersEdit struct {
	AdminId  int `json:"admin_id"`
	UserId   int `json:"user_id"`
	LevelOld int `json:"level_old"`
	LevelNew int `json:"level_new"`
}

type MessageTypingState struct {
	State  string `json:"state"`
	FromId int    `json:"from_id"`
	ToId   int    `json:"to_id"`
}

type GroupChangeSettings struct {
	UserId  int
	Changes struct {
		Title               ValueChanging `json:"title"`
		Description         ValueChanging `json:"description"`
		Access              ValueChanging `json:"access"`
		ScreenName          ValueChanging `json:"screen_name"`
		PublicCategory      ValueChanging `json:"public_category"`
		PublicSubcategory   ValueChanging `json:"public_subcategory"`
		AgeLimits           ValueChanging `json:"age_limits"`
		Website             ValueChanging `json:"website"`
		EnableStatusDefault ValueChanging `json:"enable_status_default"`
		EnableAudio         ValueChanging `json:"enable_audio"`
		EnablePhoto         ValueChanging `json:"enable_photo"`
		EnableVideo         ValueChanging `json:"enable_video"`
		EnableMarket        ValueChanging `json:"enable_market"`
	}
}

type GroupChangePhoto struct {
	UserId int         `json:"user_id"`
	Photo  types.Photo `json:"photo"`
}

type ValueChanging struct {
	OldValue interface{} `json:"old_value"`
	NewValue interface{} `json:"new_value"` // документация не объясняет, какой типа там хранится
}

var (
	UpdateTypes = map[string]reflect.Type{
		"message_new":            reflect.TypeOf(MessageNew{}),
		"message_reply":          reflect.TypeOf(MessageReply{}),
		"message_edit":           reflect.TypeOf(MessageEdit{}),
		"message_allow":          reflect.TypeOf(MessageAllow{}),
		"photo_new":              reflect.TypeOf(PhotoNew{}),
		"photo_comment_new":      reflect.TypeOf(PhotoCommentNew{}),
		"photo_comment_edit":     reflect.TypeOf(PhotoCommentEdit{}),
		"photo_comment_restore":  reflect.TypeOf(PhotoCommentRestore{}),
		"photo_comment_delete":   reflect.TypeOf(PhotoCommentDelete{}),
		"audio_new":              reflect.TypeOf(AudioNew{}),
		"video_new":              reflect.TypeOf(VideoNew{}),
		"video_comment_new":      reflect.TypeOf(VideoCommentNew{}),
		"video_comment_edit":     reflect.TypeOf(VideoCommentEdit{}),
		"video_comment_restore":  reflect.TypeOf(VideoCommentRestore{}),
		"video_comment_delete":   reflect.TypeOf(VideoCommentDelete{}),
		"wall_post_new":          reflect.TypeOf(WallPostNew{}),
		"wall_repost":            reflect.TypeOf(WallRepost{}),
		"wall_reply_new":         reflect.TypeOf(WallReplyNew{}),
		"wall_reply_edit":        reflect.TypeOf(WallReplyEdit{}),
		"wall_reply_restore":     reflect.TypeOf(WallReplyRestore{}),
		"wall_reply_delete":      reflect.TypeOf(WallReplyDelete{}),
		"board_post_new":         reflect.TypeOf(BoardPostNew{}),
		"board_post_edit":        reflect.TypeOf(BoardPostEdit{}),
		"board_post_restore":     reflect.TypeOf(BoardPostRestore{}),
		"board_post_delete":      reflect.TypeOf(BoardPostDelete{}),
		"market_comment_new":     reflect.TypeOf(MarketCommentNew{}),
		"market_comment_edit":    reflect.TypeOf(MarketCommentEdit{}),
		"market_comment_restore": reflect.TypeOf(MarketCommentRestore{}),
		"market_comment_delete":  reflect.TypeOf(MarketCommentDelete{}),
		"group_leave":            reflect.TypeOf(GroupLeave{}),
		"group_join":             reflect.TypeOf(GroupJoin{}),
		"user_block":             reflect.TypeOf(UserBlock{}),
		"user_unblock":           reflect.TypeOf(UserUnblock{}),
		"poll_vote_new":          reflect.TypeOf(PollVoteNew{}),
		"group_officers_edit":    reflect.TypeOf(GroupOfficersEdit{}),
		"group_change_settings":  reflect.TypeOf(GroupChangeSettings{}),
		"group_change_photo":     reflect.TypeOf(GroupChangePhoto{}),
		"message_typing_state":   reflect.TypeOf(MessageTypingState{}),
	}
)

func validAction(name string) bool {
	_, ok := UpdateTypes[name]
	return ok
}
