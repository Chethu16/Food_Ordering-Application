package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/Chethu16/foodordering-system/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type Repo struct{
	DB *sql.DB
}


func(rg *Repo) Register(w http.ResponseWriter,r *http.Request){
	var registerDetails models.Users
	err:=json.NewDecoder(r.Body).Decode(&registerDetails)
	if err!=nil{
		json.NewEncoder(w).Encode(map[string]string{"message":"Register decode error"})
		return
	}
	enp,err:=bcrypt.GenerateFromPassword([]byte(registerDetails.UserPassword),bcrypt.DefaultCost)
	if err!=nil{
		json.NewEncoder(w).Encode(map[string]string{"message":"Password Encrypt error"})
		return
	}
	_,err=rg.DB.Exec(`INSERT INTO users VALUES($1,$2,$3,$4,$5,$6)`,uuid.NewString(),registerDetails.UserName,registerDetails.UserEmail,string(enp),registerDetails.UserPhone,registerDetails.UserAddress)
	if err!=nil{
		json.NewEncoder(w).Encode(map[string]string{"messsage":"Query execution error"})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message":"Registerd Succesfully"})
}

func (rg *Repo)Login(w http.ResponseWriter,r *http.Request){
	var loginDetails models.Users
	err:=json.NewDecoder(r.Body).Decode(&loginDetails)
	if err!=nil{
		json.NewEncoder(w).Encode(map[string]string{"message":"login decode error"})
		return
	}

	var dbuserid,dbuserpassword string
	err=rg.DB.QueryRow(`SELECT user_id,user_password FROM users WHERE user_email=$1`,loginDetails.UserEmail).Scan(&dbuserid,&dbuserpassword)
	fmt.Println(err)
	if err!=nil{
		json.NewEncoder(w).Encode(map[string]string{"message":"Password scan error"})
		return
	}
	err=bcrypt.CompareHashAndPassword([]byte(dbuserpassword),[]byte(loginDetails.UserPassword))
	if err!=nil{
		json.NewEncoder(w).Encode(map[string]string{"message":"Incorrect Password"})
		return
	}
		json.NewEncoder(w).Encode(map[string]string{"user_id":dbuserid})
}
func(rg *Repo)DeleteUser(w http.ResponseWriter,r *http.Request){
	var vars = mux.Vars(r)
	var deleteuser = vars["deleteuser_id"]

	_,err:=rg.DB.Exec(`DELETE FROM users WHERE user_id=$1`,deleteuser)
	if err!=nil{
		json.NewEncoder(w).Encode(map[string]string{"message":"Deleteuser Query execution error"})
		return
	}
		json.NewEncoder(w).Encode(map[string]string{"message":"User deleted Succesfully"})
}
	