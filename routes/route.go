package routes

import (
	"database/sql"

	"github.com/Chethu16/foodordering-system/repository"
	"github.com/gorilla/mux"
)



func InitializeRoutes(r *mux.Router,connection *sql.DB){
	var user = repository.Repo{
		DB: connection,
	}
	var hotel = repository.Hotelstruct{
		DB: connection,
	}
	var item = repository.ItemStruct{
		DB: connection,
	}

	r.HandleFunc("/register",user.Register).Methods("POST")
	r.HandleFunc("/login",user.Login).Methods("POST")
	r.HandleFunc("/addhotel",hotel.AddHotel).Methods("POST")
	r.HandleFunc("/additem",item.AddItem).Methods("POST")
	r.HandleFunc("/deletehotel/{hotel_id}",hotel.DeleteHotel).Methods("GET")
	r.HandleFunc("/gethotel/{gethotel_id}",hotel.Gethotel).Methods("GET")
}