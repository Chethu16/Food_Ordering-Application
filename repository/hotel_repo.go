package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Chethu16/foodordering-system/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Hotelstruct struct{
	DB *sql.DB
}

func (h *Hotelstruct)AddHotel(w http.ResponseWriter, r *http.Request){
	var addhotelDetails models.Hotels
	err:=json.NewDecoder(r.Body).Decode(&addhotelDetails)
	if err!=nil{
		json.NewEncoder(w).Encode(map[string]string{"message":"Addhotel json decode error"})
		return
	}

	_,err=h.DB.Exec(`INSERT INTO hotels VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`,addhotelDetails.UserId,uuid.NewString(),addhotelDetails.HotelName,addhotelDetails.HotelCategory,addhotelDetails.HotelType,addhotelDetails.HotelRating,addhotelDetails.HotelLocation,addhotelDetails.HotelOpen,addhotelDetails.HotelClose,addhotelDetails.HotelImage)
	if err!=nil{
		json.NewEncoder(w).Encode(map[string]string{"message":"Query execution error"})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message":"Hotels added Succesfully"})
}
func(h *Hotelstruct)DeleteHotel(w http.ResponseWriter,r *http.Request){
	var vars = mux.Vars(r)
	var hotelid =vars["hotel_id"]
	_,err:=h.DB.Exec(`DELETE FROM hotels WHERE hotel_id=$1`,hotelid)
	if err!=nil{
		json.NewEncoder(w).Encode(map[string]string{"message":"Delete Hotel Query execution error"})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message":"Hotel deleted Succesfully"})
}
func (h *Hotelstruct)Gethotel(w http.ResponseWriter,r *http.Request){
	var vars =mux.Vars(r)
	var gethotel = vars["gethotel_id"]

	res,err:=h.DB.Query(`SELECT hotel_name,hotel_category,hotel_type,hotel_rating,hotel_location,hotel_open,hotel_close,hotel_image FROM hotels WHERE hotel_id=$1`,gethotel)
		fmt.Println(err)
	if err!=nil{
		json.NewEncoder(w).Encode(map[string]string{"message":"Get hotel query execution error"})
		return
	}
	defer res.Close()
	var hotels []models.Hotels

	for res.Next(){
		var hotel models.Hotels
		err=res.Scan(&hotel.HotelName,&hotel.HotelCategory,&hotel.HotelType,&hotel.HotelRating,&hotel.HotelLocation,&hotel.HotelOpen,&hotel.HotelClose,&hotel.HotelImage)
		if err!=nil{
			json.NewEncoder(w).Encode(map[string]string{"message":"Error occured 1"})
			return
		}
		hotels=append(hotels, hotel)
	}
	if res.Err()!=nil{
		json.NewEncoder(w).Encode(map[string]string{"message":"Error occured 2"})
		return
	}
	json.NewEncoder(w).Encode(map[string][]models.Hotels{"Hotels":hotels})
}