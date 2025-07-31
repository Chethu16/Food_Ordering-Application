package models

type Hotels struct{
	LocationId string `json:"location_id"`
	UserId string `json:"user_id"`
	HotelId string `json:"hotel_id"`
	HotelName string `json:"hotel_name"`
	HotelCategory string `json:"hotel_category"`
	HotelType string `json:"hotel_type"`
	HotelRating string `json:"hotel_rating"`
	HotelLocation string `json:"hotel_location"`
	HotelOpen string `json:"hotel_open"`
	HotelClose string `json:"hotel_close"`
	HotelImage string `json:"hotel_image"`

}