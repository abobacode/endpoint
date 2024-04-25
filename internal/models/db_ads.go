package models

type AdsStruct struct {
	VideoID int    `db:"video_id"`
	Pre     []byte `db:"pre"`
	Mid     []byte `db:"mid"`
	Pause   []byte `db:"pause"`
	Post    []byte `db:"post"`
}
