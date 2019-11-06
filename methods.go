package vk

import (
	"reflect"

	"github.com/xelaj/vk/methods"
	"github.com/xelaj/vk/types"
)

func (c *Client) AccountBan(ownerId int) error {
	return methods.AccountBan(c.client(), ownerId)
}

/*
func (c *Client) AccountGetActiveOffers(param) (response) {
	return methods.AccountGetActiveOffers(c.client(), param)
}

*/

func (c *Client) AccountGetAppPermissions(userId int) (types.Permissions, error) {
	return methods.AccountGetAppPermissions(c.client(), userId)
}

func (c *Client) AccountGetBanned(offset, count int) (*methods.AccountGetBannedResponse, error) {
	return methods.AccountGetBanned(c.client(), offset, count)
}

/*

func (c *Client) AccountGetCounters(param) (response) {
	return methods.AccountGetCounters(c.client(), param)
}

func (c *Client) AccountGetInfo(param) (response) {
	return methods.AccountGetInfo(c.client(), param)
}

func (c *Client) AccountGetProfileInfo(param) (response) {
	return methods.AccountGetProfileInfo(c.client(), param)
}

func (c *Client) AccountGetPushSettings(param) (response) {
	return methods.AccountGetPushSettings(c.client(), param)
}

func (c *Client) AccountRegisterDevice(param) (response) {
	return methods.AccountRegisterDevice(c.client(), param)
}

func (c *Client) AccountSaveProfileInfo(param) (response) {
	return methods.AccountSaveProfileInfo(c.client(), param)
}

func (c *Client) AccountSetInfo(param) (response) {
	return methods.AccountSetInfo(c.client(), param)
}

func (c *Client) AccountGetProfileInfo(param) (response) {
	return methods.AccountGetProfileInfo(c.client(), param)
}

func (c *Client) AccountSetOffline(param) (response) {
	return methods.AccountSetOffline(c.client(), param)
}

func (c *Client) AccountSetOnline(param) (response) {
	return methods.AccountSetOnline(c.client(), param)
}

func (c *Client) AccountSetPushSettings(param) (response) {
	return methods.AccountSetPushSettings(c.client(), param)
}

func (c *Client) AccountSetSilenceMode(param) (response) {
	return methods.AccountSetSilenceMode(c.client(), param)
}

*/

func (c *Client) AccountUnban(ownerId int) error {
	return methods.AccountUnban(c.client(), ownerId)
}

/*

func (c *Client) AccountUnregisterDevice(param) (response) {
	return methods.AccountUnregisterDevice(c.client(), param)
}

func (c *Client) AdsAddOfficeUsers(param) (response) {
	return methods.AdsAddOfficeUsers(c.client(), param)
}

func (c *Client) AdsCheckLink(param) (response) {
	return methods.AdsCheckLink(c.client(), param)
}

func (c *Client) AdsCreateAds(param) (response) {
	return methods.AdsCreateAds(c.client(), param)
}

func (c *Client) AdsCreateCampaigns(param) (response) {
	return methods.AdsCreateCampaigns(c.client(), param)
}

func (c *Client) AdsCreateClients(param) (response) {
	return methods.AdsCreateClients(c.client(), param)
}

func (c *Client) AdsCreateLookalikeRequest(param) (response) {
	return methods.AdsCreateLookalikeRequest(c.client(), param)
}

func (c *Client) AdsCreateTargetGroup(param) (response) {
	return methods.AdsCreateTargetGroup(c.client(), param)
}

func (c *Client) AdsCreateTargetPixel(param) (response) {
	return methods.AdsCreateTargetPixel(c.client(), param)
}

func (c *Client) AdsDeleteAds(param) (response) {
	return methods.AdsDeleteAds(c.client(), param)
}

func (c *Client) AdsDeleteCampaigns(param) (response) {
	return methods.AdsDeleteCampaigns(c.client(), param)
}

func (c *Client) AdsDeleteClients(param) (response) {
	return methods.AdsDeleteClients(c.client(), param)
}

func (c *Client) AdsDeleteTargetGroup(param) (response) {
	return methods.AdsDeleteTargetGroup(c.client(), param)
}

func (c *Client) AdsDeleteTargetPixel(param) (response) {
	return methods.AdsDeleteTargetPixel(c.client(), param)
}

func (c *Client) AdsGetAccounts(param) (response) {
	return methods.AdsGetAccounts(c.client(), param)
}

func (c *Client) AdsGetAds(param) (response) {
	return methods.AdsGetAds(c.client(), param)
}

func (c *Client) AdsGetAdsLayout(param) (response) {
	return methods.AdsGetAdsLayout(c.client(), param)
}

func (c *Client) AdsGetAdsTargeting(param) (response) {
	return methods.AdsGetAdsTargeting(c.client(), param)
}

func (c *Client) AdsGetBudget(param) (response) {
	return methods.AdsGetBudget(c.client(), param)
}

func (c *Client) AdsGetCampaigns(param) (response) {
	return methods.AdsGetCampaigns(c.client(), param)
}

func (c *Client) AdsGetCategories(param) (response) {
	return methods.AdsGetCategories(c.client(), param)
}

func (c *Client) AdsGetClients(param) (response) {
	return methods.AdsGetClients(c.client(), param)
}

func (c *Client) AdsGetDemographics(param) (response) {
	return methods.AdsGetDemographics(c.client(), param)
}

func (c *Client) AdsGetFloodStats(param) (response) {
	return methods.AdsGetFloodStats(c.client(), param)
}

func (c *Client) AdsGetLookalikeRequests(param) (response) {
	return methods.AdsGetLookalikeRequests(c.client(), param)
}

func (c *Client) AdsGetOfficeUsers(param) (response) {
	return methods.AdsGetOfficeUsers(c.client(), param)
}

func (c *Client) AdsGetPostsReach(param) (response) {
	return methods.AdsGetPostsReach(c.client(), param)
}

func (c *Client) AdsGetRejectionReason(param) (response) {
	return methods.AdsGetRejectionReason(c.client(), param)
}

func (c *Client) AdsGetStatistics(param) (response) {
	return methods.AdsGetStatistics(c.client(), param)
}

func (c *Client) AdsGetSuggestions(param) (response) {
	return methods.AdsGetSuggestions(c.client(), param)
}

func (c *Client) AdsGetTargetGroups(param) (response) {
	return methods.AdsGetTargetGroups(c.client(), param)
}

func (c *Client) AdsGetTargetPixels(param) (response) {
	return methods.AdsGetTargetPixels(c.client(), param)
}

func (c *Client) AdsGetTargetingStats(param) (response) {
	return methods.AdsGetTargetingStats(c.client(), param)
}

func (c *Client) AdsGetUploadURL(param) (response) {
	return methods.AdsGetUploadURL(c.client(), param)
}

func (c *Client) AdsGetVideoUploadURL(param) (response) {
	return methods.AdsGetVideoUploadURL(c.client(), param)
}

func (c *Client) AdsImportTargetContacts(param) (response) {
	return methods.AdsImportTargetContacts(c.client(), param)
}

func (c *Client) AdsRemoveOfficeUsers(param) (response) {
	return methods.AdsRemoveOfficeUsers(c.client(), param)
}

func (c *Client) AdsRemoveTargetContacts(param) (response) {
	return methods.AdsRemoveTargetContacts(c.client(), param)
}

func (c *Client) AdsSaveLookalikeRequestResult(param) (response) {
	return methods.AdsSaveLookalikeRequestResult(c.client(), param)
}

func (c *Client) AdsShareTargetGroup(param) (response) {
	return methods.AdsShareTargetGroup(c.client(), param)
}

func (c *Client) AdsUpdateAds(param) (response) {
	return methods.AdsUpdateAds(c.client(), param)
}

func (c *Client) AdsUpdateCampaigns(param) (response) {
	return methods.AdsUpdateCampaigns(c.client(), param)
}

func (c *Client) AdsUpdateClients(param) (response) {
	return methods.AdsUpdateClients(c.client(), param)
}

func (c *Client) AdsUpdateTargetGroup(param) (response) {
	return methods.AdsUpdateTargetGroup(c.client(), param)
}

func (c *Client) AdsUpdateTargetPixel(param) (response) {
	return methods.AdsUpdateTargetPixel(c.client(), param)
}

func (c *Client) AppwdigetsGetAppImageUploadServer(param) (response) {
	return methods.AppwdigetsGetAppImageUploadServer(c.client(), param)
}

func (c *Client) AppwdigetsGetAppImages(param) (response) {
	return methods.AppwdigetsGetAppImages(c.client(), param)
}

func (c *Client) AppwdigetsGetGroupImageUploadServer(param) (response) {
	return methods.AppwdigetsGetGroupImageUploadServer(c.client(), param)
}

func (c *Client) AppwdigetsGetGroupImages(param) (response) {
	return methods.AppwdigetsGetGroupImages(c.client(), param)
}

func (c *Client) AppwdigetsGetImagesById(param) (response) {
	return methods.AppwdigetsGetImagesById(c.client(), param)
}

func (c *Client) AppwdigetsSaveAppImage(param) (response) {
	return methods.AppwdigetsSaveAppImage(c.client(), param)
}

func (c *Client) AppwdigetsSaveGroupImage(param) (response) {
	return methods.AppwdigetsSaveGroupImage(c.client(), param)
}

func (c *Client) AppwdigetsUpdate(param) (response) {
	return methods.AppwdigetsUpdate(c.client(), param)
}

func (c *Client) AppsDeleteAppRequests(param) (response) {
	return methods.AppsDeleteAppRequests(c.client(), param)
}

*/

