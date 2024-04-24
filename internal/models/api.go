package models

type DataProject struct {
	ServiceName string  `json:"service_name"`
	ContentData Project `json:"content_data"`
}

type Project struct {
	ForeignID   int           `json:"foreign_id"`
	Title       string        `json:"title"`
	Story       string        `json:"story"`
	PosterHoriz string        `json:"poster_horiz"`
	Poster      string        `json:"poster"`
	Genres      []string      `json:"genres"`
	Countries   []int         `json:"countries"`
	Categories  []int         `json:"categories"`
	Tags        []string      `json:"tags"`
	IsFree      bool          `json:"is_free"`
	Year        int           `json:"year"`
	AgeLimit    int           `json:"age_limit"`
	Premiere    string        `json:"premiere"`
	Seasons     []SeasonsType `json:"seasons"`
}

type SeasonsType struct {
	SeasonID string         `json:"season_id"`
	Title    string         `json:"title"`
	Episodes []EpisodesType `json:"episodes"`
}

type EpisodesType struct {
	SeasonID    string `json:"season_id"`
	ForeignID   int    `json:"foreign_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	Image       string `json:"image"`
	IsFree      bool   `json:"is_free"`
	Src         string `json:"src"`
}

type Deleted struct {
	DeletedVideos []int `json:"deleted_videos"`
}
