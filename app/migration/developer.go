package main

import (
  "context"
  "encoding/json"
  "fmt"
  "github.com/olivere/elastic"
  elasticsearch "github.com/user/sites/app/services"
  "io/ioutil"
   _ "reflect"
  "strconv"
)

type developer struct {
  Name        string   `json:"name"`
  Email       string   `json:"email"`
  Designation string   `json:"designation"`
  Experience  string   `json:"experience"`
  Skills      []interface{} `json:"skills"`
}



func jsonToStruct() []developer {

  developers := []developer{}

  developerFileData, err := ioutil.ReadFile("../crowdSourceData/developer.json")

  if err != nil {

    fmt.Println(err)
  }

  json.Unmarshal(developerFileData, &developers)

  return developers

}

func BulkImportDevData() bool {

  client := elasticsearch.CreateClient()

  bulkRequest := client.Bulk()
  n := 0
  var dev developer
  developers := jsonToStruct()
  for _, dev = range developers {
    n++
    str := elastic.NewBulkIndexRequest().Index("developer").Type("details").Id(strconv.Itoa(n)).Doc(dev)
    bulkRequest = bulkRequest.Add(str)
  }

  // Do sends the bulk requests to Elasticsearch
  response, err := bulkRequest.Do(context.Background())
  if err != nil {
    // ...
    fmt.Println(err)
    return false
  }


  fmt.Printf("migration ran successfully imported %v\n documents",len(response.Items))

  return true
}


func main() {
  BulkImportDevData()
}