func (c *Client) AppsGet(additional map[string]interface{}) (*methods.AppsGetResponse, error) {
	return methods.AppsGet(c.client(), additional)
}

/*

func (c *Client) AppsGetCatalog(param) (response) {
	return methods.AppsGetCatalog(c.client(), param)
}

func (c *Client) AppsGetFriendsList(param) (response) {
	return methods.AppsGetFriendsList(c.client(), param)
}

func (c *Client) AppsGetLeaderboard(param) (response) {
	return methods.AppsGetLeaderboard(c.client(), param)
}

func (c *Client) AppsGetScopes(param) (response) {
	return methods.AppsGetScopes(c.client(), param)
}

func (c *Client) AppsGetScore(param) (response) {
	return methods.AppsGetScore(c.client(), param)
}

func (c *Client) AppsPromoHasActiveGift(param) (response) {
	return methods.AppsPromoHasActiveGift(c.client(), param)
}

func (c *Client) AppsPromoUseGift(param) (response) {
	return methods.AppsPromoUseGift(c.client(), param)
}

func (c *Client) AppsSendRequest(param) (response) {
	return methods.AppsSendRequest(c.client(), param)
}

func (c *Client) AudioGet(param) (response) {
	return methods.AudioGet(c.client(), param)
}

func (c *Client) AudioGetById(param) (response) {
	return methods.AudioGetById(c.client(), param)
}

func (c *Client) AudioGetLyrics(param) (response) {
	return methods.AudioGetLyrics(c.client(), param)
}

func (c *Client) AudioSearch(param) (response) {
	return methods.AudioSearch(c.client(), param)
}

func (c *Client) AudioGetUploadServer(param) (response) {
	return methods.AudioGetUploadServer(c.client(), param)
}

func (c *Client) AudioSave(param) (response) {
	return methods.AudioSave(c.client(), param)
}

func (c *Client) AudioAdd(param) (response) {
	return methods.AudioAdd(c.client(), param)
}

func (c *Client) AudioDelete(param) (response) {
	return methods.AudioDelete(c.client(), param)
}

func (c *Client) AudioEdit(param) (response) {
	return methods.AudioEdit(c.client(), param)
}

func (c *Client) AudioReorder(param) (response) {
	return methods.AudioReorder(c.client(), param)
}

func (c *Client) AudioRestore(param) (response) {
	return methods.AudioRestore(c.client(), param)
}

func (c *Client) AudioGetAlbums(param) (response) {
	return methods.AudioGetAlbums(c.client(), param)
}

func (c *Client) AudioAddAlbum(param) (response) {
	return methods.AudioAddAlbum(c.client(), param)
}

func (c *Client) AudioEditAlbum(param) (response) {
	return methods.AudioEditAlbum(c.client(), param)
}

func (c *Client) AudioDeleteAlbum(param) (response) {
	return methods.AudioDeleteAlbum(c.client(), param)
}

func (c *Client) AudioMoveToAlbum(param) (response) {
	return methods.AudioMoveToAlbum(c.client(), param)
}

func (c *Client) AudioSetBroadcast(param) (response) {
	return methods.AudioSetBroadcast(c.client(), param)
}

func (c *Client) AudioGetBroadcastList(param) (response) {
	return methods.AudioGetBroadcastList(c.client(), param)
}

func (c *Client) AudioGetRecommendations(param) (response) {
	return methods.AudioGetRecommendations(c.client(), param)
}

func (c *Client) AudioGetPopular(param) (response) {
	return methods.AudioGetPopular(c.client(), param)
}

func (c *Client) AudioGetCount(param) (response) {
	return methods.AudioGetCount(c.client(), param)
}

func (c *Client) AuthCheckPhone(param) (response) {
	return methods.AuthCheckPhone(c.client(), param)
}

func (c *Client) AuthRestore(param) (response) {
	return methods.AuthRestore(c.client(), param)
}

func (c *Client) BoardAddTopic(param) (response) {
	return methods.BoardAddTopic(c.client(), param)
}

func (c *Client) BoardCloseTopic(param) (response) {
	return methods.BoardCloseTopic(c.client(), param)
}

func (c *Client) BoardCreateComment(param) (response) {
	return methods.BoardCreateComment(c.client(), param)
}

func (c *Client) BoardDeleteComment(param) (response) {
	return methods.BoardDeleteComment(c.client(), param)
}

func (c *Client) BoardDeleteTopic(param) (response) {
	return methods.BoardDeleteTopic(c.client(), param)
}

func (c *Client) BoardEditComment(param) (response) {
	return methods.BoardEditComment(c.client(), param)
}

func (c *Client) BoardEditTopic(param) (response) {
	return methods.BoardEditTopic(c.client(), param)
}

func (c *Client) BoardFixTopic(param) (response) {
	return methods.BoardFixTopic(c.client(), param)
}

func (c *Client) BoardGetComments(param) (response) {
	return methods.BoardGetComments(c.client(), param)
}

func (c *Client) BoardGetTopics(param) (response) {
	return methods.BoardGetTopics(c.client(), param)
}

func (c *Client) BoardOpenTopic(param) (response) {
	return methods.BoardOpenTopic(c.client(), param)
}

func (c *Client) BoardRestoreComment(param) (response) {
	return methods.BoardRestoreComment(c.client(), param)
}

func (c *Client) BoardUnfixTopic(param) (response) {
	return methods.BoardUnfixTopic(c.client(), param)
}

func (c *Client) DatabaseGetChairs(param) (response) {
	return methods.DatabaseGetChairs(c.client(), param)
}

func (c *Client) DatabaseGetCities(param) (response) {
	return methods.DatabaseGetCities(c.client(), param)
}

func (c *Client) DatabaseGetCitiesById(param) (response) {
	return methods.DatabaseGetCitiesById(c.client(), param)
}

func (c *Client) DatabaseGetCountries(param) (response) {
	return methods.DatabaseGetCountries(c.client(), param)
}

func (c *Client) DatabaseGetCountriesById(param) (response) {
	return methods.DatabaseGetCountriesById(c.client(), param)
}

func (c *Client) DatabaseGetFaculties(param) (response) {
	return methods.DatabaseGetFaculties(c.client(), param)
}

func (c *Client) DatabaseGetMetroStations(param) (response) {
	return methods.DatabaseGetMetroStations(c.client(), param)
}

func (c *Client) DatabaseGetMetroStationsById(param) (response) {
	return methods.DatabaseGetMetroStationsById(c.client(), param)
}

func (c *Client) DatabaseGetRegions(param) (response) {
	return methods.DatabaseGetRegions(c.client(), param)
}

func (c *Client) DatabaseGetSchoolClasses(param) (response) {
	return methods.DatabaseGetSchoolClasses(c.client(), param)
}

func (c *Client) DatabaseGetSchools(param) (response) {
	return methods.DatabaseGetSchools(c.client(), param)
}

func (c *Client) DatabaseGetUniversities(param) (response) {
	return methods.DatabaseGetUniversities(c.client(), param)
}

func (c *Client) DocsAdd(param) (response) {
	return methods.DocsAdd(c.client(), param)
}

func (c *Client) DocsDelete(param) (response) {
	return methods.DocsDelete(c.client(), param)
}

func (c *Client) DocsEdit(param) (response) {
	return methods.DocsEdit(c.client(), param)
}

func (c *Client) DocsGet(param) (response) {
	return methods.DocsGet(c.client(), param)
}

func (c *Client) DocsGetById(param) (response) {
	return methods.DocsGetById(c.client(), param)
}

func (c *Client) DocsGetMessagesUploadServer(param) (response) {
	return methods.DocsGetMessagesUploadServer(c.client(), param)
}

func (c *Client) DocsGetTypes(param) (response) {
	return methods.DocsGetTypes(c.client(), param)
}

func (c *Client) DocsGetUploadServer(param) (response) {
	return methods.DocsGetUploadServer(c.client(), param)
}

func (c *Client) DocsGetWallUploadServer(param) (response) {
	return methods.DocsGetWallUploadServer(c.client(), param)
}

func (c *Client) DocsSave(param) (response) {
	return methods.DocsSave(c.client(), param)
}

func (c *Client) DocsSearch(param) (response) {
	return methods.DocsSearch(c.client(), param)
}

*/

