package models

type Epg struct {
	ID             int64  `db:"id"`
	TimeStart      string `db:"timestart"`
	TimeStop       string `db:"timestop"`
	Title          string `db:"title"`
	Desc           string `db:"desc"`
	EpgID          int    `db:"epg_id"`
	Date           string `db:"date"`
	CDNVideo       uint8  `db:"cdnvideo"`
	Rating         uint8  `db:"rating"`
	DirectorID     int64  `db:"director_id"`
	Year           uint8  `db:"year"`
	TimeZone       string `db:"time_zone"`
	ArchiveEnabled uint8  `db:"archive_enabled"`
	PostersJSON    string `db:"posters_json"`
	ParserType     int    `db:"parser_type"`
	AdUUID         string `db:"ad_uuid"`
	VideoURL       string `db:"video_url"`
	UnixStart      int64  `db:"unixstart"`
	UnixStop       int64  `db:"unixstop"`
}
