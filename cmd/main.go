package main

import (
	//"os"
	"github.com/gorilla/mux"
	"log"
	"github.com/keeperOfTheShadows/keeperOfTheShadows.github.io/pkg/models"
	//"html/template"
	"github.com/keeperOfTheShadows/keeperOfTheShadows.github.io/pkg/router"
)


func main(){
	models.Init()
	//router.RunRouter()
	r := mux.NewRouter()

	router.RoutingGroup(r)
	

	port := 8080
	log.Printf("Server start on port: %d", port)
}