func (c *Client) Execucte(method string, params map[string]interface{}, resType reflect.Type) (interface{}, error) {
	return methods.ExecuteMethod(c.client(), method, params, resType)
}

/*

func (c *Client) FaveAddArticle(param) (response) {
	return methods.FaveAddArticle(c.client(), param)
}

func (c *Client) FaveAddLink(param) (response) {
	return methods.FaveAddLink(c.client(), param)
}

func (c *Client) FaveAddPage(param) (response) {
	return methods.FaveAddPage(c.client(), param)
}

func (c *Client) FaveAddPost(param) (response) {
	return methods.FaveAddPost(c.client(), param)
}

func (c *Client) FaveAddProduct(param) (response) {
	return methods.FaveAddProduct(c.client(), param)
}

func (c *Client) FaveAddTag(param) (response) {
	return methods.FaveAddTag(c.client(), param)
}

func (c *Client) FaveAddVideo(param) (response) {
	return methods.FaveAddVideo(c.client(), param)
}

func (c *Client) FaveEditTag(param) (response) {
	return methods.FaveEditTag(c.client(), param)
}

func (c *Client) FaveGet(param) (response) {
	return methods.FaveGet(c.client(), param)
}

func (c *Client) FaveGetPages(param) (response) {
	return methods.FaveGetPages(c.client(), param)
}

func (c *Client) FaveGetTags(param) (response) {
	return methods.FaveGetTags(c.client(), param)
}

func (c *Client) FaveMarkSeen(param) (response) {
	return methods.FaveMarkSeen(c.client(), param)
}

func (c *Client) FaveRemoveArticle(param) (response) {
	return methods.FaveRemoveArticle(c.client(), param)
}

func (c *Client) FaveRemoveLink(param) (response) {
	return methods.FaveRemoveLink(c.client(), param)
}

func (c *Client) FaveRemovePage(param) (response) {
	return methods.FaveRemovePage(c.client(), param)
}

func (c *Client) FaveRemovePost(param) (response) {
	return methods.FaveRemovePost(c.client(), param)
}

func (c *Client) FaveRemoveProduct(param) (response) {
	return methods.FaveRemoveProduct(c.client(), param)
}

func (c *Client) FaveRemoveTag(param) (response) {
	return methods.FaveRemoveTag(c.client(), param)
}

func (c *Client) FaveRemoveVideo(param) (response) {
	return methods.FaveRemoveVideo(c.client(), param)
}

func (c *Client) FaveReorderTags(param) (response) {
	return methods.FaveReorderTags(c.client(), param)
}

func (c *Client) FaveSetPageTags(param) (response) {
	return methods.FaveSetPageTags(c.client(), param)
}

func (c *Client) FaveSetTags(param) (response) {
	return methods.FaveSetTags(c.client(), param)
}

func (c *Client) FaveTrackPageInteraction(param) (response) {
	return methods.FaveTrackPageInteraction(c.client(), param)
}

func (c *Client) FriendsAdd(param) (response) {
	return methods.FriendsAdd(c.client(), param)
}

func (c *Client) FriendsAddList(param) (response) {
	return methods.FriendsAddList(c.client(), param)
}

func (c *Client) FriendsAreFriends(param) (response) {
	return methods.FriendsAreFriends(c.client(), param)
}

func (c *Client) FriendsDelete(param) (response) {
	return methods.FriendsDelete(c.client(), param)
}

func (c *Client) FriendsDeleteAllRequests(param) (response) {
	return methods.FriendsDeleteAllRequests(c.client(), param)
}

func (c *Client) FriendsDeleteList(param) (response) {
	return methods.FriendsDeleteList(c.client(), param)
}

func (c *Client) FriendsEdit(param) (response) {
	return methods.FriendsEdit(c.client(), param)
}

func (c *Client) FriendsEditList(param) (response) {
	return methods.FriendsEditList(c.client(), param)
}

*/

func (c *Client) FriendsGet(userId int, Order string, fields []string, nameCase string) (*methods.FriendsGetResponse, error) {
	return methods.FriendsGet(c.client(), userId, Order, fields, nameCase, MAX_COUNT_FRIENDS)
}

func (c *Client) FriendsGetId(userId int, Order string, nameCase string) (*methods.FriendsGetIdResponse, error) {
	return methods.FriendsGetId(c.client(), userId, Order, nameCase, MAX_COUNT_FRIENDS)
}

/*

func (c *Client) FriendsGetAppUsers(param) (response) {
	return methods.FriendsGetAppUsers(c.client(), param)
}

func (c *Client) FriendsGetByPhones(param) (response) {
	return methods.FriendsGetByPhones(c.client(), param)
}

func (c *Client) FriendsGetLists(param) (response) {
	return methods.FriendsGetLists(c.client(), param)
}

func (c *Client) FriendsGetMutual(param) (response) {
	return methods.FriendsGetMutual(c.client(), param)
}

func (c *Client) FriendsGetOnline(param) (response) {
	return methods.FriendsGetOnline(c.client(), param)
}

func (c *Client) FriendsGetRecent(param) (response) {
	return methods.FriendsGetRecent(c.client(), param)
}

func (c *Client) FriendsGetRequests(param) (response) {
	return methods.FriendsGetRequests(c.client(), param)
}

func (c *Client) FriendsGetSuggestions(param) (response) {
	return methods.FriendsGetSuggestions(c.client(), param)
}

func (c *Client) FriendsSearch(param) (response) {
	return methods.FriendsSearch(c.client(), param)
}

func (c *Client) GiftsGet(param) (response) {
	return methods.GiftsGet(c.client(), param)
}

func (c *Client) GroupsAddAddress(param) (response) {
	return methods.GroupsAddAddress(c.client(), param)
}

func (c *Client) GroupsAddCallbackServer(param) (response) {
	return methods.GroupsAddCallbackServer(c.client(), param)
}

func (c *Client) GroupsAddLink(param) (response) {
	return methods.GroupsAddLink(c.client(), param)
}

func (c *Client) GroupsApproveRequest(param) (response) {
	return methods.GroupsApproveRequest(c.client(), param)
}

func (c *Client) GroupsBan(param) (response) {
	return methods.GroupsBan(c.client(), param)
}

func (c *Client) GroupsCreate(param) (response) {
	return methods.GroupsCreate(c.client(), param)
}

func (c *Client) GroupsDeleteAddress(param) (response) {
	return methods.GroupsDeleteAddress(c.client(), param)
}

func (c *Client) GroupsDeleteCallbackServer(param) (response) {
	return methods.GroupsDeleteCallbackServer(c.client(), param)
}

func (c *Client) GroupsDeleteLink(param) (response) {
	return methods.GroupsDeleteLink(c.client(), param)
}

func (c *Client) GroupsDisableOnline(param) (response) {
	return methods.GroupsDisableOnline(c.client(), param)
}

func (c *Client) GroupsEdit(param) (response) {
	return methods.GroupsEdit(c.client(), param)
}

func (c *Client) GroupsEditAddress(param) (response) {
	return methods.GroupsEditAddress(c.client(), param)
}

func (c *Client) GroupsEditCallbackServer(param) (response) {
	return methods.GroupsEditCallbackServer(c.client(), param)
}

func (c *Client) GroupsEditLink(param) (response) {
	return methods.GroupsEditLink(c.client(), param)
}

func (c *Client) GroupsEditManager(param) (response) {
	return methods.GroupsEditManager(c.client(), param)
}

func (c *Client) GroupsEnableOnline(param) (response) {
	return methods.GroupsEnableOnline(c.client(), param)
}

*/

