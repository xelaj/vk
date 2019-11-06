package types

type App struct {
	AuthorGroup     int     `json:"author_group"`
	AuthorId        int     `json:"author_id"`
	AuthorUrl       string  `json:"author_url"`
	Banner1120      string  `json:"banner_1120"`
	Banner560       string  `json:"banner_560"`
	Friends         []int   `json:"friends"`
	CatalogPosition int     `json:"catalog_position"`
	Description     string  `json:"description"`
	Genre           string  `json:"genre"`
	GenreId         int     `json:"genre_id"`
	Icon139         string  `json:"icon_139"`
	Icon150         string  `json:"icon_150"`
	Icon278         string  `json:"icon_278"`
	Icon75          string  `json:"icon_75"`
	Id              int     `json:"id"`
	International   int     `json:"international"`
	IsInCatalog     int     `json:"is_in_catalog"`
	LeaderboardType int     `json:"leaderboard_type"`
	Members_count   int     `json:"members_count"`
	PlatformId      int     `json:"platform_id"`
	PublishedDate   int     `json:"published_date"`
	ScreenName      string  `json:"screen_name"`
	Screenshots     []Photo `json:"screenshots"`
	Section         string  `json:"section"`
	Title           string  `json:"title"`
	Type            string  `json:"type"`
}
