package models

type Users struct{
	UserId string `json:"user_id"`
	UserName string `json:"user_name"`
	UserEmail string `json:"user_email"`
	UserPassword string `json:"user_password"`
	UserPhone string `json:"user_phone"`
	UserAddress string `json:"user_address"`
}
