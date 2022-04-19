package models

type Product struct {
	UUID          uint   `json:"uuid"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Price         uint   `json:"price"`
	Image         string `json:"image"`
	UUID_category uint   `json:"uuid_category"`
	UUID_user     uint   `json:"uuid_user"`
}
