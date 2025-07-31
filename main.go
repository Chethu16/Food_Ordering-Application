package main

import (
	"net/http"

	"github.com/Chethu16/foodordering-system/datatbase"
	"github.com/Chethu16/foodordering-system/routes"
	"github.com/gorilla/mux"
)

func main(){

	databaseconnection:=datatbase.DatabaseConnection("postgresql://foodordering_nbl4_user:eEh7AkLZ0KLBXximitxnuHYBydBpBhYZ@dpg-d229qf15pdvs738np31g-a.oregon-postgres.render.com/foodordering_nbl4")
	defer databaseconnection.Close()
	datatbase.Initializing(databaseconnection)

	var route = mux.NewRouter()

	routes.InitializeRoutes(route,databaseconnection)



	http.ListenAndServe(":8003",route)
}