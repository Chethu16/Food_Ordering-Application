package models

type Items struct{
	UserId string `json:"user_id"`
	HotelId string `json:"hotel_id"`
	ItemId string `json:"item_id"`
	ItemName string `json:"item_name"`
	ItemPrice string `json:"item_price"`
	ItemRating string `json:"item_rating"`
	ItemType string `json:"item_type"`
	ItemImage string `json:"item_image"`
}