package entity

import (
  "fmt"
  "encoding/json"
  "reflect"
  "context"
  "strings"
  "github.com/olivere/elastic"
  elasticsearch "github.com/user/sites/app/services"
  SearchableDeveloper "github.com/user/sites/app/services/struct"
)



type Developer struct {
  Name        string   `json:"name"`
  Email       string   `json:"email"`
  Designation string   `json:"designation"`
  Experience  int   `json:"experience"`
  Skills      []interface{} `json:"skills"`
  CurrentCompany string   `json:"current_company"`
  IsIntern bool   `json:"is_intern"`
  ActivelyJobSearching bool   `json:"actively_job_searching"`
  Address string   `json:"address"`
  State string   `json:"state"`
  GithubUrl string   `json:"github_url"`
  LinkedinUrl string   `json:"linkedin_url"`

}


type DeveloperCollection struct {
  Developers []Developer `json:"developers"`
  Total int64 `json:"total"`
}


func in_array(val interface{}, array interface{}) (exists bool, index int) {
  exists = false
  index = -1

  switch reflect.TypeOf(array).Kind() {
  case reflect.Slice:
    s := reflect.ValueOf(array)

    for i := 0; i < s.Len(); i++ {
      if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
        index = i
        exists = true
        return
      }
    }
  }

  return
}

func GetAllDeveloperData(offset, limit int,searchFields SearchableDeveloper.DeveloperSearchableFields) DeveloperCollection  {
  ctx := context.Background()
  client := elasticsearch.CreateClient()


  searchQuery := elastic.NewBoolQuery()

    if searchFields.Name != "" {
      searchQuery.Must(elastic.NewMultiMatchQuery(searchFields.Name,"name").Type("phrase_prefix"))
    }

    if searchFields.LookingForJob == true {
      searchQuery.Must(elastic.NewTermQuery("actively_job_searching",true))
    }

  if searchFields.Intern == true {
    searchQuery.Must(elastic.NewTermQuery("is_intern",true))
  }


  if searchFields.Address != "" {
    searchQuery.Must(elastic.NewMultiMatchQuery(searchFields.Address,"address").Type("phrase_prefix"))
  }

  if searchFields.State != "" {
    searchQuery.Must(elastic.NewMultiMatchQuery(searchFields.State,"state").Type("phrase_prefix"))
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

  query := client.Search().
  Index("developer").
  Query(searchQuery).
  From(offset).Size(limit).
  Pretty(true)

  if searchFields.Sort != "" {
    sortableFields := []string{"experience", "name", "current_company","is_intern","actively_job_searching"}
    fmt.Println(in_array(searchFields.Sort,sortableFields))
    valueExists, _ := in_array(searchFields.Sort,sortableFields)

    if searchFields.Sort == "name" || searchFields.Sort == "company" {
      searchFields.Sort = searchFields.Sort+".raw"
    }
    if valueExists {
      sortType := true
      if searchFields.SortType == "desc" {
        sortType = false
      }
      query.Sort(searchFields.Sort, sortType)
    }
  }

  searchResult,err := query.Do(ctx)



  if err != nil {
    fmt.Println(err);
    fmt.Println("could not search")
  }

  var developers []Developer

  if searchResult.TotalHits() > 0 {
    for _, hit := range searchResult.Hits.Hits {
      var devData Developer
        err := json.Unmarshal(hit.Source, &devData)
        if err != nil {
            developers = append(developers,devData)
        }

        developers = append(developers,devData)
      }

  }

  return DeveloperCollection{
    Developers : developers,
    Total : searchResult.TotalHits(),
  }

}

