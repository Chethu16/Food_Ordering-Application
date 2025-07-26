package datatbase

import (
	"database/sql"
	"log"
	

	_ "github.com/lib/pq"
)

func DatabaseConnection(url string) *sql.DB{
	cnn,err:=sql.Open("postgres",url)
	if err!=nil{
		log.Fatalf("Unable to connect database :%v",err)
		
	}
	log.Println("Databse Connected Successfully")
	return cnn

}
func Initializing(databaseconnection *sql.DB){
	var queries = []string{
		`CREATE TABLE IF NOT EXISTS users(
		user_id VARCHAR NOT NULL PRIMARY KEY,
		user_name VARCHAR NOT NULL,
		user_email VARCHAR NOT NULL,
		user_password VARCHAR NOT NULL,
		user_phone_no VARCHAR NOT NULL,
		user_address VARCHAR NOT NULL
		
		
		)`,
		`CREATE TABLE IF NOT EXISTS hotels(
		user_id VARCHAR NOT NULL,
		hotel_id VARCHAR NOT NULL PRIMARY KEY,
		hotel_name VARCHAR NOT NULL,
		hotel_category VARCHAR NOT NULL,
		hotel_type VARCHAR NOT NULL,
		hotel_rating VARCHAR NOT NULL, 
		hotel_location VARCHAR NOT NULL,
		hotel_open VARCHAR NOT NULL,
		hotel_close VARCHAR NOT NULL,
		hotel_image VARCHAR NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE

		)`,
		`CREATE TABLE IF NOT EXISTS items(
		user_id VARCHAR NOT NULL,
		hotel_id VARCHAR NOT NULL,
		item_id VARCHAR NOT NULL PRIMARY KEY,
		item_name VARCHAR NOT NULL,
		item_price VARCHAR NOT NULL,
		item_rating VARCHAR NOT NULL,
		item_type VARCHAR NOT NULL,
		item_image VARCHAR NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
		FOREIGN KEY (hotel_id) REFERENCES hotels(hotel_id) ON DELETE CASCADE
		)`,
	}
	for _,query:=range queries{
		_,err:=databaseconnection.Exec(query) 
		if err!=nil{
			log.Fatalf("Unable to initialize :%v",err)

		}
	}
	log.Println("Database Initialized Succesfully")
}