func (c *Client) GroupsGet(userId int, filter, fields []string) (*methods.GroupsGetResponse, error) {
	return methods.GroupsGet(c.client(), userId, filter, fields, MAX_COUNT_GROUPS)
}

func (c *Client) GroupsGetById(groupIds []int, fields []string) (*methods.GroupsGetByIdResponse, error) {
	return methods.GroupsGetById(c.client(), groupIds, fields, MAX_COUNT_GROUPS)
}

/*

func (c *Client) GroupsGetAddresses(param) (response) {
	return methods.GroupsGetAddresses(c.client(), param)
}

func (c *Client) GroupsGetBanned(param) (response) {
	return methods.GroupsGetBanned(c.client(), param)
}

func (c *Client) GroupsGetById(param) (response) {
	return methods.GroupsGetById(c.client(), param)
}

func (c *Client) GroupsGetCallbackConfirmationCode(param) (response) {
	return methods.GroupsGetCallbackConfirmationCode(c.client(), param)
}

func (c *Client) GroupsGetCallbackServers(param) (response) {
	return methods.GroupsGetCallbackServers(c.client(), param)
}

func (c *Client) GroupsGetCallbackSettings(param) (response) {
	return methods.GroupsGetCallbackSettings(c.client(), param)
}

func (c *Client) GroupsGetCatalog(param) (response) {
	return methods.GroupsGetCatalog(c.client(), param)
}

func (c *Client) GroupsGetCatalogInfo(param) (response) {
	return methods.GroupsGetCatalogInfo(c.client(), param)
}

func (c *Client) GroupsGetInvitedUsers(param) (response) {
	return methods.GroupsGetInvitedUsers(c.client(), param)
}

func (c *Client) GroupsGetInvites(param) (response) {
	return methods.GroupsGetInvites(c.client(), param)
}

*/

func (c *Client) GroupsGetLongPollServer(groupId int) (*methods.GroupsGetLongPollServerResponse, error) {
	return methods.GroupsGetLongPollServer(c.client(), groupId)
}

/*

func (c *Client) GroupsGetLongPollSettings(param) (response) {
	return methods.GroupsGetLongPollSettings(c.client(), param)
}

func (c *Client) GroupsGetMembers(param) (response) {
	return methods.GroupsGetMembers(c.client(), param)
}

func (c *Client) GroupsGetOnlineStatus(param) (response) {
	return methods.GroupsGetOnlineStatus(c.client(), param)
}

func (c *Client) GroupsGetRequests(param) (response) {
	return methods.GroupsGetRequests(c.client(), param)
}

func (c *Client) GroupsGetSettings(param) (response) {
	return methods.GroupsGetSettings(c.client(), param)
}

*/

func (c *Client) GroupsGetTokenPermissions() (types.CommunityPermissions, error) {
	return methods.GroupsGetTokenPermissions(c.client())
}

/*

func (c *Client) GroupsInvite(param) (response) {
	return methods.GroupsInvite(c.client(), param)
}

func (c *Client) GroupsIsMember(param) (response) {
	return methods.GroupsIsMember(c.client(), param)
}

func (c *Client) GroupsJoin(param) (response) {
	return methods.GroupsJoin(c.client(), param)
}

func (c *Client) GroupsLeave(param) (response) {
	return methods.GroupsLeave(c.client(), param)
}

func (c *Client) GroupsRemoveUser(param) (response) {
	return methods.GroupsRemoveUser(c.client(), param)
}

func (c *Client) GroupsReorderLink(param) (response) {
	return methods.GroupsReorderLink(c.client(), param)
}

func (c *Client) GroupsSearch(param) (response) {
	return methods.GroupsSearch(c.client(), param)
}

func (c *Client) GroupsSetCallbackSettings(param) (response) {
	return methods.GroupsSetCallbackSettings(c.client(), param)
}

func (c *Client) GroupsSetLongPollSettings(param) (response) {
	return methods.GroupsSetLongPollSettings(c.client(), param)
}

func (c *Client) GroupsSetSettings(param) (response) {
	return methods.GroupsSetSettings(c.client(), param)
}

func (c *Client) GroupsUnban(param) (response) {
	return methods.GroupsUnban(c.client(), param)
}

func (c *Client) LeadformsCreate(param) (response) {
	return methods.LeadformsCreate(c.client(), param)
}

func (c *Client) LeadformsDelete(param) (response) {
	return methods.LeadformsDelete(c.client(), param)
}

func (c *Client) LeadformsGet(param) (response) {
	return methods.LeadformsGet(c.client(), param)
}

func (c *Client) LeadformsGetLeads(param) (response) {
	return methods.LeadformsGetLeads(c.client(), param)
}

func (c *Client) LeadformsGetUploadURL(param) (response) {
	return methods.LeadformsGetUploadURL(c.client(), param)
}

func (c *Client) LeadformsList(param) (response) {
	return methods.LeadformsList(c.client(), param)
}

func (c *Client) LeadformsUpdate(param) (response) {
	return methods.LeadformsUpdate(c.client(), param)
}

func (c *Client) LeadsCheckUser(param) (response) {
	return methods.LeadsCheckUser(c.client(), param)
}

func (c *Client) LeadsComplete(param) (response) {
	return methods.LeadsComplete(c.client(), param)
}

func (c *Client) LeadsGetStats(param) (response) {
	return methods.LeadsGetStats(c.client(), param)
}

func (c *Client) LeadsGetUsers(param) (response) {
	return methods.LeadsGetUsers(c.client(), param)
}

func (c *Client) LeadsMetricHit(param) (response) {
	return methods.LeadsMetricHit(c.client(), param)
}

func (c *Client) LeadsStart(param) (response) {
	return methods.LeadsStart(c.client(), param)
}

func (c *Client) LikesAdd(param) (response) {
	return methods.LikesAdd(c.client(), param)
}

func (c *Client) LikesDelete(param) (response) {
	return methods.LikesDelete(c.client(), param)
}

func (c *Client) LikesGetList(param) (response) {
	return methods.LikesGetList(c.client(), param)
}

func (c *Client) LikesIsLiked(param) (response) {
	return methods.LikesIsLiked(c.client(), param)
}

func (c *Client) MarketAdd(param) (response) {
	return methods.MarketAdd(c.client(), param)
}

func (c *Client) MarketAddAlbum(param) (response) {
	return methods.MarketAddAlbum(c.client(), param)
}

func (c *Client) MarketAddToAlbum(param) (response) {
	return methods.MarketAddToAlbum(c.client(), param)
}

func (c *Client) MarketCreateComment(param) (response) {
	return methods.MarketCreateComment(c.client(), param)
}

func (c *Client) MarketDelete(param) (response) {
	return methods.MarketDelete(c.client(), param)
}

func (c *Client) MarketDeleteAlbum(param) (response) {
	return methods.MarketDeleteAlbum(c.client(), param)
}

func (c *Client) MarketDeleteComment(param) (response) {
	return methods.MarketDeleteComment(c.client(), param)
}

func (c *Client) MarketEdit(param) (response) {
	return methods.MarketEdit(c.client(), param)
}

func (c *Client) MarketEditAlbum(param) (response) {
	return methods.MarketEditAlbum(c.client(), param)
}

func (c *Client) MarketEditComment(param) (response) {
	return methods.MarketEditComment(c.client(), param)
}

func (c *Client) MarketGet(param) (response) {
	return methods.MarketGet(c.client(), param)
}

func (c *Client) MarketGetAlbumById(param) (response) {
	return methods.MarketGetAlbumById(c.client(), param)
}

func (c *Client) MarketGetAlbums(param) (response) {
	return methods.MarketGetAlbums(c.client(), param)
}

func (c *Client) MarketGetById(param) (response) {
	return methods.MarketGetById(c.client(), param)
}

func (c *Client) MarketGetCategories(param) (response) {
	return methods.MarketGetCategories(c.client(), param)
}

func (c *Client) MarketGetComments(param) (response) {
	return methods.MarketGetComments(c.client(), param)
}

func (c *Client) MarketRemoveFromAlbum(param) (response) {
	return methods.MarketRemoveFromAlbum(c.client(), param)
}

func (c *Client) MarketReorderAlbums(param) (response) {
	return methods.MarketReorderAlbums(c.client(), param)
}

func (c *Client) MarketReorderItems(param) (response) {
	return methods.MarketReorderItems(c.client(), param)
}

func (c *Client) MarketReport(param) (response) {
	return methods.MarketReport(c.client(), param)
}

func (c *Client) MarketReportComment(param) (response) {
	return methods.MarketReportComment(c.client(), param)
}

func (c *Client) MarketRestore(param) (response) {
	return methods.MarketRestore(c.client(), param)
}

func (c *Client) MarketRestoreComment(param) (response) {
	return methods.MarketRestoreComment(c.client(), param)
}

func (c *Client) MarketSearch(param) (response) {
	return methods.MarketSearch(c.client(), param)
}

func (c *Client) MessagesAddChatUser(param) (response) {
	return methods.MessagesAddChatUser(c.client(), param)
}

func (c *Client) MessagesAllowMessagesFromGroup(param) (response) {
	return methods.MessagesAllowMessagesFromGroup(c.client(), param)
}

func (c *Client) MessagesCreateChat(param) (response) {
	return methods.MessagesCreateChat(c.client(), param)
}

func (c *Client) MessagesDelete(param) (response) {
	return methods.MessagesDelete(c.client(), param)
}

func (c *Client) MessagesDeleteChatPhoto(param) (response) {
	return methods.MessagesDeleteChatPhoto(c.client(), param)
}

func (c *Client) MessagesDeleteConversation(param) (response) {
	return methods.MessagesDeleteConversation(c.client(), param)
}

func (c *Client) MessagesDenyMessagesFromGroup(param) (response) {
	return methods.MessagesDenyMessagesFromGroup(c.client(), param)
}

func (c *Client) MessagesEdit(param) (response) {
	return methods.MessagesEdit(c.client(), param)
}

func (c *Client) MessagesEditChat(param) (response) {
	return methods.MessagesEditChat(c.client(), param)
}

func (c *Client) MessagesGetByConversationMessageId(param) (response) {
	return methods.MessagesGetByConversationMessageId(c.client(), param)
}

*/

