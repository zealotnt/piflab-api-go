package models

type Product struct {
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Provider  string `json:"provider"`
	Rating    int    `json:"rating"`
	Status    string `json:"status"`
	ImageUrl  string `json:"image_url"`
	DetailUrl string `json:"detail_url"`
}
