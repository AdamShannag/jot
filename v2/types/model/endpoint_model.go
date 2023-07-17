package model

type Endpoint struct {
	Name string `json:"name"`
	Urls []Url  `json:"urls"`
}
