// для тех кому не нравится струткур прав доступа и хочет переделать:
// https://vk.com/dev/permissions в документации четко написано, что для пабликов и пользователей разные права
// с разной битовой маской. если есть варианты, как лучше описать права доступа, сделайте pull request, я сам от
// этого кода не в восторге

package types

type Permissions interface {
	Type() TokenType
	Bitmask() int

	Notify() bool
	Friends() bool
	Photos() bool
	Audio() bool
	Video() bool
	Stories() bool
	Pages() bool
	LeftLinks() bool
	Status() bool
	Notes() bool
	Messages() bool
	Wall() bool
	Ads() bool
	Offline() bool
	Docs() bool
	Groups() bool
	Notifications() bool
	Stats() bool
	Email() bool
	Market() bool
	AppWidget() bool
	Manage() bool
}

type TokenType int

const (
	UserToken TokenType = iota + 1
	CommunityToken
)

func (t TokenType) String() string {
	if t == UserToken {
		return "<vk_token: user>"
	}
	if t == CommunityToken {
		return "<vk_token: community>"
	}
	return "<vk_token: unknown>"
}

type UserPermissions int

const (
	UPermNotify UserPermissions = 1 << iota
	UPermFriends
	UPermPhotos
	UPermAudio
	UPermVideo
	_
	UPermStories
	UPermPages
	UPermLeftLinks
	_
	UPermStatus
	UPermNotes
	UPermMessages
	UPermWall
	_
	UPermAds
	UPermOffline
	UPermDocs
	UPermGroups
	UPermNotifications
	UPermStats
	_
	UPermEmail
	_
	_
	_
	_
	_
	UPermMarket
)

// type Permissions int

// const (
// 	PermNotify Permissions = 1 << iota
// 	PermFriends
// 	PermPhotos
// 	PermAudio
// 	PermVideo
// 	_
// 	PermStories
// 	PermPages
// 	PermLeftLinks
// 	_
// 	PermStatus
// 	PermNotes
// 	PermMessages
// 	PermWall
// 	_
// 	PermAds
// 	PermOffline
// 	PermDocs
// 	PermGroups
// 	PermNotifications
// 	PermStats
// 	_
// 	PermEmail
// 	_
// 	_
// 	_
// 	_
// 	_
// 	PermMarket
// )

func (u UserPermissions) Type() TokenType {
	return UserToken
}

func (u UserPermissions) Bitmask() int {
	return int(u)
}

func (u UserPermissions) Notify() bool {
	return u&UPermNotify != 0
}
func (u UserPermissions) Friends() bool {
	return u&UPermFriends != 0
}
func (u UserPermissions) Photos() bool {
	return u&UPermPhotos != 0
}
func (u UserPermissions) Audio() bool {
	return u&UPermAudio != 0
}
func (u UserPermissions) Video() bool {
	return u&UPermVideo != 0
}
func (u UserPermissions) Stories() bool {
	return u&UPermStories != 0
}
func (u UserPermissions) Pages() bool {
	return u&UPermPages != 0
}
func (u UserPermissions) LeftLinks() bool {
	return u&UPermLeftLinks != 0
}
func (u UserPermissions) Status() bool {
	return u&UPermStatus != 0
}
func (u UserPermissions) Notes() bool {
	return u&UPermNotes != 0
}
func (u UserPermissions) Messages() bool {
	return u&UPermMessages != 0
}
func (u UserPermissions) Wall() bool {
	return u&UPermWall != 0
}
func (u UserPermissions) Ads() bool {
	return u&UPermAds != 0
}
func (u UserPermissions) Offline() bool {
	return u&UPermOffline != 0
}
func (u UserPermissions) Docs() bool {
	return u&UPermDocs != 0
}
func (u UserPermissions) Groups() bool {
	return u&UPermGroups != 0
}
func (u UserPermissions) Notifications() bool {
	return u&UPermNotifications != 0
}
func (u UserPermissions) Stats() bool {
	return u&UPermStats != 0
}
func (u UserPermissions) Email() bool {
	return u&UPermEmail != 0
}
func (u UserPermissions) Market() bool {
	return u&UPermMarket != 0
}
func (u UserPermissions) AppWidget() bool {
	return false
}
func (u UserPermissions) Manage() bool {
	return false
}

type CommunityPermissions int

const (
	CPermStories CommunityPermissions = 1 << iota
	_
	CPermPhotos
	_
	_
	_
	CPermAppWidget
	_
	_
	_
	_
	_
	CPermMessages
	_
	_
	_
	_
	CPermDocs
	CPermManage
)

func (c CommunityPermissions) Type() TokenType {
	return CommunityToken
}

func (c CommunityPermissions) Bitmask() int {
	return int(c)
}

func (c CommunityPermissions) Notify() bool {
	return false
}
func (c CommunityPermissions) Friends() bool {
	return false
}
func (c CommunityPermissions) Photos() bool {
	return c&CPermPhotos != 0
}
func (c CommunityPermissions) Audio() bool {
	return false
}
func (c CommunityPermissions) Video() bool {
	return false
}
func (c CommunityPermissions) Stories() bool {
	return c&CPermStories != 0
}
func (c CommunityPermissions) Pages() bool {
	return false
}
func (c CommunityPermissions) LeftLinks() bool {
	return false
}
func (c CommunityPermissions) Status() bool {
	return false
}
func (c CommunityPermissions) Notes() bool {
	return false
}
func (c CommunityPermissions) Messages() bool {
	return c&CPermMessages != 0
}
func (c CommunityPermissions) Wall() bool {
	return false
}
func (c CommunityPermissions) Ads() bool {
	return false
}
func (c CommunityPermissions) Offline() bool {
	return false
}
func (c CommunityPermissions) Docs() bool {
	return c&CPermDocs != 0
}
func (c CommunityPermissions) Groups() bool {
	return false
}
func (c CommunityPermissions) Notifications() bool {
	return false
}
func (c CommunityPermissions) Stats() bool {
	return false
}
func (c CommunityPermissions) Email() bool {
	return false
}
func (c CommunityPermissions) Market() bool {
	return false
}
func (c CommunityPermissions) AppWidget() bool {
	return c&CPermAppWidget != 0
}
func (c CommunityPermissions) Manage() bool {
	return c&CPermManage != 0
}
