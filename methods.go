package vk

import (
	"reflect"

	"github.com/xelaj/vk/methods"
	"github.com/xelaj/vk/types"
)

func (c *Client) AccountBan(ownerId int) error {
	return methods.AccountBan(c.Client(), ownerId)
}

/*

func (c *Client) AccountGetActiveOffers(param) (response) {
	return methods.AccountGetActiveOffers(c.Client(), param)
}

*/

func (c *Client) AccountGetAppPermissions(userId int) (types.Permissions, error) {
	return methods.AccountGetAppPermissions(c.Client(), userId)
}

func (c *Client) AccountGetBanned(offset, count int) (*methods.AccountGetBannedResponse, error) {
	return methods.AccountGetBanned(c.Client(), offset, count)
}

/*

func (c *Client) AccountGetCounters(param) (response) {
	return methods.AccountGetCounters(c.Client(), param)
}

func (c *Client) AccountGetInfo(param) (response) {
	return methods.AccountGetInfo(c.Client(), param)
}

func (c *Client) AccountGetProfileInfo(param) (response) {
	return methods.AccountGetProfileInfo(c.Client(), param)
}

func (c *Client) AccountGetPushSettings(param) (response) {
	return methods.AccountGetPushSettings(c.Client(), param)
}

func (c *Client) AccountRegisterDevice(param) (response) {
	return methods.AccountRegisterDevice(c.Client(), param)
}

func (c *Client) AccountSaveProfileInfo(param) (response) {
	return methods.AccountSaveProfileInfo(c.Client(), param)
}

func (c *Client) AccountSetInfo(param) (response) {
	return methods.AccountSetInfo(c.Client(), param)
}

func (c *Client) AccountGetProfileInfo(param) (response) {
	return methods.AccountGetProfileInfo(c.Client(), param)
}

func (c *Client) AccountSetOffline(param) (response) {
	return methods.AccountSetOffline(c.Client(), param)
}

func (c *Client) AccountSetOnline(param) (response) {
	return methods.AccountSetOnline(c.Client(), param)
}

func (c *Client) AccountSetPushSettings(param) (response) {
	return methods.AccountSetPushSettings(c.Client(), param)
}

func (c *Client) AccountSetSilenceMode(param) (response) {
	return methods.AccountSetSilenceMode(c.Client(), param)
}

*/

func (c *Client) AccountUnban(ownerId int) error {
	return methods.AccountUnban(c.Client(), ownerId)
}

/*

func (c *Client) AccountUnregisterDevice(param) (response) {
	return methods.AccountUnregisterDevice(c.Client(), param)
}

func (c *Client) AdsAddOfficeUsers(param) (response) {
	return methods.AdsAddOfficeUsers(c.Client(), param)
}

func (c *Client) AdsCheckLink(param) (response) {
	return methods.AdsCheckLink(c.Client(), param)
}

func (c *Client) AdsCreateAds(param) (response) {
	return methods.AdsCreateAds(c.Client(), param)
}

func (c *Client) AdsCreateCampaigns(param) (response) {
	return methods.AdsCreateCampaigns(c.Client(), param)
}

func (c *Client) AdsCreateClients(param) (response) {
	return methods.AdsCreateClients(c.Client(), param)
}

func (c *Client) AdsCreateLookalikeRequest(param) (response) {
	return methods.AdsCreateLookalikeRequest(c.Client(), param)
}

func (c *Client) AdsCreateTargetGroup(param) (response) {
	return methods.AdsCreateTargetGroup(c.Client(), param)
}

func (c *Client) AdsCreateTargetPixel(param) (response) {
	return methods.AdsCreateTargetPixel(c.Client(), param)
}

func (c *Client) AdsDeleteAds(param) (response) {
	return methods.AdsDeleteAds(c.Client(), param)
}

func (c *Client) AdsDeleteCampaigns(param) (response) {
	return methods.AdsDeleteCampaigns(c.Client(), param)
}

func (c *Client) AdsDeleteClients(param) (response) {
	return methods.AdsDeleteClients(c.Client(), param)
}

func (c *Client) AdsDeleteTargetGroup(param) (response) {
	return methods.AdsDeleteTargetGroup(c.Client(), param)
}

func (c *Client) AdsDeleteTargetPixel(param) (response) {
	return methods.AdsDeleteTargetPixel(c.Client(), param)
}

func (c *Client) AdsGetAccounts(param) (response) {
	return methods.AdsGetAccounts(c.Client(), param)
}

func (c *Client) AdsGetAds(param) (response) {
	return methods.AdsGetAds(c.Client(), param)
}

func (c *Client) AdsGetAdsLayout(param) (response) {
	return methods.AdsGetAdsLayout(c.Client(), param)
}

func (c *Client) AdsGetAdsTargeting(param) (response) {
	return methods.AdsGetAdsTargeting(c.Client(), param)
}

func (c *Client) AdsGetBudget(param) (response) {
	return methods.AdsGetBudget(c.Client(), param)
}

func (c *Client) AdsGetCampaigns(param) (response) {
	return methods.AdsGetCampaigns(c.Client(), param)
}

func (c *Client) AdsGetCategories(param) (response) {
	return methods.AdsGetCategories(c.Client(), param)
}

func (c *Client) AdsGetClients(param) (response) {
	return methods.AdsGetClients(c.Client(), param)
}

func (c *Client) AdsGetDemographics(param) (response) {
	return methods.AdsGetDemographics(c.Client(), param)
}

func (c *Client) AdsGetFloodStats(param) (response) {
	return methods.AdsGetFloodStats(c.Client(), param)
}

func (c *Client) AdsGetLookalikeRequests(param) (response) {
	return methods.AdsGetLookalikeRequests(c.Client(), param)
}

func (c *Client) AdsGetOfficeUsers(param) (response) {
	return methods.AdsGetOfficeUsers(c.Client(), param)
}

func (c *Client) AdsGetPostsReach(param) (response) {
	return methods.AdsGetPostsReach(c.Client(), param)
}

func (c *Client) AdsGetRejectionReason(param) (response) {
	return methods.AdsGetRejectionReason(c.Client(), param)
}

func (c *Client) AdsGetStatistics(param) (response) {
	return methods.AdsGetStatistics(c.Client(), param)
}

func (c *Client) AdsGetSuggestions(param) (response) {
	return methods.AdsGetSuggestions(c.Client(), param)
}

func (c *Client) AdsGetTargetGroups(param) (response) {
	return methods.AdsGetTargetGroups(c.Client(), param)
}

func (c *Client) AdsGetTargetPixels(param) (response) {
	return methods.AdsGetTargetPixels(c.Client(), param)
}

func (c *Client) AdsGetTargetingStats(param) (response) {
	return methods.AdsGetTargetingStats(c.Client(), param)
}

func (c *Client) AdsGetUploadURL(param) (response) {
	return methods.AdsGetUploadURL(c.Client(), param)
}

func (c *Client) AdsGetVideoUploadURL(param) (response) {
	return methods.AdsGetVideoUploadURL(c.Client(), param)
}

func (c *Client) AdsImportTargetContacts(param) (response) {
	return methods.AdsImportTargetContacts(c.Client(), param)
}

func (c *Client) AdsRemoveOfficeUsers(param) (response) {
	return methods.AdsRemoveOfficeUsers(c.Client(), param)
}

func (c *Client) AdsRemoveTargetContacts(param) (response) {
	return methods.AdsRemoveTargetContacts(c.Client(), param)
}

func (c *Client) AdsSaveLookalikeRequestResult(param) (response) {
	return methods.AdsSaveLookalikeRequestResult(c.Client(), param)
}

func (c *Client) AdsShareTargetGroup(param) (response) {
	return methods.AdsShareTargetGroup(c.Client(), param)
}

func (c *Client) AdsUpdateAds(param) (response) {
	return methods.AdsUpdateAds(c.Client(), param)
}

func (c *Client) AdsUpdateCampaigns(param) (response) {
	return methods.AdsUpdateCampaigns(c.Client(), param)
}

func (c *Client) AdsUpdateClients(param) (response) {
	return methods.AdsUpdateClients(c.Client(), param)
}

func (c *Client) AdsUpdateTargetGroup(param) (response) {
	return methods.AdsUpdateTargetGroup(c.Client(), param)
}

func (c *Client) AdsUpdateTargetPixel(param) (response) {
	return methods.AdsUpdateTargetPixel(c.Client(), param)
}

func (c *Client) AppwdigetsGetAppImageUploadServer(param) (response) {
	return methods.AppwdigetsGetAppImageUploadServer(c.Client(), param)
}

func (c *Client) AppwdigetsGetAppImages(param) (response) {
	return methods.AppwdigetsGetAppImages(c.Client(), param)
}

func (c *Client) AppwdigetsGetGroupImageUploadServer(param) (response) {
	return methods.AppwdigetsGetGroupImageUploadServer(c.Client(), param)
}

func (c *Client) AppwdigetsGetGroupImages(param) (response) {
	return methods.AppwdigetsGetGroupImages(c.Client(), param)
}

func (c *Client) AppwdigetsGetImagesById(param) (response) {
	return methods.AppwdigetsGetImagesById(c.Client(), param)
}

func (c *Client) AppwdigetsSaveAppImage(param) (response) {
	return methods.AppwdigetsSaveAppImage(c.Client(), param)
}

func (c *Client) AppwdigetsSaveGroupImage(param) (response) {
	return methods.AppwdigetsSaveGroupImage(c.Client(), param)
}

func (c *Client) AppwdigetsUpdate(param) (response) {
	return methods.AppwdigetsUpdate(c.Client(), param)
}

func (c *Client) AppsDeleteAppRequests(param) (response) {
	return methods.AppsDeleteAppRequests(c.Client(), param)
}

*/

