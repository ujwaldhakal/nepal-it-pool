package elasticsearch

import (
  "github.com/olivere/elastic"
  "fmt"
  "os"
)

func CreateClient() *elastic.Client {

  var elasticSearchUrl string = os.Getenv("ELASTIC_HOST_URL")

  client, err := elastic.NewClient(
    elastic.SetURL(elasticSearchUrl),
    elastic.SetSniff(false),
    )
  if err != nil {
    fmt.Println(err)
  }

  return client
}

