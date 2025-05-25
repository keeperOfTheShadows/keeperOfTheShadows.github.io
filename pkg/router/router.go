package router

import (
	"github.com/gorilla/mux"
	//"github.com/keeperOfTheShadows/keeperOfTheShadows.github.io/pkg/models"
	"github.com/keeperOfTheShadows/keeperOfTheShadows.github.io/pkg/controllers"
	//"os"
	//"github.com/gin-gonic/gin"
	//"image"
	//"fmt"
	//"net/http"
)


//type (
//	ImageFile struct{
//		image_file	image.Image
//	}
//)

//var (
//	Images []ImageFile
//)

//var (
//	Router = gin.Default()
//	UnProtected = Router.Group("/unprotected")
	//UnProtected.POST("/register", controllers.Register)
	//UnProtected.POST("/users", controllers.)
	//Protected = Router.Group("/protected")
//)


/*func getImageFromFilePath(filePath string)(error) {
	var Image models.ImagesFile

	f, err := os.Open(filePath)
	if err != nil{
		return fmt.Errorf("Error in reading Image files", err)
	}
	defer f.Close()

	//Images = append(Images, image)
	image, _, err := image.Decode(f)
	if err != nil{
		fmt.Println("Error in decoding the image", err)
	}

	var Picture models.ImageFile{FileName:filePath, Data: image}
	

	Image, err = Picture.models.AddImage()
	if err != nil{
		return fmt.Errorf("Error in Adding the image", err)
	}
	
}*/

var RoutingGroup = func(router *mux.Router){
	router.HandleFunc("https://keeperoftheshadows.github.io/images", controllers.UploadHandler).Methods("POST")
	router.HandleFunc("https://keeperoftheshadows.github.io/home", controllers.HomeHandler).Methods("GET")
}

/*func RunRouter(){ 
	Router.Use(gin.Recovery())

	/*var AllImages []models.ImagesFile
	
	AllImages = models.GetAllImages()

	for i, _ := range AllImages {
		Unprotected.GET(AllImages[i].FileName, getImageFromFilePath)
	}*/

//	UnProtected.POST("/images", controllers.UploadHandler)
	
