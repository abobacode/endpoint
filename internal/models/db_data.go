package models

type DataStruct struct {
	//ID    int    `db:"id"` // autoincrement in database
	Title string `db:"title"`
	URL   string `db:"url"`
}
