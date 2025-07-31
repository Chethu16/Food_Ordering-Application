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

	fmt.Println(additemDetails.UserId)

	_,err=is.DB.Exec(`INSERT INTO items VALUES($1,$2,$3,$4,$5,$6,$7,$8)`,additemDetails.UserId,additemDetails.HotelId,uuid.NewString(),additemDetails.ItemName,additemDetails.ItemPrice,additemDetails.ItemRating,additemDetails.ItemType,additemDetails.ItemImage)
	fmt.Println(err)
	if err!=nil{
		json.NewEncoder(w).Encode(map[string]string{"message":"Add items Query execution error"})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message":"Item Added Succesfully"})
}
func(is *ItemStruct)GetItem(w http.ResponseWriter,r *http.Request){
	var vars = mux.Vars(r)
	var getitem = vars["getitem_id"]
	res,err:=is.DB.Query(`SELECT item_name,item_price,item_rating,item_type,item_image FROM items WHERE hotel_id=$1`,getitem)
	if err!=nil{
		json.NewEncoder(w).Encode(map[string]string{"Message":"Get items query execution error"})
		return
	}
	defer res.Close()
	var getitems []models.Items
	for res.Next(){
		var Getitems models.Items
		err=res.Scan(&Getitems.ItemName,&Getitems.ItemPrice,&Getitems.ItemType,&Getitems.ItemRating,&Getitems.ItemImage)
		if err!=nil{
			json.NewEncoder(w).Encode(map[string]string{"Message":"Error occured get items 1"})
			return
		}
		getitems=append(getitems, Getitems)
	}
	if res.Err()!=nil{
		json.NewEncoder(w).Encode(map[string]string{"Message":"Error occured get items 2"})
		return
	}
	json.NewEncoder(w).Encode(map[string][]models.Items{"Items":getitems})
}
func (is *ItemStruct)DeleteItem(w http.ResponseWriter,r *http.Request){
	var vars = mux.Vars(r)
	var deleteitem = vars ["deletetem_id"]

	_,err:=is.DB.Exec(`DELETE FROM items WHERE hotel_id=$1`,deleteitem)
	if err!=nil{
		json.NewEncoder(w).Encode(map[string]string{"Message":"Delete Items query execution error"})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"Message":"Items deleted Succesfully"})
}