package models

type Ads struct {
	//ID   int    `db:"id"`
	URL  string `db:"url"`
	Type string `db:"type"`
}

type AdsStruct struct {
	VideoID int    `db:"video_id"`
	Pre     []byte `db:"pre"`
	Mid     []byte `db:"mid"`
	Pause   []byte `db:"pause"`
	Post    []byte `db:"post"`
	Etag    string `db:"etag"`
}

type AdsBlock struct {
	Pre   []byte `db:"pre"`
	Mid   []byte `db:"mid"`
	Pause []byte `db:"pause"`
	Post  []byte `db:"post"`
}
