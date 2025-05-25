package controllers

import (
	"html/template"
	"net/http"
	"io"
	"fmt"
	"os"
	"path/filepath"
	//"github.com/gin-gonic/gin"
	//"github.com/keeperOfTheShadows/keeperOfTheShadows.github.io/pkg/models"
)

var (
	tmpl *template.Template
)

func init(){
	var err error
	tmpl, err = template.ParseGlob("templates/*.html")
	if err != nil{
		fmt.Println("Error parsing templates:", err)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request){
	tmpl.ExecuteTemplate(w, "images", nil)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	//Handle the multipart form size
	r.ParseMultipartForm(10 << 20)

	//Retrieve the file from form data
	file, handler, err := r.FormFile("Image")
	if err != nil {
		if err == http.ErrMissingFile {
			fmt.Println("No file submitted", err)
		}else{
			fmt.Println("There is an error in uploading the file", err)
		}
	}

	defer file.Close()

	fileName := filepath.Ext(handler.Filename)

	filePath := filepath.Join("images", fileName)

	dst, err := os.Create(filePath)
	if err != nil{
		fmt.Println("Error in saving file in the server", err)
		return
	}
	
	defer dst.Close()
	
	if _, err := io.Copy(dst, file); err != nil{
		fmt.Println("Error in copy file to server", err)
		return 
	}
}


/*func getImagesFromFilePath(c *gin.Context){
	var Image models.ImagesFile
	
	f, err := os.Open(filePath)
}*/
