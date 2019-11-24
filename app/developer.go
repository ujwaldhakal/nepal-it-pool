package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"fmt"
	"encoding/json"
	_ "reflect"
	"github.com/user/sites/app/entity"
  SearchableDeveloper "github.com/user/sites/app/services/struct"

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
  var searchQuery SearchableDeveloper.DeveloperSearchableFields
  router.BindQuery(&searchQuery)

  developerFileData, err := ioutil.ReadFile("crowdSourceData/developer.json")

    if err != nil {

        fmt.Println(err)
	}


	offset := router.Query("offset")
	limit := router.Query("limit")


	json.Unmarshal([]byte(developerFileData), &developer)

	if offset == "" {
		offset = "0"
	}

	if limit == "" {
		limit = "10"
	}


	DataLimit, err := strconv.Atoi(limit)
	DataOffset, err := strconv.Atoi(offset)

	fmt.Println(limit)
	fmt.Println(offset)

	data := entity.GetAllDeveloperData(DataOffset,DataLimit,searchQuery)
	router.JSON(200, gin.H{
		"status": "up",
		"data" : data.Developers,
		"total" : data.Total,
	})
}


func registerDeveloperRoutes(routes *gin.Engine) {

	routes.GET("/", listDevelopers)
}


