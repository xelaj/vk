package types

type User struct {
	Id                     int        `json:"id"`
	FirstName              string     `json:"first_name"`
	LastName               string     `json:"last_name"`
	Deactivated            string     `json:"deactivated"`
	IsClosed               bool       `json:"is_closed"`
	CanAccessClosed        bool       `json:"can_access_closed"`
	About                  string     `json:"about"`
	Activities             string     `json:"activities"`
	BirthDate              string     `json:"bdate"`
	Blacklisted            int        `json:"blacklisted"`
	BlacklistedByMe        int        `json:"blacklisted_by_me"`
	Books                  string     `json:"books"`
	CanPost                int        `json:"can_post"`
	CanSeeAllPosts         int        `json:"can_see_all_posts"`
	CanSeeAudio            int        `json:"can_see_audio"`
	CanSendFriendRequest   int        `json:"can_send_frien_request"`
	CanWritePrivateMessage int        `json:"can_write_private_message"`
	Career                 []struct{} `json:"career"`
	City                   struct{}   `json:"city"`
	CommonCount            int
	Connections            map[string]string
	Contacts               struct {
		MobilePhone string
		HomePhone   string
	} `json:"contacts"`
	Counters struct {
		Albums        int `json:"albums"`
		Videos        int `json:"videos"`
		Audios        int `json:"audios"`
		Photos        int `json:"photos"`
		Notes         int `json:"notes"`
		Friends       int `json:"friends"`
		Groups        int `json:"groups"`
		OnlineFriends int `json:"online_friends"`
		MutualFriends int `json:"mutual_friends"`
		UserVideos    int `json:"user_videos"`
		Followers     int `json:"followers"`
		Pages         int `json:"pages"`
	} `json:"counters"`
	Country struct {
		Id    int
		Title string
	} `json:"country"`
	CropPhoto        struct{}
	Domain           string
	Education        interface{} `json:"-"`
	Exports          interface{} `json:"-"`
	FirstNameNom     string      `json:"first_name_nom"`
	FirstNameGen     string      `json:"first_name_gen"`
	FirstNameDat     string      `json:"first_name_dat"`
	FirstNameAcc     string      `json:"first_name_acc"`
	FirstNameIns     string      `json:"first_name_ins"`
	FirstNameAbl     string      `json:"first_name_abl"`
	FollowersCount   int
	FriendStatus     int
	Games            string
	HasMobile        int
	HasPhoto         int
	HomeTown         string
	Interests        string
	IsFavorite       int
	IsFriend         int
	IsHiddenFromFeed int
	LastNameNom      string `json:"last_name_nom"`
	LastNameGen      string `json:"last_name_gen"`
	LastNameDat      string `json:"last_name_dat"`
	LastNameAcc      string `json:"last_name_acc"`
	LastNameIns      string `json:"last_name_ins"`
	LastNameAbl      string `json:"last_name_abl"`
	LastSeen         struct {
		Time     int `json:"time"`
		Platform int `json:"platform"`
	} `json:"last_seen"`
	//Lists      string `json:"lists"`
	//MaidenName string `json:"maiden_name
	//Military   string `json:"milirary
	Movies     string `json:"movies"`
	Music      string `json:"music"`
	Nickname   string `json:"nickname"`
	Occupation struct {
		Type string  `json:"type"`
		Id   float64 `json:"id"` // и здесь я охуел не меньше
		Name string  `json:"name"`
	} `json:"occupation"`
	Online       int `json:"online"`
	OnlineApp    int `json:"online_app"`
	OnlineMobile int `json:"online_mobile"`
	Personal     struct {
		Political  int      `json:"political"`
		Langs      []string `json:"langs"`
		Religion   string   `json:"religion"`
		InspiredBy string   `json:"inspired_by"`
		PeopleMain int      `json:"people_main"`
		LifeMain   int      `json:"life_main"`
		Smoking    int      `json:"smoking"`
		Alcohol    int      `json:"alcohol"`
	}
	Photo50      string `json:"photo_50"`
	Photo100     string `json:"photo_100"`
	Photo200Orig string `json:"photo_200_orig"`
	Photo200     string `json:"photo_200"`
	Photo400Orig string `json:"photo_400_orig"`
	PhotoId      string `json:"photo_id"`
	PhotoMax     string `json:"photo_max"`
	PhotoMaxOrig string `json:"photo_max_orig"`
	Quotes       string
	Relatives    []struct {
		Id   int
		Name string
		Type string
	} `json:"relatives"`
	Relation int
	Schools  []struct {
		Id            string // да, я тоже охуел
		Country       int
		City          int
		Name          string
		YearFrom      int
		YearTo        int
		YearGraduated int
		Class         string
		Speciality    string
		Type          int
		TypeStr       string
	}
	ScreenName   string
	Sex          int
	Site         string
	Status       string
	Timezone     int
	Trending     int
	Tv           string
	Universities []struct{}
	Verified     int
	WallDefault  string
}
