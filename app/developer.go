package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"github.com/user/sites/app/entity"
	 "strconv"
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

	developerFileData, err := ioutil.ReadFile("crowdSourceData/developer.json")
    
    if err != nil {

        fmt.Println(err)
	}

	offset := router.Query("offset")
	limit := router.Query("limit")

	
	json.Unmarshal([]byte(developerFileData), &developer)

	if offset == "" {
		offset = "0"
		fmt.Println("ok i do")

	}

	if limit == "" {
		limit = "1"
	
	}

	
	
	DataLimit, err := strconv.Atoi(limit)
	DataOffset, err := strconv.Atoi(offset)

	fmt.Println(limit)
	fmt.Println(offset)

	
	router.JSON(200, gin.H{
		"status": "up",
		"data" : entity.GetAllDeveloperData(DataOffset,DataLimit),
	})
}


func registerDeveloperRoutes(routes *gin.Engine) {

	routes.GET("/developers", listDevelopers)
}


