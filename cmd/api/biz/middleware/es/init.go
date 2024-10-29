package es

import (
	"fmt"
	"log"

	"github.com/wushiling50/aster/config"

	elasticsearch "github.com/elastic/go-elasticsearch"
)

var (
	EsClient *elasticsearch.Client
)

// 初始化es
func Init() {
	esConn := fmt.Sprintf("http://%s", config.Elasticsearch.Addr)
	cfg := elasticsearch.Config{
		Addresses: []string{esConn},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Panic(err)
	}
	EsClient = client
}