func (c *Client) AppsGet(additional map[string]interface{}) (*methods.AppsGetResponse, error) {
	return methods.AppsGet(c.Client(), additional)
}

/*

func (c *Client) AppsGetCatalog(param) (response) {
	return methods.AppsGetCatalog(c.Client(), param)
}

func (c *Client) AppsGetFriendsList(param) (response) {
	return methods.AppsGetFriendsList(c.Client(), param)
}

func (c *Client) AppsGetLeaderboard(param) (response) {
	return methods.AppsGetLeaderboard(c.Client(), param)
}

func (c *Client) AppsGetScopes(param) (response) {
	return methods.AppsGetScopes(c.Client(), param)
}

func (c *Client) AppsGetScore(param) (response) {
	return methods.AppsGetScore(c.Client(), param)
}

func (c *Client) AppsPromoHasActiveGift(param) (response) {
	return methods.AppsPromoHasActiveGift(c.Client(), param)
}

func (c *Client) AppsPromoUseGift(param) (response) {
	return methods.AppsPromoUseGift(c.Client(), param)
}

func (c *Client) AppsSendRequest(param) (response) {
	return methods.AppsSendRequest(c.Client(), param)
}

func (c *Client) AudioGet(param) (response) {
	return methods.AudioGet(c.Client(), param)
}

func (c *Client) AudioGetById(param) (response) {
	return methods.AudioGetById(c.Client(), param)
}

func (c *Client) AudioGetLyrics(param) (response) {
	return methods.AudioGetLyrics(c.Client(), param)
}

func (c *Client) AudioSearch(param) (response) {
	return methods.AudioSearch(c.Client(), param)
}

func (c *Client) AudioGetUploadServer(param) (response) {
	return methods.AudioGetUploadServer(c.Client(), param)
}

func (c *Client) AudioSave(param) (response) {
	return methods.AudioSave(c.Client(), param)
}

func (c *Client) AudioAdd(param) (response) {
	return methods.AudioAdd(c.Client(), param)
}

func (c *Client) AudioDelete(param) (response) {
	return methods.AudioDelete(c.Client(), param)
}

func (c *Client) AudioEdit(param) (response) {
	return methods.AudioEdit(c.Client(), param)
}

func (c *Client) AudioReorder(param) (response) {
	return methods.AudioReorder(c.Client(), param)
}

func (c *Client) AudioRestore(param) (response) {
	return methods.AudioRestore(c.Client(), param)
}

func (c *Client) AudioGetAlbums(param) (response) {
	return methods.AudioGetAlbums(c.Client(), param)
}

func (c *Client) AudioAddAlbum(param) (response) {
	return methods.AudioAddAlbum(c.Client(), param)
}

func (c *Client) AudioEditAlbum(param) (response) {
	return methods.AudioEditAlbum(c.Client(), param)
}

func (c *Client) AudioDeleteAlbum(param) (response) {
	return methods.AudioDeleteAlbum(c.Client(), param)
}

func (c *Client) AudioMoveToAlbum(param) (response) {
	return methods.AudioMoveToAlbum(c.Client(), param)
}

func (c *Client) AudioSetBroadcast(param) (response) {
	return methods.AudioSetBroadcast(c.Client(), param)
}

func (c *Client) AudioGetBroadcastList(param) (response) {
	return methods.AudioGetBroadcastList(c.Client(), param)
}

func (c *Client) AudioGetRecommendations(param) (response) {
	return methods.AudioGetRecommendations(c.Client(), param)
}

func (c *Client) AudioGetPopular(param) (response) {
	return methods.AudioGetPopular(c.Client(), param)
}

func (c *Client) AudioGetCount(param) (response) {
	return methods.AudioGetCount(c.Client(), param)
}

func (c *Client) AuthCheckPhone(param) (response) {
	return methods.AuthCheckPhone(c.Client(), param)
}

func (c *Client) AuthRestore(param) (response) {
	return methods.AuthRestore(c.Client(), param)
}

func (c *Client) BoardAddTopic(param) (response) {
	return methods.BoardAddTopic(c.Client(), param)
}

func (c *Client) BoardCloseTopic(param) (response) {
	return methods.BoardCloseTopic(c.Client(), param)
}

func (c *Client) BoardCreateComment(param) (response) {
	return methods.BoardCreateComment(c.Client(), param)
}

func (c *Client) BoardDeleteComment(param) (response) {
	return methods.BoardDeleteComment(c.Client(), param)
}

func (c *Client) BoardDeleteTopic(param) (response) {
	return methods.BoardDeleteTopic(c.Client(), param)
}

func (c *Client) BoardEditComment(param) (response) {
	return methods.BoardEditComment(c.Client(), param)
}

func (c *Client) BoardEditTopic(param) (response) {
	return methods.BoardEditTopic(c.Client(), param)
}

func (c *Client) BoardFixTopic(param) (response) {
	return methods.BoardFixTopic(c.Client(), param)
}

func (c *Client) BoardGetComments(param) (response) {
	return methods.BoardGetComments(c.Client(), param)
}

func (c *Client) BoardGetTopics(param) (response) {
	return methods.BoardGetTopics(c.Client(), param)
}

func (c *Client) BoardOpenTopic(param) (response) {
	return methods.BoardOpenTopic(c.Client(), param)
}

func (c *Client) BoardRestoreComment(param) (response) {
	return methods.BoardRestoreComment(c.Client(), param)
}

func (c *Client) BoardUnfixTopic(param) (response) {
	return methods.BoardUnfixTopic(c.Client(), param)
}

func (c *Client) DatabaseGetChairs(param) (response) {
	return methods.DatabaseGetChairs(c.Client(), param)
}

func (c *Client) DatabaseGetCities(param) (response) {
	return methods.DatabaseGetCities(c.Client(), param)
}

func (c *Client) DatabaseGetCitiesById(param) (response) {
	return methods.DatabaseGetCitiesById(c.Client(), param)
}

func (c *Client) DatabaseGetCountries(param) (response) {
	return methods.DatabaseGetCountries(c.Client(), param)
}

func (c *Client) DatabaseGetCountriesById(param) (response) {
	return methods.DatabaseGetCountriesById(c.Client(), param)
}

func (c *Client) DatabaseGetFaculties(param) (response) {
	return methods.DatabaseGetFaculties(c.Client(), param)
}

func (c *Client) DatabaseGetMetroStations(param) (response) {
	return methods.DatabaseGetMetroStations(c.Client(), param)
}

func (c *Client) DatabaseGetMetroStationsById(param) (response) {
	return methods.DatabaseGetMetroStationsById(c.Client(), param)
}

func (c *Client) DatabaseGetRegions(param) (response) {
	return methods.DatabaseGetRegions(c.Client(), param)
}

func (c *Client) DatabaseGetSchoolClasses(param) (response) {
	return methods.DatabaseGetSchoolClasses(c.Client(), param)
}

func (c *Client) DatabaseGetSchools(param) (response) {
	return methods.DatabaseGetSchools(c.Client(), param)
}

func (c *Client) DatabaseGetUniversities(param) (response) {
	return methods.DatabaseGetUniversities(c.Client(), param)
}

func (c *Client) DocsAdd(param) (response) {
	return methods.DocsAdd(c.Client(), param)
}

func (c *Client) DocsDelete(param) (response) {
	return methods.DocsDelete(c.Client(), param)
}

func (c *Client) DocsEdit(param) (response) {
	return methods.DocsEdit(c.Client(), param)
}

func (c *Client) DocsGet(param) (response) {
	return methods.DocsGet(c.Client(), param)
}

func (c *Client) DocsGetById(param) (response) {
	return methods.DocsGetById(c.Client(), param)
}

func (c *Client) DocsGetMessagesUploadServer(param) (response) {
	return methods.DocsGetMessagesUploadServer(c.Client(), param)
}

func (c *Client) DocsGetTypes(param) (response) {
	return methods.DocsGetTypes(c.Client(), param)
}

func (c *Client) DocsGetUploadServer(param) (response) {
	return methods.DocsGetUploadServer(c.Client(), param)
}

func (c *Client) DocsGetWallUploadServer(param) (response) {
	return methods.DocsGetWallUploadServer(c.Client(), param)
}

func (c *Client) DocsSave(param) (response) {
	return methods.DocsSave(c.Client(), param)
}

func (c *Client) DocsSearch(param) (response) {
	return methods.DocsSearch(c.Client(), param)
}

*/

