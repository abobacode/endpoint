package models

type AdsBlocksStruct struct {
	BlockPrefix string    `json:"block_prefix"`
	PrettyURL   string    `json:"pretty_url"`
	AdsBlocks   []DataAds `json:"ads_blocks"`
}

type DataAds struct {
	URL   string `json:"url"`
	Type  string `json:"type"`
	VodID []int  `json:"vod_ids"`
}
