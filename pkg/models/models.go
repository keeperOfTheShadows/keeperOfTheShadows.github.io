package models

import (
	"github.com/keeperOfTheShadows/keeperOfTheShadows.github.io/pkg/config"
	"fmt"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

var (
	DB *gorm.DB
)

type (
	ImagesFile struct{
		ID int `gorm:"primaryKey;autoIncrement"`
		FileName string `gorm:"not null"`
		MimeType string	`gorm:"not null"`
		UploadAt time.Time `gorm:"autoCreateTime"`
		Data	[]byte	`gorm:"type:longblob"`
	}
)

func Init(){
	DB = config.Connection()
	if DB == nil{
		panic("DB is not Initialized")
	}
	err := DB.AutoMigrate(&ImagesFile{})
	if err != nil{
		log.Fatal("Failed to migrate database:", err)
	}
}

func (i *ImagesFile) AddImage() (*ImagesFile, error){
	result := DB.Create(&i)
	if result.Error != nil {
		return nil, fmt.Errorf("Error in Adding image")
	}
	return i, nil
}

func GetAllImages() []ImagesFile{
	var result []ImagesFile
	DB.Find(&result)
	return result
}


