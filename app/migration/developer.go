package main

import (
  "context"
  "encoding/json"
  "fmt"
  "github.com/olivere/elastic"
  elasticsearch "github.com/user/sites/app/services"
  entity "github.com/user/sites/app/entity"
  "io/ioutil"
   _ "reflect"
  "strconv"
)



const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
			"properties":{
				"name":{
					"type":"text"
				},
				"email":{
					"type":"text"
				},
				"designation":{
					"type":"text"
				},
				"experience":{
					"type":"integer"
				},
				"skills":{
					"type":"text"
				},
				"current_company":{
					"type":"text"
				},
				"is_intern":{
					"type":"boolean"
				},
				"actively_job_searching":{
					"type":"boolean"
				},
				"address":{
					"type":"text"
				},
				"state":{
					"type":"integer"
				},
				"github_url":{
					"type":"text"
				},
        "linkedin_url":{
					"type":"text"
				}

		}
	}
}`


func jsonToStruct() []entity.Developer {

  developers := []entity.Developer{}

  developerFileData, err := ioutil.ReadFile("../crowdSourceData/developer.json")

  if err != nil {

    fmt.Println(err)
  }

  json.Unmarshal(developerFileData, &developers)

  return developers

}


func deleteIndex(client *elastic.Client) {
  deleteIndex, err := client.DeleteIndex("developer").Do(context.Background())
  if err != nil {
    // Handle error
    panic(err)
  }
  if !deleteIndex.Acknowledged {
    fmt.Println("deleted index for migration")
  }
}

func createIndex(client *elastic.Client) {

  // Create a new index.
  createIndex, err := client.CreateIndex("developer").BodyString(mapping).Do(context.Background())
  if err != nil {
    // Handle error
    panic(err)
  }
  if !createIndex.Acknowledged {
    fmt.Println("cannot create index")
  }
}

func BulkImportDevData() bool {

  client := elasticsearch.CreateClient()


  // Use the IndexExists service to check if a specified index exists.
  exists, err := client.IndexExists("developer").Do(context.Background())
  if err != nil {
    // Handle error
    panic(err)
  }

  if exists {
      deleteIndex(client) // delete existing index
  }

  if !exists {
    createIndex(client)
  }

  bulkRequest := client.Bulk()
  n := 0
  var dev entity.Developer
  developers := jsonToStruct()
  for _, dev = range developers {
    n++
    str := elastic.NewBulkIndexRequest().Index("developer").Id(strconv.Itoa(n)).Doc(dev)
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
