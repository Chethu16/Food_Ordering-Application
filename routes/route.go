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

	r.HandleFunc("/register",user.Register).Methods("POST")
	r.HandleFunc("/login",user.Login).Methods("POST")
}