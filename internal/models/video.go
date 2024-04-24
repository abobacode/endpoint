package models

import "time"

type Video struct {
	ProjectID          int       `db:"project_id"`
	VideoID            int       `db:"video_id"`
	Title              string    `db:"title"`
	TitleProject       string    `db:"title_project"`
	Description        string    `db:"description"`
	DescriptionProject string    `db:"description_project"`
	Categories         []byte    `db:"categories"`
	Duration           int       `db:"duration"`
	ViewsCount         int       `db:"views_count"`
	PosterURL          string    `db:"poster_url"`
	PosterSqrURL       string    `db:"poster_sqr_url"`
	URL                string    `db:"url"`
	AiredAt            time.Time `db:"aired_at"`
	VideoType          string    `db:"video_type"`
	AgeRestriction     int       `db:"age_restriction"`
	Format             string    `db:"format"`
	Src                string    `db:"src"`
	Type               string    `db:"type"`
	Status             string    `db:"status"`
	StatusDesc         string    `db:"status_description"`
	Etag               string    `db:"etag"`
	M3u8URL            string    `db:"m3u8_url"`
	UpdatedAt          string    `db:"updated_at"`
}

type VideoModifiedURL struct {
	URL        string `db:"url"`
	Etag       string `db:"etag"`
	Status     string `db:"status"`
	StatusDesc string `db:"status_description"`
}

type VideoModifiedAll struct {
	ProjectID          int       `db:"project_id"`
	Title              string    `db:"title"`
	TitleProject       string    `db:"title_project"`
	Description        string    `db:"description"`
	DescriptionProject string    `db:"description_project"`
	Categories         []byte    `db:"categories"`
	Duration           int       `db:"duration"`
	ViewsCount         int       `db:"views_count"`
	PosterURL          string    `db:"poster_url"`
	PosterSqrURL       string    `db:"poster_sqr_url"`
	AiredAt            time.Time `db:"aired_at"`
	VideoType          string    `db:"video_type"`
	AgeRestriction     int       `db:"age_restriction"`
	Format             string    `db:"format"`
	Src                string    `db:"src"`
	Type               string    `db:"type"`
	Status             string    `db:"status"`
	StatusDesc         string    `db:"status_description"`
	M3u8URL            string    `db:"m3u8_url"`
	UpdatedAt          string    `db:"updated_at"`
}

type VideoDelete struct {
	Status     string `db:"status"`
	StatusDesc string `db:"status_description"`
}