func (c *Client) Execucte(method string, params map[string]interface{}, resType reflect.Type) (interface{}, error) {
	return methods.ExecuteMethod(c.Client(), method, params, resType)
}

/*

func (c *Client) FaveAddArticle(param) (response) {
	return methods.FaveAddArticle(c.Client(), param)
}

func (c *Client) FaveAddLink(param) (response) {
	return methods.FaveAddLink(c.Client(), param)
}

func (c *Client) FaveAddPage(param) (response) {
	return methods.FaveAddPage(c.Client(), param)
}

func (c *Client) FaveAddPost(param) (response) {
	return methods.FaveAddPost(c.Client(), param)
}

func (c *Client) FaveAddProduct(param) (response) {
	return methods.FaveAddProduct(c.Client(), param)
}

func (c *Client) FaveAddTag(param) (response) {
	return methods.FaveAddTag(c.Client(), param)
}

func (c *Client) FaveAddVideo(param) (response) {
	return methods.FaveAddVideo(c.Client(), param)
}

func (c *Client) FaveEditTag(param) (response) {
	return methods.FaveEditTag(c.Client(), param)
}

func (c *Client) FaveGet(param) (response) {
	return methods.FaveGet(c.Client(), param)
}

func (c *Client) FaveGetPages(param) (response) {
	return methods.FaveGetPages(c.Client(), param)
}

func (c *Client) FaveGetTags(param) (response) {
	return methods.FaveGetTags(c.Client(), param)
}

func (c *Client) FaveMarkSeen(param) (response) {
	return methods.FaveMarkSeen(c.Client(), param)
}

func (c *Client) FaveRemoveArticle(param) (response) {
	return methods.FaveRemoveArticle(c.Client(), param)
}

func (c *Client) FaveRemoveLink(param) (response) {
	return methods.FaveRemoveLink(c.Client(), param)
}

func (c *Client) FaveRemovePage(param) (response) {
	return methods.FaveRemovePage(c.Client(), param)
}

func (c *Client) FaveRemovePost(param) (response) {
	return methods.FaveRemovePost(c.Client(), param)
}

func (c *Client) FaveRemoveProduct(param) (response) {
	return methods.FaveRemoveProduct(c.Client(), param)
}

func (c *Client) FaveRemoveTag(param) (response) {
	return methods.FaveRemoveTag(c.Client(), param)
}

func (c *Client) FaveRemoveVideo(param) (response) {
	return methods.FaveRemoveVideo(c.Client(), param)
}

func (c *Client) FaveReorderTags(param) (response) {
	return methods.FaveReorderTags(c.Client(), param)
}

func (c *Client) FaveSetPageTags(param) (response) {
	return methods.FaveSetPageTags(c.Client(), param)
}

func (c *Client) FaveSetTags(param) (response) {
	return methods.FaveSetTags(c.Client(), param)
}

func (c *Client) FaveTrackPageInteraction(param) (response) {
	return methods.FaveTrackPageInteraction(c.Client(), param)
}

func (c *Client) FriendsAdd(param) (response) {
	return methods.FriendsAdd(c.Client(), param)
}

func (c *Client) FriendsAddList(param) (response) {
	return methods.FriendsAddList(c.Client(), param)
}

func (c *Client) FriendsAreFriends(param) (response) {
	return methods.FriendsAreFriends(c.Client(), param)
}

func (c *Client) FriendsDelete(param) (response) {
	return methods.FriendsDelete(c.Client(), param)
}

func (c *Client) FriendsDeleteAllRequests(param) (response) {
	return methods.FriendsDeleteAllRequests(c.Client(), param)
}

func (c *Client) FriendsDeleteList(param) (response) {
	return methods.FriendsDeleteList(c.Client(), param)
}

func (c *Client) FriendsEdit(param) (response) {
	return methods.FriendsEdit(c.Client(), param)
}

func (c *Client) FriendsEditList(param) (response) {
	return methods.FriendsEditList(c.Client(), param)
}

*/


func (c *Client) FriendsGet(userId int, Order string, fields []string, nameCase string) (*methods.FriendsGetResponse, error) {
	return methods.FriendsGet(c.Client(), userId, Order, fields, nameCase, MAX_COUNT_FRIENDS)
}

func (c *Client) FriendsGetId(userId int, Order string, nameCase string) (*methods.FriendsGetIdResponse, error) {
	return methods.FriendsGetId(c.Client(), userId, Order, nameCase, MAX_COUNT_FRIENDS)
}

/*

func (c *Client) FriendsGetAppUsers(param) (response) {
	return methods.FriendsGetAppUsers(c.Client(), param)
}

func (c *Client) FriendsGetByPhones(param) (response) {
	return methods.FriendsGetByPhones(c.Client(), param)
}

func (c *Client) FriendsGetLists(param) (response) {
	return methods.FriendsGetLists(c.Client(), param)
}

func (c *Client) FriendsGetMutual(param) (response) {
	return methods.FriendsGetMutual(c.Client(), param)
}

func (c *Client) FriendsGetOnline(param) (response) {
	return methods.FriendsGetOnline(c.Client(), param)
}

func (c *Client) FriendsGetRecent(param) (response) {
	return methods.FriendsGetRecent(c.Client(), param)
}

func (c *Client) FriendsGetRequests(param) (response) {
	return methods.FriendsGetRequests(c.Client(), param)
}

func (c *Client) FriendsGetSuggestions(param) (response) {
	return methods.FriendsGetSuggestions(c.Client(), param)
}

func (c *Client) FriendsSearch(param) (response) {
	return methods.FriendsSearch(c.Client(), param)
}

func (c *Client) GiftsGet(param) (response) {
	return methods.GiftsGet(c.Client(), param)
}

func (c *Client) GroupsAddAddress(param) (response) {
	return methods.GroupsAddAddress(c.Client(), param)
}

func (c *Client) GroupsAddCallbackServer(param) (response) {
	return methods.GroupsAddCallbackServer(c.Client(), param)
}

func (c *Client) GroupsAddLink(param) (response) {
	return methods.GroupsAddLink(c.Client(), param)
}

func (c *Client) GroupsApproveRequest(param) (response) {
	return methods.GroupsApproveRequest(c.Client(), param)
}

func (c *Client) GroupsBan(param) (response) {
	return methods.GroupsBan(c.Client(), param)
}

func (c *Client) GroupsCreate(param) (response) {
	return methods.GroupsCreate(c.Client(), param)
}

func (c *Client) GroupsDeleteAddress(param) (response) {
	return methods.GroupsDeleteAddress(c.Client(), param)
}

func (c *Client) GroupsDeleteCallbackServer(param) (response) {
	return methods.GroupsDeleteCallbackServer(c.Client(), param)
}

func (c *Client) GroupsDeleteLink(param) (response) {
	return methods.GroupsDeleteLink(c.Client(), param)
}

func (c *Client) GroupsDisableOnline(param) (response) {
	return methods.GroupsDisableOnline(c.Client(), param)
}

func (c *Client) GroupsEdit(param) (response) {
	return methods.GroupsEdit(c.Client(), param)
}

func (c *Client) GroupsEditAddress(param) (response) {
	return methods.GroupsEditAddress(c.Client(), param)
}

func (c *Client) GroupsEditCallbackServer(param) (response) {
	return methods.GroupsEditCallbackServer(c.Client(), param)
}

func (c *Client) GroupsEditLink(param) (response) {
	return methods.GroupsEditLink(c.Client(), param)
}

func (c *Client) GroupsEditManager(param) (response) {
	return methods.GroupsEditManager(c.Client(), param)
}

func (c *Client) GroupsEnableOnline(param) (response) {
	return methods.GroupsEnableOnline(c.Client(), param)
}

*/

