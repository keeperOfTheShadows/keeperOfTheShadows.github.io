package main

import (
	"os"
	"log"
	"github.com/keeperOfTheShadows/keeperOfTheShadows.github.io/pkg/models"
	"html/template"
	"github.com/keeperOfTheShadows/keeperOfTheShadows.github.io/pkg/router"
)

var (
	tmpl *template.Template
)

func init(){
	tmpl, _ = template.ParseGlob("templates/*.html")
}

func main(){
	models.Init()
	router.RunRouter()

	port := 8080
	log.Printf("Server start on port: %s", port)
}
