package repository

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Chethu16/foodordering-system/models"
	"github.com/google/uuid"
)


type AddlocationStruct struct{
	DB *sql.DB
}

func(al *AddlocationStruct)AddLocation(w http.ResponseWriter,r *http.Request){
	var addlocationDetail models.Location
	err:=json.NewDecoder(r.Body).Decode(&addlocationDetail)
	if err!=nil{
		json.NewEncoder(w).Encode(map[string]string{"message":"AddLocation decode error"})
		return
	}
	_,err=al.DB.Exec(`INSERT INTO locations VALUES($1,$2)`,uuid.NewString(),addlocationDetail.LocationName)
	if err!=nil{
		json.NewEncoder(w).Encode(map[string]string{"message":"Addlocation Query execution error"})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message":"Location Added Succesfully"})
}