func (c *Client) GroupsGet(userId int, filter, fields []string) (*methods.GroupsGetResponse, error) {
	return methods.GroupsGet(c.Client(), userId, filter, fields, MAX_COUNT_GROUPS)
}

func (c *Client) GroupsGetById(groupIds []int, fields []string) (*methods.GroupsGetByIdResponse, error) {
	return methods.GroupsGetById(c.Client(), groupIds, fields, MAX_COUNT_GROUPS)
}

/*

func (c *Client) GroupsGetAddresses(param) (response) {
	return methods.GroupsGetAddresses(c.Client(), param)
}

func (c *Client) GroupsGetBanned(param) (response) {
	return methods.GroupsGetBanned(c.Client(), param)
}

func (c *Client) GroupsGetById(param) (response) {
	return methods.GroupsGetById(c.Client(), param)
}

func (c *Client) GroupsGetCallbackConfirmationCode(param) (response) {
	return methods.GroupsGetCallbackConfirmationCode(c.Client(), param)
}

func (c *Client) GroupsGetCallbackServers(param) (response) {
	return methods.GroupsGetCallbackServers(c.Client(), param)
}

func (c *Client) GroupsGetCallbackSettings(param) (response) {
	return methods.GroupsGetCallbackSettings(c.Client(), param)
}

func (c *Client) GroupsGetCatalog(param) (response) {
	return methods.GroupsGetCatalog(c.Client(), param)
}

func (c *Client) GroupsGetCatalogInfo(param) (response) {
	return methods.GroupsGetCatalogInfo(c.Client(), param)
}

func (c *Client) GroupsGetInvitedUsers(param) (response) {
	return methods.GroupsGetInvitedUsers(c.Client(), param)
}

func (c *Client) GroupsGetInvites(param) (response) {
	return methods.GroupsGetInvites(c.Client(), param)
}

*/

func (c *Client) GroupsGetLongPollServer(groupId int) (*methods.GroupsGetLongPollServerResponse, error) {
	return methods.GroupsGetLongPollServer(c.Client(), groupId)
}

/*

func (c *Client) GroupsGetLongPollSettings(param) (response) {
	return methods.GroupsGetLongPollSettings(c.Client(), param)
}

func (c *Client) GroupsGetMembers(param) (response) {
	return methods.GroupsGetMembers(c.Client(), param)
}

func (c *Client) GroupsGetOnlineStatus(param) (response) {
	return methods.GroupsGetOnlineStatus(c.Client(), param)
}

func (c *Client) GroupsGetRequests(param) (response) {
	return methods.GroupsGetRequests(c.Client(), param)
}

func (c *Client) GroupsGetSettings(param) (response) {
	return methods.GroupsGetSettings(c.Client(), param)
}

*/

func (c *Client) GroupsGetTokenPermissions() (types.CommunityPermissions, error) {
	return methods.GroupsGetTokenPermissions(c.Client())
}

/*

func (c *Client) GroupsInvite(param) (response) {
	return methods.GroupsInvite(c.Client(), param)
}

func (c *Client) GroupsIsMember(param) (response) {
	return methods.GroupsIsMember(c.Client(), param)
}

func (c *Client) GroupsJoin(param) (response) {
	return methods.GroupsJoin(c.Client(), param)
}

func (c *Client) GroupsLeave(param) (response) {
	return methods.GroupsLeave(c.Client(), param)
}

func (c *Client) GroupsRemoveUser(param) (response) {
	return methods.GroupsRemoveUser(c.Client(), param)
}

func (c *Client) GroupsReorderLink(param) (response) {
	return methods.GroupsReorderLink(c.Client(), param)
}

func (c *Client) GroupsSearch(param) (response) {
	return methods.GroupsSearch(c.Client(), param)
}

func (c *Client) GroupsSetCallbackSettings(param) (response) {
	return methods.GroupsSetCallbackSettings(c.Client(), param)
}

func (c *Client) GroupsSetLongPollSettings(param) (response) {
	return methods.GroupsSetLongPollSettings(c.Client(), param)
}

func (c *Client) GroupsSetSettings(param) (response) {
	return methods.GroupsSetSettings(c.Client(), param)
}

func (c *Client) GroupsUnban(param) (response) {
	return methods.GroupsUnban(c.Client(), param)
}

func (c *Client) LeadformsCreate(param) (response) {
	return methods.LeadformsCreate(c.Client(), param)
}

func (c *Client) LeadformsDelete(param) (response) {
	return methods.LeadformsDelete(c.Client(), param)
}

func (c *Client) LeadformsGet(param) (response) {
	return methods.LeadformsGet(c.Client(), param)
}

func (c *Client) LeadformsGetLeads(param) (response) {
	return methods.LeadformsGetLeads(c.Client(), param)
}

func (c *Client) LeadformsGetUploadURL(param) (response) {
	return methods.LeadformsGetUploadURL(c.Client(), param)
}

func (c *Client) LeadformsList(param) (response) {
	return methods.LeadformsList(c.Client(), param)
}

func (c *Client) LeadformsUpdate(param) (response) {
	return methods.LeadformsUpdate(c.Client(), param)
}

func (c *Client) LeadsCheckUser(param) (response) {
	return methods.LeadsCheckUser(c.Client(), param)
}

func (c *Client) LeadsComplete(param) (response) {
	return methods.LeadsComplete(c.Client(), param)
}

func (c *Client) LeadsGetStats(param) (response) {
	return methods.LeadsGetStats(c.Client(), param)
}

func (c *Client) LeadsGetUsers(param) (response) {
	return methods.LeadsGetUsers(c.Client(), param)
}

func (c *Client) LeadsMetricHit(param) (response) {
	return methods.LeadsMetricHit(c.Client(), param)
}

func (c *Client) LeadsStart(param) (response) {
	return methods.LeadsStart(c.Client(), param)
}

func (c *Client) LikesAdd(param) (response) {
	return methods.LikesAdd(c.Client(), param)
}

func (c *Client) LikesDelete(param) (response) {
	return methods.LikesDelete(c.Client(), param)
}

func (c *Client) LikesGetList(param) (response) {
	return methods.LikesGetList(c.Client(), param)
}

func (c *Client) LikesIsLiked(param) (response) {
	return methods.LikesIsLiked(c.Client(), param)
}

func (c *Client) MarketAdd(param) (response) {
	return methods.MarketAdd(c.Client(), param)
}

func (c *Client) MarketAddAlbum(param) (response) {
	return methods.MarketAddAlbum(c.Client(), param)
}

func (c *Client) MarketAddToAlbum(param) (response) {
	return methods.MarketAddToAlbum(c.Client(), param)
}

func (c *Client) MarketCreateComment(param) (response) {
	return methods.MarketCreateComment(c.Client(), param)
}

func (c *Client) MarketDelete(param) (response) {
	return methods.MarketDelete(c.Client(), param)
}

func (c *Client) MarketDeleteAlbum(param) (response) {
	return methods.MarketDeleteAlbum(c.Client(), param)
}

func (c *Client) MarketDeleteComment(param) (response) {
	return methods.MarketDeleteComment(c.Client(), param)
}

func (c *Client) MarketEdit(param) (response) {
	return methods.MarketEdit(c.Client(), param)
}

func (c *Client) MarketEditAlbum(param) (response) {
	return methods.MarketEditAlbum(c.Client(), param)
}

func (c *Client) MarketEditComment(param) (response) {
	return methods.MarketEditComment(c.Client(), param)
}

func (c *Client) MarketGet(param) (response) {
	return methods.MarketGet(c.Client(), param)
}

func (c *Client) MarketGetAlbumById(param) (response) {
	return methods.MarketGetAlbumById(c.Client(), param)
}

func (c *Client) MarketGetAlbums(param) (response) {
	return methods.MarketGetAlbums(c.Client(), param)
}

func (c *Client) MarketGetById(param) (response) {
	return methods.MarketGetById(c.Client(), param)
}

func (c *Client) MarketGetCategories(param) (response) {
	return methods.MarketGetCategories(c.Client(), param)
}

func (c *Client) MarketGetComments(param) (response) {
	return methods.MarketGetComments(c.Client(), param)
}

func (c *Client) MarketRemoveFromAlbum(param) (response) {
	return methods.MarketRemoveFromAlbum(c.Client(), param)
}

func (c *Client) MarketReorderAlbums(param) (response) {
	return methods.MarketReorderAlbums(c.Client(), param)
}

func (c *Client) MarketReorderItems(param) (response) {
	return methods.MarketReorderItems(c.Client(), param)
}

func (c *Client) MarketReport(param) (response) {
	return methods.MarketReport(c.Client(), param)
}

func (c *Client) MarketReportComment(param) (response) {
	return methods.MarketReportComment(c.Client(), param)
}

func (c *Client) MarketRestore(param) (response) {
	return methods.MarketRestore(c.Client(), param)
}

func (c *Client) MarketRestoreComment(param) (response) {
	return methods.MarketRestoreComment(c.Client(), param)
}

func (c *Client) MarketSearch(param) (response) {
	return methods.MarketSearch(c.Client(), param)
}

func (c *Client) MessagesAddChatUser(param) (response) {
	return methods.MessagesAddChatUser(c.Client(), param)
}

func (c *Client) MessagesAllowMessagesFromGroup(param) (response) {
	return methods.MessagesAllowMessagesFromGroup(c.Client(), param)
}

func (c *Client) MessagesCreateChat(param) (response) {
	return methods.MessagesCreateChat(c.Client(), param)
}

func (c *Client) MessagesDelete(param) (response) {
	return methods.MessagesDelete(c.Client(), param)
}

func (c *Client) MessagesDeleteChatPhoto(param) (response) {
	return methods.MessagesDeleteChatPhoto(c.Client(), param)
}

func (c *Client) MessagesDeleteConversation(param) (response) {
	return methods.MessagesDeleteConversation(c.Client(), param)
}

func (c *Client) MessagesDenyMessagesFromGroup(param) (response) {
	return methods.MessagesDenyMessagesFromGroup(c.Client(), param)
}

func (c *Client) MessagesEdit(param) (response) {
	return methods.MessagesEdit(c.Client(), param)
}

func (c *Client) MessagesEditChat(param) (response) {
	return methods.MessagesEditChat(c.Client(), param)
}

func (c *Client) MessagesGetByConversationMessageId(param) (response) {
	return methods.MessagesGetByConversationMessageId(c.Client(), param)
}

*/