func (c *Client) MessagesGetById(messageIds []int, previewLen, groupId int) (*methods.MessagesGetByIdResponse, error) {
	return methods.MessagesGetById(c, messageIds, previewLen, groupId)
}

/*

func (c *Client) MessagesGetChat(param) (response) {
	return methods.MessagesGetChat(c.client(), param)
}

func (c *Client) MessagesGetChatPreview(param) (response) {
	return methods.MessagesGetChatPreview(c.client(), param)
}

func (c *Client) MessagesGetConversationMembers(param) (response) {
	return methods.MessagesGetConversationMembers(c.client(), param)
}

func (c *Client) MessagesGetConversations(param) (response) {
	return methods.MessagesGetConversations(c.client(), param)
}

func (c *Client) MessagesGetConversationsById(param) (response) {
	return methods.MessagesGetConversationsById(c.client(), param)
}

func (c *Client) MessagesGetHistory(param) (response) {
	return methods.MessagesGetHistory(c.client(), param)
}

func (c *Client) MessagesGetHistoryAttachments(param) (response) {
	return methods.MessagesGetHistoryAttachments(c.client(), param)
}

func (c *Client) MessagesGetImportantMessages(param) (response) {
	return methods.MessagesGetImportantMessages(c.client(), param)
}

func (c *Client) MessagesGetInviteLink(param) (response) {
	return methods.MessagesGetInviteLink(c.client(), param)
}

func (c *Client) MessagesGetLastActivity(param) (response) {
	return methods.MessagesGetLastActivity(c.client(), param)
}

func (c *Client) MessagesGetLongPollHistory(param) (response) {
	return methods.MessagesGetLongPollHistory(c.client(), param)
}

func (c *Client) MessagesGetLongPollServer(param) (response) {
	return methods.MessagesGetLongPollServer(c.client(), param)
}

func (c *Client) MessagesIsMessagesFromGroupAllowed(param) (response) {
	return methods.MessagesIsMessagesFromGroupAllowed(c.client(), param)
}

func (c *Client) MessagesJoinChatByInviteLink(param) (response) {
	return methods.MessagesJoinChatByInviteLink(c.client(), param)
}

func (c *Client) MessagesMarkAsAnsweredConversation(param) (response) {
	return methods.MessagesMarkAsAnsweredConversation(c.client(), param)
}

func (c *Client) MessagesMarkAsImportant(param) (response) {
	return methods.MessagesMarkAsImportant(c.client(), param)
}

func (c *Client) MessagesMarkAsImportantConversation(param) (response) {
	return methods.MessagesMarkAsImportantConversation(c.client(), param)
}

func (c *Client) MessagesMarkAsRead(param) (response) {
	return methods.MessagesMarkAsRead(c.client(), param)
}

func (c *Client) MessagesPin(param) (response) {
	return methods.MessagesPin(c.client(), param)
}

func (c *Client) MessagesRemoveChatUser(param) (response) {
	return methods.MessagesRemoveChatUser(c.client(), param)
}

func (c *Client) MessagesRestore(param) (response) {
	return methods.MessagesRestore(c.client(), param)
}

func (c *Client) MessagesSearch(param) (response) {
	return methods.MessagesSearch(c.client(), param)
}

func (c *Client) MessagesSearchConversations(param) (response) {
	return methods.MessagesSearchConversations(c.client(), param)
}

*/

func (c *Client) MessagesSend(m methods.MessagesSendRequest) (*methods.MessagesSendResponse, error) {
	return methods.MessagesSend(c.client(), m)
}

