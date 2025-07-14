package systems

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
)

var EsClient *elasticsearch.Client

func InitElasticsearch() error {
	conf := GetConfig().Elasticsearch
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: conf.Addresses,
	})
	if err != nil {
		return err
	}
	EsClient = es
	fmt.Println("✅ ElasticSearch链接成功")
	return nil
}