func (c *Client) MessagesGetById(messageIds []int, previewLen, groupId int) (*methods.MessagesGetByIdResponse, error) {
	return methods.MessagesGetById(c, messageIds, previewLen, groupId)
}

/*

func (c *Client) MessagesGetChat(param) (response) {
	return methods.MessagesGetChat(c.Client(), param)
}

func (c *Client) MessagesGetChatPreview(param) (response) {
	return methods.MessagesGetChatPreview(c.Client(), param)
}

func (c *Client) MessagesGetConversationMembers(param) (response) {
	return methods.MessagesGetConversationMembers(c.Client(), param)
}

*/

func (c *Client) MessagesGetConversations(r methods.MessagesGetConversationsRequest) (*methods.MessagesGetConversationsResponse, error) {
	return methods.MessagesGetConversations(c.Client(), r)
}

/*

func (c *Client) MessagesGetConversationsById(param) (response) {
	return methods.MessagesGetConversationsById(c.Client(), param)
}

func (c *Client) MessagesGetHistory(param) (response) {
	return methods.MessagesGetHistory(c.Client(), param)
}

func (c *Client) MessagesGetHistoryAttachments(param) (response) {
	return methods.MessagesGetHistoryAttachments(c.Client(), param)
}

func (c *Client) MessagesGetImportantMessages(param) (response) {
	return methods.MessagesGetImportantMessages(c.Client(), param)
}

func (c *Client) MessagesGetInviteLink(param) (response) {
	return methods.MessagesGetInviteLink(c.Client(), param)
}

func (c *Client) MessagesGetLastActivity(param) (response) {
	return methods.MessagesGetLastActivity(c.Client(), param)
}

func (c *Client) MessagesGetLongPollHistory(param) (response) {
	return methods.MessagesGetLongPollHistory(c.Client(), param)
}

*/

func (c *Client) MessagesGetLongPollServer(needPts bool, groupID int) (*methods.MessagesGetLongPollServerResponse, error) {
	return methods.MessagesGetLongPollServer(c.Client(), needPts, groupID)
}

/*

func (c *Client) MessagesIsMessagesFromGroupAllowed(param) (response) {
	return methods.MessagesIsMessagesFromGroupAllowed(c.Client(), param)
}

func (c *Client) MessagesJoinChatByInviteLink(param) (response) {
	return methods.MessagesJoinChatByInviteLink(c.Client(), param)
}

func (c *Client) MessagesMarkAsAnsweredConversation(param) (response) {
	return methods.MessagesMarkAsAnsweredConversation(c.Client(), param)
}

func (c *Client) MessagesMarkAsImportant(param) (response) {
	return methods.MessagesMarkAsImportant(c.Client(), param)
}

func (c *Client) MessagesMarkAsImportantConversation(param) (response) {
	return methods.MessagesMarkAsImportantConversation(c.Client(), param)
}

func (c *Client) MessagesMarkAsRead(param) (response) {
	return methods.MessagesMarkAsRead(c.Client(), param)
}

func (c *Client) MessagesPin(param) (response) {
	return methods.MessagesPin(c.Client(), param)
}

func (c *Client) MessagesRemoveChatUser(param) (response) {
	return methods.MessagesRemoveChatUser(c.Client(), param)
}

func (c *Client) MessagesRestore(param) (response) {
	return methods.MessagesRestore(c.Client(), param)
}

func (c *Client) MessagesSearch(param) (response) {
	return methods.MessagesSearch(c.Client(), param)
}

func (c *Client) MessagesSearchConversations(param) (response) {
	return methods.MessagesSearchConversations(c.Client(), param)
}

*/

func (c *Client) MessagesSend(m methods.MessagesSendRequest) error {
	return methods.MessagesSend(c.Client(), m)
}

