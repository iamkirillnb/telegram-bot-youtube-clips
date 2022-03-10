package entities

type RestResponse struct {
	Items []Item `json:"items"`
}

type Item struct {
	Id Info `json:"id"`
}

type Info struct {
	VideoId string `json:"videoId"`
}
