package repository

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Chethu16/foodordering-system/models"
	"github.com/google/uuid"
)

type ItemStruct struct{
	DB *sql.DB
}

func (is *ItemStruct)AddItem(w http.ResponseWriter,r *http.Request){
	var additemDetails models.Items
	err:=json.NewDecoder(r.Body).Decode(&additemDetails)
	if err!=nil{
		json.NewEncoder(w).Encode(map[string]string{"message":"AddItem json decode error"})
		return
	}

	_,err=is.DB.Exec(`INSERT INTO items VALUES($1,$2,$3,$4,$5,$6,$7,$8)`,additemDetails.UserId,additemDetails.HotelId,uuid.NewString(),additemDetails.ItemName,additemDetails.ItemPrice,additemDetails.ItemRating,additemDetails.ItemType,additemDetails.ItemImage)
	if err!=nil{
		json.NewEncoder(w).Encode(map[string]string{"message":"Add items Query execution error"})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message":"Item Added Succesfully"})
}