/*

func (c *Client) MessagesSetActivity(param) (response) {
	return methods.MessagesSetActivity(c.Client(), param)
}

func (c *Client) MessagesSetChatPhoto(param) (response) {
	return methods.MessagesSetChatPhoto(c.Client(), param)
}

func (c *Client) MessagesUnpin(param) (response) {
	return methods.MessagesUnpin(c.Client(), param)
}

func (c *Client) NewsfeedAddBan(param) (response) {
	return methods.NewsfeedAddBan(c.Client(), param)
}

func (c *Client) NewsfeedDeleteBan(param) (response) {
	return methods.NewsfeedDeleteBan(c.Client(), param)
}

func (c *Client) NewsfeedDeleteList(param) (response) {
	return methods.NewsfeedDeleteList(c.Client(), param)
}

func (c *Client) NewsfeedGet(param) (response) {
	return methods.NewsfeedGet(c.Client(), param)
}

func (c *Client) NewsfeedGetBanned(param) (response) {
	return methods.NewsfeedGetBanned(c.Client(), param)
}

func (c *Client) NewsfeedGetComments(param) (response) {
	return methods.NewsfeedGetComments(c.Client(), param)
}

func (c *Client) NewsfeedGetLists(param) (response) {
	return methods.NewsfeedGetLists(c.Client(), param)
}

func (c *Client) NewsfeedGetMentions(param) (response) {
	return methods.NewsfeedGetMentions(c.Client(), param)
}

func (c *Client) NewsfeedGetRecommended(param) (response) {
	return methods.NewsfeedGetRecommended(c.Client(), param)
}

func (c *Client) NewsfeedGetSuggestedSources(param) (response) {
	return methods.NewsfeedGetSuggestedSources(c.Client(), param)
}

func (c *Client) NewsfeedIgnoreItem(param) (response) {
	return methods.NewsfeedIgnoreItem(c.Client(), param)
}

func (c *Client) NewsfeedSaveList(param) (response) {
	return methods.NewsfeedSaveList(c.Client(), param)
}

func (c *Client) NewsfeedSearch(param) (response) {
	return methods.NewsfeedSearch(c.Client(), param)
}

func (c *Client) NewsfeedUnignoreItem(param) (response) {
	return methods.NewsfeedUnignoreItem(c.Client(), param)
}

func (c *Client) NewsfeedUnsubscribe(param) (response) {
	return methods.NewsfeedUnsubscribe(c.Client(), param)
}

func (c *Client) NotesAdd(param) (response) {
	return methods.NotesAdd(c.Client(), param)
}

func (c *Client) NotesCreateComment(param) (response) {
	return methods.NotesCreateComment(c.Client(), param)
}

func (c *Client) NotesDelete(param) (response) {
	return methods.NotesDelete(c.Client(), param)
}

func (c *Client) NotesDeleteComment(param) (response) {
	return methods.NotesDeleteComment(c.Client(), param)
}

func (c *Client) NotesEdit(param) (response) {
	return methods.NotesEdit(c.Client(), param)
}

func (c *Client) NotesEditComment(param) (response) {
	return methods.NotesEditComment(c.Client(), param)
}

func (c *Client) NotesGet(param) (response) {
	return methods.NotesGet(c.Client(), param)
}

func (c *Client) NotesGetById(param) (response) {
	return methods.NotesGetById(c.Client(), param)
}

func (c *Client) NotesGetComments(param) (response) {
	return methods.NotesGetComments(c.Client(), param)
}

func (c *Client) NotesRestoreComment(param) (response) {
	return methods.NotesRestoreComment(c.Client(), param)
}

func (c *Client) NotificationsGet(param) (response) {
	return methods.NotificationsGet(c.Client(), param)
}

func (c *Client) NotificationsMarkAsViewed(param) (response) {
	return methods.NotificationsMarkAsViewed(c.Client(), param)
}

func (c *Client) NotificationsSendMessage(param) (response) {
	return methods.NotificationsSendMessage(c.Client(), param)
}

func (c *Client) OrdersCancelSubscription(param) (response) {
	return methods.OrdersCancelSubscription(c.Client(), param)
}

func (c *Client) OrdersChangeState(param) (response) {
	return methods.OrdersChangeState(c.Client(), param)
}

func (c *Client) OrdersGet(param) (response) {
	return methods.OrdersGet(c.Client(), param)
}

func (c *Client) OrdersGetAmount(param) (response) {
	return methods.OrdersGetAmount(c.Client(), param)
}

func (c *Client) OrdersGetById(param) (response) {
	return methods.OrdersGetById(c.Client(), param)
}

func (c *Client) OrdersGetUserSubscriptionById(param) (response) {
	return methods.OrdersGetUserSubscriptionById(c.Client(), param)
}

func (c *Client) OrdersGetUserSubscriptions(param) (response) {
	return methods.OrdersGetUserSubscriptions(c.Client(), param)
}

func (c *Client) OrdersUpdateSubscription(param) (response) {
	return methods.OrdersUpdateSubscription(c.Client(), param)
}

func (c *Client) PagesClearCache(param) (response) {
	return methods.PagesClearCache(c.Client(), param)
}

func (c *Client) PagesGet(param) (response) {
	return methods.PagesGet(c.Client(), param)
}

func (c *Client) PagesGetHistory(param) (response) {
	return methods.PagesGetHistory(c.Client(), param)
}

func (c *Client) PagesGetTitles(param) (response) {
	return methods.PagesGetTitles(c.Client(), param)
}

func (c *Client) PagesGetVersion(param) (response) {
	return methods.PagesGetVersion(c.Client(), param)
}

func (c *Client) PagesParseWiki(param) (response) {
	return methods.PagesParseWiki(c.Client(), param)
}

func (c *Client) PagesSave(param) (response) {
	return methods.PagesSave(c.Client(), param)
}

func (c *Client) PagesSaveAccess(param) (response) {
	return methods.PagesSaveAccess(c.Client(), param)
}

func (c *Client) PhotoConfirmTag(param) (response) {
	return methods.PhotoConfirmTag(c.Client(), param)
}

func (c *Client) PhotoCopy(param) (response) {
	return methods.PhotoCopy(c.Client(), param)
}

func (c *Client) PhotoCreateAlbum(param) (response) {
	return methods.PhotoCreateAlbum(c.Client(), param)
}

func (c *Client) PhotoCreateComment(param) (response) {
	return methods.PhotoCreateComment(c.Client(), param)
}

func (c *Client) PhotoDelete(param) (response) {
	return methods.PhotoDelete(c.Client(), param)
}

func (c *Client) PhotoDeleteAlbum(param) (response) {
	return methods.PhotoDeleteAlbum(c.Client(), param)
}

func (c *Client) PhotoDeleteComment(param) (response) {
	return methods.PhotoDeleteComment(c.Client(), param)
}

func (c *Client) PhotoEdit(param) (response) {
	return methods.PhotoEdit(c.Client(), param)
}

func (c *Client) PhotoEditAlbum(param) (response) {
	return methods.PhotoEditAlbum(c.Client(), param)
}

func (c *Client) PhotoEditComment(param) (response) {
	return methods.PhotoEditComment(c.Client(), param)
}

func (c *Client) PhotoGet(param) (response) {
	return methods.PhotoGet(c.Client(), param)
}

func (c *Client) PhotoGetAlbums(param) (response) {
	return methods.PhotoGetAlbums(c.Client(), param)
}

func (c *Client) PhotoGetAlbumsCount(param) (response) {
	return methods.PhotoGetAlbumsCount(c.Client(), param)
}

func (c *Client) PhotoGetAll(param) (response) {
	return methods.PhotoGetAll(c.Client(), param)
}

func (c *Client) PhotoGetAllComments(param) (response) {
	return methods.PhotoGetAllComments(c.Client(), param)
}

func (c *Client) PhotoGetById(param) (response) {
	return methods.PhotoGetById(c.Client(), param)
}

func (c *Client) PhotoGetChatUploadServer(param) (response) {
	return methods.PhotoGetChatUploadServer(c.Client(), param)
}

func (c *Client) PhotoGetComments(param) (response) {
	return methods.PhotoGetComments(c.Client(), param)
}

func (c *Client) PhotoGetMarketAlbumUploadServer(param) (response) {
	return methods.PhotoGetMarketAlbumUploadServer(c.Client(), param)
}

func (c *Client) PhotoGetMarketUploadServer(param) (response) {
	return methods.PhotoGetMarketUploadServer(c.Client(), param)
}

func (c *Client) PhotoGetMessagesUploadServer(param) (response) {
	return methods.PhotoGetMessagesUploadServer(c.Client(), param)
}

func (c *Client) PhotoGetNewTags(param) (response) {
	return methods.PhotoGetNewTags(c.Client(), param)
}

func (c *Client) PhotoGetOwnerCoverPhotoUploadServer(param) (response) {
	return methods.PhotoGetOwnerCoverPhotoUploadServer(c.Client(), param)
}

func (c *Client) PhotoGetOwnerPhotoUploadServer(param) (response) {
	return methods.PhotoGetOwnerPhotoUploadServer(c.Client(), param)
}

func (c *Client) PhotoGetTags(param) (response) {
	return methods.PhotoGetTags(c.Client(), param)
}

func (c *Client) PhotoGetUploadServer(param) (response) {
	return methods.PhotoGetUploadServer(c.Client(), param)
}

func (c *Client) PhotoGetUserPhotos(param) (response) {
	return methods.PhotoGetUserPhotos(c.Client(), param)
}

func (c *Client) PhotoGetWallUploadServer(param) (response) {
	return methods.PhotoGetWallUploadServer(c.Client(), param)
}

func (c *Client) PhotoMakeCover(param) (response) {
	return methods.PhotoMakeCover(c.Client(), param)
}

func (c *Client) PhotoMove(param) (response) {
	return methods.PhotoMove(c.Client(), param)
}

func (c *Client) PhotoPutTag(param) (response) {
	return methods.PhotoPutTag(c.Client(), param)
}

func (c *Client) PhotoRemoveTag(param) (response) {
	return methods.PhotoRemoveTag(c.Client(), param)
}

func (c *Client) PhotoReorderAlbums(param) (response) {
	return methods.PhotoReorderAlbums(c.Client(), param)
}

func (c *Client) PhotoReorderPhotos(param) (response) {
	return methods.PhotoReorderPhotos(c.Client(), param)
}

func (c *Client) PhotoReport(param) (response) {
	return methods.PhotoReport(c.Client(), param)
}

func (c *Client) PhotoReportComment(param) (response) {
	return methods.PhotoReportComment(c.Client(), param)
}

func (c *Client) PhotoRestore(param) (response) {
	return methods.PhotoRestore(c.Client(), param)
}

func (c *Client) PhotoRestoreComment(param) (response) {
	return methods.PhotoRestoreComment(c.Client(), param)
}

func (c *Client) PhotoSave(param) (response) {
	return methods.PhotoSave(c.Client(), param)
}

func (c *Client) PhotoSaveMarketAlbumPhoto(param) (response) {
	return methods.PhotoSaveMarketAlbumPhoto(c.Client(), param)
}

func (c *Client) PhotoSaveMarketPhoto(param) (response) {
	return methods.PhotoSaveMarketPhoto(c.Client(), param)
}

func (c *Client) PhotoSaveMessagesPhoto(param) (response) {
	return methods.PhotoSaveMessagesPhoto(c.Client(), param)
}

func (c *Client) PhotoSaveOwnerCoverPhoto(param) (response) {
	return methods.PhotoSaveOwnerCoverPhoto(c.Client(), param)
}

func (c *Client) PhotoSaveOwnerPhoto(param) (response) {
	return methods.PhotoSaveOwnerPhoto(c.Client(), param)
}

func (c *Client) PhotoSaveWallPhoto(param) (response) {
	return methods.PhotoSaveWallPhoto(c.Client(), param)
}

func (c *Client) PhotoSearch(param) (response) {
	return methods.PhotoSearch(c.Client(), param)
}

func (c *Client) PollsAddVote(param) (response) {
	return methods.PollsAddVote(c.Client(), param)
}

func (c *Client) PollsCreate(param) (response) {
	return methods.PollsCreate(c.Client(), param)
}

func (c *Client) PollsDeleteVote(param) (response) {
	return methods.PollsDeleteVote(c.Client(), param)
}

func (c *Client) PollsEdit(param) (response) {
	return methods.PollsEdit(c.Client(), param)
}

func (c *Client) PollsGetBackgrounds(param) (response) {
	return methods.PollsGetBackgrounds(c.Client(), param)
}

func (c *Client) PollsGetById(param) (response) {
	return methods.PollsGetById(c.Client(), param)
}

func (c *Client) PollsGetPhotoUploadServer(param) (response) {
	return methods.PollsGetPhotoUploadServer(c.Client(), param)
}

func (c *Client) PollsGetVoters(param) (response) {
	return methods.PollsGetVoters(c.Client(), param)
}

func (c *Client) PollsSavePhoto(param) (response) {
	return methods.PollsSavePhoto(c.Client(), param)
}

func (c *Client) PrettycardsCreate(param) (response) {
	return methods.PrettycardsCreate(c.Client(), param)
}

func (c *Client) PrettycardsDelete(param) (response) {
	return methods.PrettycardsDelete(c.Client(), param)
}

func (c *Client) PrettycardsEdit(param) (response) {
	return methods.PrettycardsEdit(c.Client(), param)
}

func (c *Client) PrettycardsGet(param) (response) {
	return methods.PrettycardsGet(c.Client(), param)
}

func (c *Client) PrettycardsGetById(param) (response) {
	return methods.PrettycardsGetById(c.Client(), param)
}

func (c *Client) PrettycardsGetUploadURL(param) (response) {
	return methods.PrettycardsGetUploadURL(c.Client(), param)
}

func (c *Client) SearchGetHints(param) (response) {
	return methods.SearchGetHints(c.Client(), param)
}

func (c *Client) SecureAddAppEvent(param) (response) {
	return methods.SecureAddAppEvent(c.Client(), param)
}

func (c *Client) SecureCheckToken(param) (response) {
	return methods.SecureCheckToken(c.Client(), param)
}

func (c *Client) SecureGetAppBalance(param) (response) {
	return methods.SecureGetAppBalance(c.Client(), param)
}

func (c *Client) SecureGetSMSHistory(param) (response) {
	return methods.SecureGetSMSHistory(c.Client(), param)
}

func (c *Client) SecureGetTransactionsHistory(param) (response) {
	return methods.SecureGetTransactionsHistory(c.Client(), param)
}

func (c *Client) SecureGetUserLevel(param) (response) {
	return methods.SecureGetUserLevel(c.Client(), param)
}

func (c *Client) SecureGiveEventSticker(param) (response) {
	return methods.SecureGiveEventSticker(c.Client(), param)
}

func (c *Client) SecureSendNotification(param) (response) {
	return methods.SecureSendNotification(c.Client(), param)
}

func (c *Client) SecureSendSMSNotification(param) (response) {
	return methods.SecureSendSMSNotification(c.Client(), param)
}

func (c *Client) SecureSetCounter(param) (response) {
	return methods.SecureSetCounter(c.Client(), param)
}

func (c *Client) StatsGet(param) (response) {
	return methods.StatsGet(c.Client(), param)
}

func (c *Client) StatsGetPostReach(param) (response) {
	return methods.StatsGetPostReach(c.Client(), param)
}

func (c *Client) StatsTrackVisitor(param) (response) {
	return methods.StatsTrackVisitor(c.Client(), param)
}

func (c *Client) StatusGet(param) (response) {
	return methods.StatusGet(c.Client(), param)
}

func (c *Client) StatusSet(param) (response) {
	return methods.StatusSet(c.Client(), param)
}

*/