/*

func (c *Client) MessagesSetActivity(param) (response) {
	return methods.MessagesSetActivity(c.client(), param)
}

func (c *Client) MessagesSetChatPhoto(param) (response) {
	return methods.MessagesSetChatPhoto(c.client(), param)
}

func (c *Client) MessagesUnpin(param) (response) {
	return methods.MessagesUnpin(c.client(), param)
}

func (c *Client) NewsfeedAddBan(param) (response) {
	return methods.NewsfeedAddBan(c.client(), param)
}

func (c *Client) NewsfeedDeleteBan(param) (response) {
	return methods.NewsfeedDeleteBan(c.client(), param)
}

func (c *Client) NewsfeedDeleteList(param) (response) {
	return methods.NewsfeedDeleteList(c.client(), param)
}

func (c *Client) NewsfeedGet(param) (response) {
	return methods.NewsfeedGet(c.client(), param)
}

func (c *Client) NewsfeedGetBanned(param) (response) {
	return methods.NewsfeedGetBanned(c.client(), param)
}

func (c *Client) NewsfeedGetComments(param) (response) {
	return methods.NewsfeedGetComments(c.client(), param)
}

func (c *Client) NewsfeedGetLists(param) (response) {
	return methods.NewsfeedGetLists(c.client(), param)
}

func (c *Client) NewsfeedGetMentions(param) (response) {
	return methods.NewsfeedGetMentions(c.client(), param)
}

func (c *Client) NewsfeedGetRecommended(param) (response) {
	return methods.NewsfeedGetRecommended(c.client(), param)
}

func (c *Client) NewsfeedGetSuggestedSources(param) (response) {
	return methods.NewsfeedGetSuggestedSources(c.client(), param)
}

func (c *Client) NewsfeedIgnoreItem(param) (response) {
	return methods.NewsfeedIgnoreItem(c.client(), param)
}

func (c *Client) NewsfeedSaveList(param) (response) {
	return methods.NewsfeedSaveList(c.client(), param)
}

func (c *Client) NewsfeedSearch(param) (response) {
	return methods.NewsfeedSearch(c.client(), param)
}

func (c *Client) NewsfeedUnignoreItem(param) (response) {
	return methods.NewsfeedUnignoreItem(c.client(), param)
}

func (c *Client) NewsfeedUnsubscribe(param) (response) {
	return methods.NewsfeedUnsubscribe(c.client(), param)
}

func (c *Client) NotesAdd(param) (response) {
	return methods.NotesAdd(c.client(), param)
}

func (c *Client) NotesCreateComment(param) (response) {
	return methods.NotesCreateComment(c.client(), param)
}

func (c *Client) NotesDelete(param) (response) {
	return methods.NotesDelete(c.client(), param)
}

func (c *Client) NotesDeleteComment(param) (response) {
	return methods.NotesDeleteComment(c.client(), param)
}

func (c *Client) NotesEdit(param) (response) {
	return methods.NotesEdit(c.client(), param)
}

func (c *Client) NotesEditComment(param) (response) {
	return methods.NotesEditComment(c.client(), param)
}

func (c *Client) NotesGet(param) (response) {
	return methods.NotesGet(c.client(), param)
}

func (c *Client) NotesGetById(param) (response) {
	return methods.NotesGetById(c.client(), param)
}

func (c *Client) NotesGetComments(param) (response) {
	return methods.NotesGetComments(c.client(), param)
}

func (c *Client) NotesRestoreComment(param) (response) {
	return methods.NotesRestoreComment(c.client(), param)
}

func (c *Client) NotificationsGet(param) (response) {
	return methods.NotificationsGet(c.client(), param)
}

func (c *Client) NotificationsMarkAsViewed(param) (response) {
	return methods.NotificationsMarkAsViewed(c.client(), param)
}

func (c *Client) NotificationsSendMessage(param) (response) {
	return methods.NotificationsSendMessage(c.client(), param)
}

func (c *Client) OrdersCancelSubscription(param) (response) {
	return methods.OrdersCancelSubscription(c.client(), param)
}

func (c *Client) OrdersChangeState(param) (response) {
	return methods.OrdersChangeState(c.client(), param)
}

func (c *Client) OrdersGet(param) (response) {
	return methods.OrdersGet(c.client(), param)
}

func (c *Client) OrdersGetAmount(param) (response) {
	return methods.OrdersGetAmount(c.client(), param)
}

func (c *Client) OrdersGetById(param) (response) {
	return methods.OrdersGetById(c.client(), param)
}

func (c *Client) OrdersGetUserSubscriptionById(param) (response) {
	return methods.OrdersGetUserSubscriptionById(c.client(), param)
}

func (c *Client) OrdersGetUserSubscriptions(param) (response) {
	return methods.OrdersGetUserSubscriptions(c.client(), param)
}

func (c *Client) OrdersUpdateSubscription(param) (response) {
	return methods.OrdersUpdateSubscription(c.client(), param)
}

func (c *Client) PagesClearCache(param) (response) {
	return methods.PagesClearCache(c.client(), param)
}

func (c *Client) PagesGet(param) (response) {
	return methods.PagesGet(c.client(), param)
}

func (c *Client) PagesGetHistory(param) (response) {
	return methods.PagesGetHistory(c.client(), param)
}

func (c *Client) PagesGetTitles(param) (response) {
	return methods.PagesGetTitles(c.client(), param)
}

func (c *Client) PagesGetVersion(param) (response) {
	return methods.PagesGetVersion(c.client(), param)
}

func (c *Client) PagesParseWiki(param) (response) {
	return methods.PagesParseWiki(c.client(), param)
}

func (c *Client) PagesSave(param) (response) {
	return methods.PagesSave(c.client(), param)
}

func (c *Client) PagesSaveAccess(param) (response) {
	return methods.PagesSaveAccess(c.client(), param)
}

func (c *Client) PhotoConfirmTag(param) (response) {
	return methods.PhotoConfirmTag(c.client(), param)
}

func (c *Client) PhotoCopy(param) (response) {
	return methods.PhotoCopy(c.client(), param)
}

func (c *Client) PhotoCreateAlbum(param) (response) {
	return methods.PhotoCreateAlbum(c.client(), param)
}

func (c *Client) PhotoCreateComment(param) (response) {
	return methods.PhotoCreateComment(c.client(), param)
}

func (c *Client) PhotoDelete(param) (response) {
	return methods.PhotoDelete(c.client(), param)
}

func (c *Client) PhotoDeleteAlbum(param) (response) {
	return methods.PhotoDeleteAlbum(c.client(), param)
}

func (c *Client) PhotoDeleteComment(param) (response) {
	return methods.PhotoDeleteComment(c.client(), param)
}

func (c *Client) PhotoEdit(param) (response) {
	return methods.PhotoEdit(c.client(), param)
}

func (c *Client) PhotoEditAlbum(param) (response) {
	return methods.PhotoEditAlbum(c.client(), param)
}

func (c *Client) PhotoEditComment(param) (response) {
	return methods.PhotoEditComment(c.client(), param)
}

func (c *Client) PhotoGet(param) (response) {
	return methods.PhotoGet(c.client(), param)
}

func (c *Client) PhotoGetAlbums(param) (response) {
	return methods.PhotoGetAlbums(c.client(), param)
}

func (c *Client) PhotoGetAlbumsCount(param) (response) {
	return methods.PhotoGetAlbumsCount(c.client(), param)
}

func (c *Client) PhotoGetAll(param) (response) {
	return methods.PhotoGetAll(c.client(), param)
}

func (c *Client) PhotoGetAllComments(param) (response) {
	return methods.PhotoGetAllComments(c.client(), param)
}

func (c *Client) PhotoGetById(param) (response) {
	return methods.PhotoGetById(c.client(), param)
}

func (c *Client) PhotoGetChatUploadServer(param) (response) {
	return methods.PhotoGetChatUploadServer(c.client(), param)
}

func (c *Client) PhotoGetComments(param) (response) {
	return methods.PhotoGetComments(c.client(), param)
}

func (c *Client) PhotoGetMarketAlbumUploadServer(param) (response) {
	return methods.PhotoGetMarketAlbumUploadServer(c.client(), param)
}

func (c *Client) PhotoGetMarketUploadServer(param) (response) {
	return methods.PhotoGetMarketUploadServer(c.client(), param)
}

func (c *Client) PhotoGetMessagesUploadServer(param) (response) {
	return methods.PhotoGetMessagesUploadServer(c.client(), param)
}

func (c *Client) PhotoGetNewTags(param) (response) {
	return methods.PhotoGetNewTags(c.client(), param)
}

func (c *Client) PhotoGetOwnerCoverPhotoUploadServer(param) (response) {
	return methods.PhotoGetOwnerCoverPhotoUploadServer(c.client(), param)
}

func (c *Client) PhotoGetOwnerPhotoUploadServer(param) (response) {
	return methods.PhotoGetOwnerPhotoUploadServer(c.client(), param)
}

func (c *Client) PhotoGetTags(param) (response) {
	return methods.PhotoGetTags(c.client(), param)
}

func (c *Client) PhotoGetUploadServer(param) (response) {
	return methods.PhotoGetUploadServer(c.client(), param)
}

func (c *Client) PhotoGetUserPhotos(param) (response) {
	return methods.PhotoGetUserPhotos(c.client(), param)
}

func (c *Client) PhotoGetWallUploadServer(param) (response) {
	return methods.PhotoGetWallUploadServer(c.client(), param)
}

func (c *Client) PhotoMakeCover(param) (response) {
	return methods.PhotoMakeCover(c.client(), param)
}

func (c *Client) PhotoMove(param) (response) {
	return methods.PhotoMove(c.client(), param)
}

func (c *Client) PhotoPutTag(param) (response) {
	return methods.PhotoPutTag(c.client(), param)
}

func (c *Client) PhotoRemoveTag(param) (response) {
	return methods.PhotoRemoveTag(c.client(), param)
}

func (c *Client) PhotoReorderAlbums(param) (response) {
	return methods.PhotoReorderAlbums(c.client(), param)
}

func (c *Client) PhotoReorderPhotos(param) (response) {
	return methods.PhotoReorderPhotos(c.client(), param)
}

func (c *Client) PhotoReport(param) (response) {
	return methods.PhotoReport(c.client(), param)
}

func (c *Client) PhotoReportComment(param) (response) {
	return methods.PhotoReportComment(c.client(), param)
}

func (c *Client) PhotoRestore(param) (response) {
	return methods.PhotoRestore(c.client(), param)
}

func (c *Client) PhotoRestoreComment(param) (response) {
	return methods.PhotoRestoreComment(c.client(), param)
}

func (c *Client) PhotoSave(param) (response) {
	return methods.PhotoSave(c.client(), param)
}

func (c *Client) PhotoSaveMarketAlbumPhoto(param) (response) {
	return methods.PhotoSaveMarketAlbumPhoto(c.client(), param)
}

func (c *Client) PhotoSaveMarketPhoto(param) (response) {
	return methods.PhotoSaveMarketPhoto(c.client(), param)
}

func (c *Client) PhotoSaveMessagesPhoto(param) (response) {
	return methods.PhotoSaveMessagesPhoto(c.client(), param)
}

func (c *Client) PhotoSaveOwnerCoverPhoto(param) (response) {
	return methods.PhotoSaveOwnerCoverPhoto(c.client(), param)
}

func (c *Client) PhotoSaveOwnerPhoto(param) (response) {
	return methods.PhotoSaveOwnerPhoto(c.client(), param)
}

func (c *Client) PhotoSaveWallPhoto(param) (response) {
	return methods.PhotoSaveWallPhoto(c.client(), param)
}

func (c *Client) PhotoSearch(param) (response) {
	return methods.PhotoSearch(c.client(), param)
}

func (c *Client) PollsAddVote(param) (response) {
	return methods.PollsAddVote(c.client(), param)
}

func (c *Client) PollsCreate(param) (response) {
	return methods.PollsCreate(c.client(), param)
}

func (c *Client) PollsDeleteVote(param) (response) {
	return methods.PollsDeleteVote(c.client(), param)
}

func (c *Client) PollsEdit(param) (response) {
	return methods.PollsEdit(c.client(), param)
}

func (c *Client) PollsGetBackgrounds(param) (response) {
	return methods.PollsGetBackgrounds(c.client(), param)
}

func (c *Client) PollsGetById(param) (response) {
	return methods.PollsGetById(c.client(), param)
}

func (c *Client) PollsGetPhotoUploadServer(param) (response) {
	return methods.PollsGetPhotoUploadServer(c.client(), param)
}

func (c *Client) PollsGetVoters(param) (response) {
	return methods.PollsGetVoters(c.client(), param)
}

func (c *Client) PollsSavePhoto(param) (response) {
	return methods.PollsSavePhoto(c.client(), param)
}

func (c *Client) PrettycardsCreate(param) (response) {
	return methods.PrettycardsCreate(c.client(), param)
}

func (c *Client) PrettycardsDelete(param) (response) {
	return methods.PrettycardsDelete(c.client(), param)
}

func (c *Client) PrettycardsEdit(param) (response) {
	return methods.PrettycardsEdit(c.client(), param)
}

func (c *Client) PrettycardsGet(param) (response) {
	return methods.PrettycardsGet(c.client(), param)
}

func (c *Client) PrettycardsGetById(param) (response) {
	return methods.PrettycardsGetById(c.client(), param)
}

func (c *Client) PrettycardsGetUploadURL(param) (response) {
	return methods.PrettycardsGetUploadURL(c.client(), param)
}

func (c *Client) SearchGetHints(param) (response) {
	return methods.SearchGetHints(c.client(), param)
}

func (c *Client) SecureAddAppEvent(param) (response) {
	return methods.SecureAddAppEvent(c.client(), param)
}

func (c *Client) SecureCheckToken(param) (response) {
	return methods.SecureCheckToken(c.client(), param)
}

func (c *Client) SecureGetAppBalance(param) (response) {
	return methods.SecureGetAppBalance(c.client(), param)
}

func (c *Client) SecureGetSMSHistory(param) (response) {
	return methods.SecureGetSMSHistory(c.client(), param)
}

func (c *Client) SecureGetTransactionsHistory(param) (response) {
	return methods.SecureGetTransactionsHistory(c.client(), param)
}

func (c *Client) SecureGetUserLevel(param) (response) {
	return methods.SecureGetUserLevel(c.client(), param)
}

func (c *Client) SecureGiveEventSticker(param) (response) {
	return methods.SecureGiveEventSticker(c.client(), param)
}

func (c *Client) SecureSendNotification(param) (response) {
	return methods.SecureSendNotification(c.client(), param)
}

func (c *Client) SecureSendSMSNotification(param) (response) {
	return methods.SecureSendSMSNotification(c.client(), param)
}

func (c *Client) SecureSetCounter(param) (response) {
	return methods.SecureSetCounter(c.client(), param)
}

func (c *Client) StatsGet(param) (response) {
	return methods.StatsGet(c.client(), param)
}

func (c *Client) StatsGetPostReach(param) (response) {
	return methods.StatsGetPostReach(c.client(), param)
}

func (c *Client) StatsTrackVisitor(param) (response) {
	return methods.StatsTrackVisitor(c.client(), param)
}

func (c *Client) StatusGet(param) (response) {
	return methods.StatusGet(c.client(), param)
}

func (c *Client) StatusSet(param) (response) {
	return methods.StatusSet(c.client(), param)
}

*/

