package model

type Video struct {
	URL string `json:"url"`
	Name string `json:"name"`
	Description string `json:"description"`
	Uploader User `json:"uploader"`
}