func (c *Client) StorageGet(keys []string, userId int) (map[string]string, error) {
	return methods.StorageGet(c.Client(), keys, userId)
}

/*

func (c *Client) StorageGetKeys(param) (response) {
	return methods.StorageGetKeys(c.Client(), param)
}

*/

func (c *Client) StorageSet(key string, value interface{}, userId int) error {
	return methods.StorageSet(c.Client(), key, value, userId)
}

/*

func (c *Client) StoriesBanOwner(param) (response) {
	return methods.StoriesBanOwner(c.Client(), param)
}

func (c *Client) StoriesDelete(param) (response) {
	return methods.StoriesDelete(c.Client(), param)
}

func (c *Client) StoriesGet(param) (response) {
	return methods.StoriesGet(c.Client(), param)
}

func (c *Client) StoriesGetBanned(param) (response) {
	return methods.StoriesGetBanned(c.Client(), param)
}

func (c *Client) StoriesGetById(param) (response) {
	return methods.StoriesGetById(c.Client(), param)
}

func (c *Client) StoriesGetPhotoUploadServer(param) (response) {
	return methods.StoriesGetPhotoUploadServer(c.Client(), param)
}

func (c *Client) StoriesGetReplies(param) (response) {
	return methods.StoriesGetReplies(c.Client(), param)
}

func (c *Client) StoriesGetStats(param) (response) {
	return methods.StoriesGetStats(c.Client(), param)
}

func (c *Client) StoriesGetVideoUploadServer(param) (response) {
	return methods.StoriesGetVideoUploadServer(c.Client(), param)
}

func (c *Client) StoriesGetViewers(param) (response) {
	return methods.StoriesGetViewers(c.Client(), param)
}

func (c *Client) StoriesHideAllReplies(param) (response) {
	return methods.StoriesHideAllReplies(c.Client(), param)
}

func (c *Client) StoriesHideReply(param) (response) {
	return methods.StoriesHideReply(c.Client(), param)
}

func (c *Client) StoriesSearch(param) (response) {
	return methods.StoriesSearch(c.Client(), param)
}

func (c *Client) StoriesUnbanOwner(param) (response) {
	return methods.StoriesUnbanOwner(c.Client(), param)
}

*/