func (c *Client) StorageGet(keys []string, userId int) (map[string]string, error) {
	return methods.StorageGet(c.client(), keys, userId)
}

/*

func (c *Client) StorageGetKeys(param) (response) {
	return methods.StorageGetKeys(c.client(), param)
}

*/

func (c *Client) StorageSet(key string, value interface{}, userId int) error {
	return methods.StorageSet(c.client(), key, value, userId)
}

/*

func (c *Client) StoriesBanOwner(param) (response) {
	return methods.StoriesBanOwner(c.client(), param)
}

func (c *Client) StoriesDelete(param) (response) {
	return methods.StoriesDelete(c.client(), param)
}

func (c *Client) StoriesGet(param) (response) {
	return methods.StoriesGet(c.client(), param)
}

func (c *Client) StoriesGetBanned(param) (response) {
	return methods.StoriesGetBanned(c.client(), param)
}

func (c *Client) StoriesGetById(param) (response) {
	return methods.StoriesGetById(c.client(), param)
}

func (c *Client) StoriesGetPhotoUploadServer(param) (response) {
	return methods.StoriesGetPhotoUploadServer(c.client(), param)
}

func (c *Client) StoriesGetReplies(param) (response) {
	return methods.StoriesGetReplies(c.client(), param)
}

func (c *Client) StoriesGetStats(param) (response) {
	return methods.StoriesGetStats(c.client(), param)
}

func (c *Client) StoriesGetVideoUploadServer(param) (response) {
	return methods.StoriesGetVideoUploadServer(c.client(), param)
}

func (c *Client) StoriesGetViewers(param) (response) {
	return methods.StoriesGetViewers(c.client(), param)
}

func (c *Client) StoriesHideAllReplies(param) (response) {
	return methods.StoriesHideAllReplies(c.client(), param)
}

func (c *Client) StoriesHideReply(param) (response) {
	return methods.StoriesHideReply(c.client(), param)
}

func (c *Client) StoriesSearch(param) (response) {
	return methods.StoriesSearch(c.client(), param)
}

func (c *Client) StoriesUnbanOwner(param) (response) {
	return methods.StoriesUnbanOwner(c.client(), param)
}

*/

