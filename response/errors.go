package response

type ResponseError struct {
	Code   ErrorCode `json:"error_code"`
	Msg    string    `json:"error_msg"`
	Params []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"request_params"`
}

func (r *ResponseError) Error() string {
	return r.Msg
}

func (r *ResponseError) ErrCode() ErrorCode {
	return r.Code
}

type ErrorCode int

const (
	ErrUnknown                     = 1
	ErrAppDisabled                 = 2
	ErrMethodUnknown               = 3
	ErrInvalidSign                 = 4
	ErrAuthFailed                  = 5
	ErrTooManyRequests             = 6
	ErrPermissionDenied            = 7
	ErrInvalidRequest              = 8
	ErrSimilarRequests             = 9
	ErrInternal                    = 10
	ErrTestModeNeedsDisabled       = 11
	ErrUnableToCombileScript       = 12
	ErrUknownScriptException       = 13
	ErrCapthaRequired              = 14
	ErrAccessDenied                = 15
	ErrHttpsRequired               = 16
	ErrValidateAccount             = 17
	ErrUserHasBeenBlocked          = 18
	ErrContentIsUnavailable        = 19
	ErrStandaloneAppRequired       = 20 // TODO:
	ErrOpenAPIAppRequired          = 21 // по большей части названия ошибок неверны
	ErrMethodHasBeenDisabled       = 23 // так как в доках особо не указаны сокращения,
	ErrConfirmationRequired        = 24 // описание ошибок размыто и иногда дублируется
	ErrGroupKeyIsInvalid           = 27
	ErrServiceKeyIsInvalid         = 28
	ErrRequestLimit                = 29
	ErrProfileIsPrivate            = 30
	ErrNotImplementedYet           = 33
	ErrParametersAreInvalid        = 100
	ErrInvlaidApiId                = 101
	ErrToManyVariables             = 103
	ErrInvalidUserId               = 112
	ErrInvalidTimestamp            = 150
	ErrAlbumPermissionDenied       = 200
	ErrAudioPermissionDenied       = 201
	ErrGroupPermissionDenied       = 201
	ErrAlbumIsFull                 = 300
	ErrPaymentsDenied              = 500
	ErrAdsCabinetPermissionDenied  = 600
	ErrTooManyRequestsAdsOffice    = 601
	ErrAdsCabinetUnknown           = 603
	ErrUserIsBlacklisted           = 900
	ErrUserRejectCommunityMessages = 901
	ErrUserRejectPrivateMessages   = 902
	ErrKeyboardIsInvalid           = 911
	ErrEnableChatbotFeatures       = 912
	ErrTooManyForwardedMessages    = 913
	ErrMessageIsTooLong            = 914
	ErrWebsocketUpgradeExpected    = 1000
	ErrWrongHttpMethod             = 1001
	ErrWrongContentType            = 1002
	ErrMissingKey                  = 1003
	ErrBadKey                      = 1004
	ErrBadStreamId                 = 1005
	ErrConnectionAlreadyExist      = 1006
	ErrCantParseJson               = 2000
	ErrTagAlreadyExist             = 2001
	ErrTagNotExist                 = 2002
	ErrCantParseRule               = 2003
	ErrTooManyFilters              = 2004
	ErrIncorrectQuotes             = 2005
	ErrTooManyRules                = 2006
	ErrDontHavePositiveFilters     = 2008
)
