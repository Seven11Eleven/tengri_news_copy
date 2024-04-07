package models

type News struct{
	Link string `json:"link"`
	Photo  string `json:"photo"`
	Title  string `json:"title"`
	Time   string `json:"time"`
}

type OneNew struct{
	Text string `json:"text"`
}