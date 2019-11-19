package entity

import (
  "fmt"
  "encoding/json"
  //_ "reflect"
  "context"
  "github.com/olivere/elastic"
  elasticsearch "github.com/user/sites/app/services"
)



type Developer struct{

	Name string `json:"name"`
	Email string `json:"email"`
	Designation string `json:"designation"`
	Experience string `json:"experience"`
	Skills []string `json:"skills"`

}

type DeveloperCollection struct {
  Developers []Developer
  Total int64
}


func GetAllDeveloperData(offset, limit int, searchQuery DeveloperSearchQuery) DeveloperCollection  {
  ctx := context.Background()
  client := elasticsearch.CreateClient()

	// Get tweet with specified ID
  termQuery := elastic.NewTermQuery("name", "ujwal")
  searchResult, err := client.Search().
  Index("developer").   // search in index "twitter"
  Query(termQuery).
  From(offset).Size(limit).   // take documents 0-9
  Pretty(true).       // pretty print request and response JSON
  Do(ctx)

  if err != nil {
    fmt.Println(err);
    fmt.Println("could not search")
  }

  var developers []Developer
  var devData Developer

  if searchResult.TotalHits() > 0 {

    for _, hit := range searchResult.Hits.Hits {

        err := json.Unmarshal(hit.Source, &devData)
        if err != nil {
            developers = append(developers,devData)
        }

        developers = append(developers,devData)
      }

    //fmt.Println(reflect.TypeOf(developers))
    fmt.Println(developers)


  }

  return DeveloperCollection{
    Developers : developers,
    Total : searchResult.TotalHits(),
  }

}

