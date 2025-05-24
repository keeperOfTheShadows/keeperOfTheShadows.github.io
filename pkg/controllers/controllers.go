package controllers

import (
	"html/template"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/keeperOfTheShadows/keeperOfTheShadows.github.io/pkg/models"
)

var (
	tmpl *template.Template
)

func init(){
	tmpl, _ = template.ParseGlob("templates/*.html")
}

func HomeHandler(w http.ResponseWriter, r *http.Request){
	tmpl.ExecuteTemplate(w, "images", nil)
}

func getImagesFromFilePath(c *gin.Context){
	var Image models.ImagesFile
	
	f, err := os.Open(filePath)
}
