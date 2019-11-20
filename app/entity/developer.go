package entity

import (
  "fmt"
  "encoding/json"
  //_ "reflect"
  "context"
  "strings"
  "github.com/olivere/elastic"
  elasticsearch "github.com/user/sites/app/services"
  SearchableDeveloper "github.com/user/sites/app/services/struct"
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


func GetAllDeveloperData(offset, limit int,searchFields SearchableDeveloper.DeveloperSearchableFields) DeveloperCollection  {
  ctx := context.Background()
  client := elasticsearch.CreateClient()

  fmt.Println("ok earching")
  fmt.Println(searchFields)

  searchQuery := elastic.NewBoolQuery()

    if searchFields.Name != "" {
      searchQuery.Must(elastic.NewMultiMatchQuery(searchFields.Name,"name").Type("phrase_prefix"))
    }

    if searchFields.Designation != "" {
      searchQuery.Must(elastic.NewMultiMatchQuery(searchFields.Designation,"designation").Type("phrase_prefix"))
    }

    if searchFields.MaxExperience != "" {
      searchQuery.Must(elastic.NewRangeQuery("experience").Lt(searchFields.MaxExperience))
    }


    if searchFields.MinExperience != "" {
      searchQuery.Must(elastic.NewRangeQuery("experience").Gt(searchFields.MinExperience))
    }

    fmt.Println(searchFields.Skills)

  if len(searchFields.Skills) > 0 {
    input := strings.Split(searchFields.Skills,",") //as newTerms only accepts multiple argument we had to turn to iterface
    values := make([]interface{}, len(input))
    for i, s := range input {
      values[i] = s
    }
    searchQuery.Must(elastic.NewTermsQuery("skills",values...))
  }

	// Get tweet with specified ID
  searchResult, err := client.Search().
  Index("developer").   // search in index "twitter"
  Query(searchQuery).
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

