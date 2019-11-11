package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"fmt"
	"encoding/json"
)


type developerType struct{

	Name string `json:"name"`
	Email string `json:"email"`
	Designation string `json:"designation"`
	Experience string `json:"experience"`
	Skills []string `json:"skills"`

}

type developerArray struct {
	developer []developerType
}

func listDevelopers(router *gin.Context) {
	var developer []developerType

	developerFileData, err := ioutil.ReadFile("CrowdSourceData/developer.json")
    
    if err != nil {
        fmt.Println(err)
	}
	
	offset := router.Query("offset")

	json.Unmarshal([]byte(developerFileData), &developer)

	if offset != "" {
		
	}


	router.JSON(200, gin.H{
		"status": "up",
		"data" : developer,
	})
}



func registerDeveloperRoutes(routes *gin.Engine) {

	routes.GET("/developers", listDevelopers)
}