func (c *Client) StreamingGetServerUrl() (*methods.StreamingGetServerUrlResponse, error) {
	return methods.StreamingGetServerUrl(c.Client())
}

/*

func (c *Client) StreamingGetSettings(param) (response) {
	return methods.StreamingGetSettings(c.Client(), param)
}

func (c *Client) StreamingGetStats(param) (response) {
	return methods.StreamingGetStats(c.Client(), param)
}

func (c *Client) StreamingGetStem(param) (response) {
	return methods.StreamingGetStem(c.Client(), param)
}

func (c *Client) StreamingSetSettings(param) (response) {
	return methods.StreamingSetSettings(c.Client(), param)
}

*/

func (c *Client) UsersGet(userIds []int, fields []string) (*methods.UsersGetResponse, error) {
	return methods.UsersGet(c.Client(), userIds, fields)
}

func (c *Client) UsersGetRegDate(userId int) (*methods.UsersGetRegDateResponse, error) {
	return methods.UsersGetRegDate(c.Client(), userId)
}

/*

func (c *Client) UsersGetFollowers(param) (response) {
	return methods.UsersGetFollowers(c.Client(), param)
}

func (c *Client) UsersGetSubscriptions(param) (response) {
	return methods.UsersGetSubscriptions(c.Client(), param)
}

func (c *Client) UsersReport(param) (response) {
	return methods.UsersReport(c.Client(), param)
}

func (c *Client) UsersSearch(param) (response) {
	return methods.UsersSearch(c.Client(), param)
}

func (c *Client) UtilsCheckLink(param) (response) {
	return methods.UtilsCheckLink(c.Client(), param)
}

func (c *Client) UtilsDeleteFromLastShortened(param) (response) {
	return methods.UtilsDeleteFromLastShortened(c.Client(), param)
}

func (c *Client) UtilsGetLastShortenedLinks(param) (response) {
	return methods.UtilsGetLastShortenedLinks(c.Client(), param)
}

func (c *Client) UtilsGetLinkStats(param) (response) {
	return methods.UtilsGetLinkStats(c.Client(), param)
}

func (c *Client) UtilsGetServerTime(param) (response) {
	return methods.UtilsGetServerTime(c.Client(), param)
}

func (c *Client) UtilsGetShortLink(param) (response) {
	return methods.UtilsGetShortLink(c.Client(), param)
}

func (c *Client) UtilsResolveScreenName(param) (response) {
	return methods.UtilsResolveScreenName(c.Client(), param)
}

func (c *Client) VideoAdd(param) (response) {
	return methods.VideoAdd(c.Client(), param)
}

func (c *Client) VideoAddAlbum(param) (response) {
	return methods.VideoAddAlbum(c.Client(), param)
}

func (c *Client) VideoAddToAlbum(param) (response) {
	return methods.VideoAddToAlbum(c.Client(), param)
}

func (c *Client) VideoCreateComment(param) (response) {
	return methods.VideoCreateComment(c.Client(), param)
}

func (c *Client) VideoDelete(param) (response) {
	return methods.VideoDelete(c.Client(), param)
}

func (c *Client) VideoDeleteAlbum(param) (response) {
	return methods.VideoDeleteAlbum(c.Client(), param)
}

func (c *Client) VideoDeleteComment(param) (response) {
	return methods.VideoDeleteComment(c.Client(), param)
}

func (c *Client) VideoEdit(param) (response) {
	return methods.VideoEdit(c.Client(), param)
}

func (c *Client) VideoEditAlbum(param) (response) {
	return methods.VideoEditAlbum(c.Client(), param)
}

func (c *Client) VideoEditComment(param) (response) {
	return methods.VideoEditComment(c.Client(), param)
}

func (c *Client) VideoGet(param) (response) {
	return methods.VideoGet(c.Client(), param)
}

func (c *Client) VideoGetAlbumById(param) (response) {
	return methods.VideoGetAlbumById(c.Client(), param)
}

func (c *Client) VideoGetAlbums(param) (response) {
	return methods.VideoGetAlbums(c.Client(), param)
}

func (c *Client) VideoGetAlbumsByVideo(param) (response) {
	return methods.VideoGetAlbumsByVideo(c.Client(), param)
}

func (c *Client) VideoGetComments(param) (response) {
	return methods.VideoGetComments(c.Client(), param)
}

func (c *Client) VideoRemoveFromAlbum(param) (response) {
	return methods.VideoRemoveFromAlbum(c.Client(), param)
}

func (c *Client) VideoReorderAlbums(param) (response) {
	return methods.VideoReorderAlbums(c.Client(), param)
}

func (c *Client) VideoReorderVideos(param) (response) {
	return methods.VideoReorderVideos(c.Client(), param)
}

func (c *Client) VideoReport(param) (response) {
	return methods.VideoReport(c.Client(), param)
}

func (c *Client) VideoReportComment(param) (response) {
	return methods.VideoReportComment(c.Client(), param)
}

func (c *Client) VideoRestore(param) (response) {
	return methods.VideoRestore(c.Client(), param)
}

func (c *Client) VideoRestoreComment(param) (response) {
	return methods.VideoRestoreComment(c.Client(), param)
}

func (c *Client) VideoSave(param) (response) {
	return methods.VideoSave(c.Client(), param)
}

func (c *Client) VideoSearch(param) (response) {
	return methods.VideoSearch(c.Client(), param)
}

func (c *Client) WallCloseComments(param) (response) {
	return methods.WallCloseComments(c.Client(), param)
}

func (c *Client) WallCreateComment(param) (response) {
	return methods.WallCreateComment(c.Client(), param)
}

func (c *Client) WallDelete(param) (response) {
	return methods.WallDelete(c.Client(), param)
}

func (c *Client) WallDeleteComment(param) (response) {
	return methods.WallDeleteComment(c.Client(), param)
}

func (c *Client) WallEdit(param) (response) {
	return methods.WallEdit(c.Client(), param)
}

func (c *Client) WallEditAdsStealth(param) (response) {
	return methods.WallEditAdsStealth(c.Client(), param)
}

func (c *Client) WallEditComment(param) (response) {
	return methods.WallEditComment(c.Client(), param)
}

*/

func (c *Client) WallGet(Owner int, filter string, maxPosts int) (*methods.WallGetResponse, error) {
	return methods.WallGet(c.Client(), Owner, filter, MAX_COUNT_POSTS)
}

/*

func (c *Client) WallGetById(param) (response) {
	return methods.WallGetById(c.Client(), param)
}

func (c *Client) WallGetComment(param) (response) {
	return methods.WallGetComment(c.Client(), param)
}

func (c *Client) WallGetComments(param) (response) {
	return methods.WallGetComments(c.Client(), param)
}

func (c *Client) WallGetReposts(param) (response) {
	return methods.WallGetReposts(c.Client(), param)
}

func (c *Client) WallOpenComments(param) (response) {
	return methods.WallOpenComments(c.Client(), param)
}

func (c *Client) WallPin(param) (response) {
	return methods.WallPin(c.Client(), param)
}

func (c *Client) WallPost(param) (response) {
	return methods.WallPost(c.Client(), param)
}

func (c *Client) WallPostAdsStealth(param) (response) {
	return methods.WallPostAdsStealth(c.Client(), param)
}

func (c *Client) WallReportComment(param) (response) {
	return methods.WallReportComment(c.Client(), param)
}

func (c *Client) WallReportPost(param) (response) {
	return methods.WallReportPost(c.Client(), param)
}

func (c *Client) WallRepost(param) (response) {
	return methods.WallRepost(c.Client(), param)
}

func (c *Client) WallRestore(param) (response) {
	return methods.WallRestore(c.Client(), param)
}

func (c *Client) WallRestoreComment(param) (response) {
	return methods.WallRestoreComment(c.Client(), param)
}

func (c *Client) WallSearch(param) (response) {
	return methods.WallSearch(c.Client(), param)
}

func (c *Client) WallUnpin(param) (response) {
	return methods.WallUnpin(c.Client(), param)
}

func (c *Client) WidgetsGetComments(param) (response) {
	return methods.WidgetsGetComments(c.Client(), param)
}

func (c *Client) WidgetsGetPages(param) (response) {
	return methods.WidgetsGetPages(c.Client(), param)
}

*/