func (c *Client) StreamingGetServerUrl() (*methods.StreamingGetServerUrlResponse, error) {
	return methods.StreamingGetServerUrl(c.client())
}

/*

func (c *Client) StreamingGetSettings(param) (response) {
	return methods.StreamingGetSettings(c.client(), param)
}

func (c *Client) StreamingGetStats(param) (response) {
	return methods.StreamingGetStats(c.client(), param)
}

func (c *Client) StreamingGetStem(param) (response) {
	return methods.StreamingGetStem(c.client(), param)
}

func (c *Client) StreamingSetSettings(param) (response) {
	return methods.StreamingSetSettings(c.client(), param)
}

*/

func (c *Client) UsersGet(userIds []int, fields []string) (*methods.UsersGetResponse, error) {
	return methods.UsersGet(c.client(), userIds, fields)
}

func (c *Client) UsersGetRegDate(userId int) (*methods.UsersGetRegDateResponse, error) {
	return methods.UsersGetRegDate(c.client(), userId)
}

/*

func (c *Client) UsersGetFollowers(param) (response) {
	return methods.UsersGetFollowers(c.client(), param)
}

func (c *Client) UsersGetSubscriptions(param) (response) {
	return methods.UsersGetSubscriptions(c.client(), param)
}

func (c *Client) UsersReport(param) (response) {
	return methods.UsersReport(c.client(), param)
}

func (c *Client) UsersSearch(param) (response) {
	return methods.UsersSearch(c.client(), param)
}

func (c *Client) UtilsCheckLink(param) (response) {
	return methods.UtilsCheckLink(c.client(), param)
}

func (c *Client) UtilsDeleteFromLastShortened(param) (response) {
	return methods.UtilsDeleteFromLastShortened(c.client(), param)
}

func (c *Client) UtilsGetLastShortenedLinks(param) (response) {
	return methods.UtilsGetLastShortenedLinks(c.client(), param)
}

func (c *Client) UtilsGetLinkStats(param) (response) {
	return methods.UtilsGetLinkStats(c.client(), param)
}

func (c *Client) UtilsGetServerTime(param) (response) {
	return methods.UtilsGetServerTime(c.client(), param)
}

func (c *Client) UtilsGetShortLink(param) (response) {
	return methods.UtilsGetShortLink(c.client(), param)
}

func (c *Client) UtilsResolveScreenName(param) (response) {
	return methods.UtilsResolveScreenName(c.client(), param)
}

func (c *Client) VideoAdd(param) (response) {
	return methods.VideoAdd(c.client(), param)
}

func (c *Client) VideoAddAlbum(param) (response) {
	return methods.VideoAddAlbum(c.client(), param)
}

func (c *Client) VideoAddToAlbum(param) (response) {
	return methods.VideoAddToAlbum(c.client(), param)
}

func (c *Client) VideoCreateComment(param) (response) {
	return methods.VideoCreateComment(c.client(), param)
}

func (c *Client) VideoDelete(param) (response) {
	return methods.VideoDelete(c.client(), param)
}

func (c *Client) VideoDeleteAlbum(param) (response) {
	return methods.VideoDeleteAlbum(c.client(), param)
}

func (c *Client) VideoDeleteComment(param) (response) {
	return methods.VideoDeleteComment(c.client(), param)
}

func (c *Client) VideoEdit(param) (response) {
	return methods.VideoEdit(c.client(), param)
}

func (c *Client) VideoEditAlbum(param) (response) {
	return methods.VideoEditAlbum(c.client(), param)
}

func (c *Client) VideoEditComment(param) (response) {
	return methods.VideoEditComment(c.client(), param)
}

func (c *Client) VideoGet(param) (response) {
	return methods.VideoGet(c.client(), param)
}

func (c *Client) VideoGetAlbumById(param) (response) {
	return methods.VideoGetAlbumById(c.client(), param)
}

func (c *Client) VideoGetAlbums(param) (response) {
	return methods.VideoGetAlbums(c.client(), param)
}

func (c *Client) VideoGetAlbumsByVideo(param) (response) {
	return methods.VideoGetAlbumsByVideo(c.client(), param)
}

func (c *Client) VideoGetComments(param) (response) {
	return methods.VideoGetComments(c.client(), param)
}

func (c *Client) VideoRemoveFromAlbum(param) (response) {
	return methods.VideoRemoveFromAlbum(c.client(), param)
}

func (c *Client) VideoReorderAlbums(param) (response) {
	return methods.VideoReorderAlbums(c.client(), param)
}

func (c *Client) VideoReorderVideos(param) (response) {
	return methods.VideoReorderVideos(c.client(), param)
}

func (c *Client) VideoReport(param) (response) {
	return methods.VideoReport(c.client(), param)
}

func (c *Client) VideoReportComment(param) (response) {
	return methods.VideoReportComment(c.client(), param)
}

func (c *Client) VideoRestore(param) (response) {
	return methods.VideoRestore(c.client(), param)
}

func (c *Client) VideoRestoreComment(param) (response) {
	return methods.VideoRestoreComment(c.client(), param)
}

func (c *Client) VideoSave(param) (response) {
	return methods.VideoSave(c.client(), param)
}

func (c *Client) VideoSearch(param) (response) {
	return methods.VideoSearch(c.client(), param)
}

func (c *Client) WallCloseComments(param) (response) {
	return methods.WallCloseComments(c.client(), param)
}

func (c *Client) WallCreateComment(param) (response) {
	return methods.WallCreateComment(c.client(), param)
}

func (c *Client) WallDelete(param) (response) {
	return methods.WallDelete(c.client(), param)
}

func (c *Client) WallDeleteComment(param) (response) {
	return methods.WallDeleteComment(c.client(), param)
}

func (c *Client) WallEdit(param) (response) {
	return methods.WallEdit(c.client(), param)
}

func (c *Client) WallEditAdsStealth(param) (response) {
	return methods.WallEditAdsStealth(c.client(), param)
}

func (c *Client) WallEditComment(param) (response) {
	return methods.WallEditComment(c.client(), param)
}

*/

func (c *Client) WallGet(Owner int, filter string, maxPosts int) (*methods.WallGetResponse, error) {
	return methods.WallGet(c.client(), Owner, filter, MAX_COUNT_POSTS)
}

/*

func (c *Client) WallGetById(param) (response) {
	return methods.WallGetById(c.client(), param)
}

func (c *Client) WallGetComment(param) (response) {
	return methods.WallGetComment(c.client(), param)
}

func (c *Client) WallGetComments(param) (response) {
	return methods.WallGetComments(c.client(), param)
}

func (c *Client) WallGetReposts(param) (response) {
	return methods.WallGetReposts(c.client(), param)
}

func (c *Client) WallOpenComments(param) (response) {
	return methods.WallOpenComments(c.client(), param)
}

func (c *Client) WallPin(param) (response) {
	return methods.WallPin(c.client(), param)
}

func (c *Client) WallPost(param) (response) {
	return methods.WallPost(c.client(), param)
}

func (c *Client) WallPostAdsStealth(param) (response) {
	return methods.WallPostAdsStealth(c.client(), param)
}

func (c *Client) WallReportComment(param) (response) {
	return methods.WallReportComment(c.client(), param)
}

func (c *Client) WallReportPost(param) (response) {
	return methods.WallReportPost(c.client(), param)
}

func (c *Client) WallRepost(param) (response) {
	return methods.WallRepost(c.client(), param)
}

func (c *Client) WallRestore(param) (response) {
	return methods.WallRestore(c.client(), param)
}

func (c *Client) WallRestoreComment(param) (response) {
	return methods.WallRestoreComment(c.client(), param)
}

func (c *Client) WallSearch(param) (response) {
	return methods.WallSearch(c.client(), param)
}

func (c *Client) WallUnpin(param) (response) {
	return methods.WallUnpin(c.client(), param)
}

func (c *Client) WidgetsGetComments(param) (response) {
	return methods.WidgetsGetComments(c.client(), param)
}

func (c *Client) WidgetsGetPages(param) (response) {
	return methods.WidgetsGetPages(c.client(), param)
}

*/
