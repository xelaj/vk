package user

import (
	"reflect"
	"time"
)

var (
	userUpdateTypes = map[int]reflect.Type{
		1:   reflect.TypeOf(MessageFlagsChange{}),
		2:   reflect.TypeOf(MessageFlagsSet{}),
		3:   reflect.TypeOf(MessageFlagsUnset{}),
		4:   reflect.TypeOf(NewMessage{}),
		5:   reflect.TypeOf(EditMessage{}),
		6:   reflect.TypeOf(IncommingMessageRead{}),
		7:   reflect.TypeOf(OutgoingMessageRead{}),
		8:   reflect.TypeOf(FriendIsOnline{}),
		9:   reflect.TypeOf(FriendIsOffline{}),
		10:  reflect.TypeOf(DialogFlagsUnset{}),
		11:  reflect.TypeOf(DialogFlagsChange{}),
		12:  reflect.TypeOf(DialogFlagsSet{}),
		13:  reflect.TypeOf(DialogMessageDelete{}),
		14:  reflect.TypeOf(DialogMessageRestore{}),
		51:  reflect.TypeOf(DialogParamsChanged{}),
		52:  reflect.TypeOf(DialogInfoChanged{}),
		61:  reflect.TypeOf(UserWritingDialog{}),
		62:  reflect.TypeOf(UserWritingChat{}),
		63:  reflect.TypeOf(UserWritingChatMultiple{}),
		64:  reflect.TypeOf(UserWritingAudio{}),
		70:  reflect.TypeOf(UserMadeCall{}),
		80:  reflect.TypeOf(CounterChanged{}),
		114: reflect.TypeOf(NotificationSettingsChanged{}),
	}
	userUpdateTypesStrings = map[string]int{
		"message_flags_change":          1,
		"message_flags_set":             2,
		"message_flags_unset":           3,
		"new_message":                   4,
		"edit_message":                  5,
		"incomming_message_read":        6,
		"outgoing_message_read":         7,
		"friend_is_online":              8,
		"friend_is_offline":             9,
		"dialog_flagsUnset":             10,
		"dialog_flagsChange":            11,
		"dialog_flagsSet":               12,
		"dialog_messageDelete":          13,
		"dialog_messageRestore":         14,
		"dialog_paramsChanged":          51,
		"dialog_infoChanged":            52,
		"user_writing_dialog":           61,
		"user_writing_chat":             62,
		"user_writing_chat_multiple":    63,
		"user_writing_audio":            64,
		"user_made_call":                70,
		"counter_changed":               80,
		"notification_settings_changed": 114,
	}
)

type Flags uint

type MessageFlagsChange struct {
	MessageID int
	Flags     Flags
}

func (m *MessageFlagsChange) Code() int {
	return 1
}

type MessageFlagsSet struct {
	MessageID int
	Mask      Flags
	PeerID    int
}

func (m *MessageFlagsSet) Code() int {
	return 2
}

type MessageFlagsUnset struct {
	MessageID int
	Mask      Flags
}

func (m *MessageFlagsUnset) Code() int {
	return 3
}

type NewMessage struct {
	MessageID   int
	Mask        Flags
	PeerID      int
	Time        time.Time
	Text        string
	Attachments map[string]interface{}
}

func (m *NewMessage) Code() int {
	return 4
}

type EditMessage struct {
	MessageID   int
	Mask        Flags
	PeerID      int
	Time        time.Time
	Text        string
	Attachments map[string]interface{}
}

func (m *EditMessage) Code() int {
	return 5
}

type IncommingMessageRead struct {
	PeerID  int
	LocalID int
}

func (m *IncommingMessageRead) Code() int {
	return 6
}

type OutgoingMessageRead struct {
	PeerID  int
	LocalID int
}

func (m *OutgoingMessageRead) Code() int {
	return 7
}

type FriendIsOnline struct {
	UserID        int `invert`
	Extra         int
	Time          time.Time
	UnknownField1 int
	UnknownField2 int
	UnknownField3 int
}

func (m *FriendIsOnline) Code() int {
	return 8
}

type FriendIsOffline struct {
	UserID        int `invert`
	Flags         Flags
	Time          time.Time
	UnknownField1 int
	UnknownField2 int
	UnknownField3 int
}

func (m *FriendIsOffline) Code() int {
	return 9
}

type DialogFlagsUnset struct {
	PeerID int
	Flags  Flags
}

func (m *DialogFlagsUnset) Code() int {
	return 10
}

type DialogFlagsChange struct {
	PeerID int
	Flags  Flags
}

func (m *DialogFlagsChange) Code() int {
	return 11
}

type DialogFlagsSet struct {
	PeerID int
	Flags  Flags
}

func (m *DialogFlagsSet) Code() int {
	return 12
}

type DialogMessageDelete struct {
	PeerID  int
	LocalID int
}

func (m *DialogMessageDelete) Code() int {
	return 13
}

type DialogMessageRestore struct {
	PeerID  int
	LocalID int
}

func (m *DialogMessageRestore) Code() int {
	return 14
}

type DialogParamsChanged struct {
	ChatID int
	Self   bool
}

func (m *DialogParamsChanged) Code() int {
	return 51
}

type DialogInfoChanged struct {
	TypeID int
	PeerID int
	Info   int
}

func (m *DialogInfoChanged) Code() int {
	return 52
}

type UserWritingDialog struct {
	UserID int
	Flags  Flags
}

func (m *UserWritingDialog) Code() int {
	return 61
}

type UserWritingChat struct {
	UserID int
	ChatID int
	// unknown integer
}

func (m *UserWritingChat) Code() int {
	return 62
}

type UserWritingChatMultiple struct {
	UserIDs    []int
	PeerID     int
	TotalCount int
	Time       time.Time
}

func (m *UserWritingChatMultiple) Code() int {
	return 63
}

type UserWritingAudio struct {
	UserIDs    []int
	PeerID     int
	TotalCount int
	Time       time.Time
}

func (m *UserWritingAudio) Code() int {
	return 64
}

type UserMadeCall struct {
	UserID int
	CallID int
}

func (m *UserMadeCall) Code() int {
	return 70
}

type CounterChanged struct {
	Count        int
	UnknownField int
}

func (m *CounterChanged) Code() int {
	return 80
}

type NotificationSettingsChanged struct {
	PeerID        int
	Unmuted       bool
	DisabledUntil int
}

func (m *NotificationSettingsChanged) Code() int {
	return 114
}
