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
  "os"
  "log"
)



const mapping = `{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
			"properties":{
				"name":{
					"type":"text",
          "fields": {
          "raw": {
            "type":  "keyword"
          }
        }
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
					"type":"text",
          "fields": {
          "raw": {
            "type":  "keyword"
          }
        }
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
					"type":"text"
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



func createDevelopersJsonFile() error {
    files, err := ioutil.ReadDir("./crowdSourceData")
        if err != nil {
            log.Fatal(err)
        }

        var finalFile string = "./crowdSourceData/developer.json"
        var developerJson string = "["
        for index, file := range files {
            developerData,  erro := ioutil.ReadFile("./crowdSourceData/" + file.Name())
            if erro != nil {
                log.Fatal(err)
            }

            if file.Name() != "developer.json" {
                if index != 0 {
                    developerJson += ","
                }

                developerJson += string(developerData)
            }
        }
        developerJson += "]"
    	f, err := os.OpenFile(finalFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    	if err != nil {
    		log.Fatal(err)
    	}

    	f.Truncate(0)
    	f.Seek(0,0)
    	f.Sync()

    	if _, err := f.WriteString(developerJson); err != nil {
    		f.Close() // ignore error; Write error takes precedence
    		log.Fatal(err)
    	}
    	if err := f.Close(); err != nil {
    		log.Fatal(err)
    	}

        fmt.Println(developerJson)

        return nil
}

func jsonToStruct() []entity.Developer {

  developers := []entity.Developer{}

  developerFileData, err := ioutil.ReadFile("/go/src/github.com/user/sites/app/crowdSourceData/developer.json")

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

    createDevelopersJsonFile()

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
  createIndex(client)

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
