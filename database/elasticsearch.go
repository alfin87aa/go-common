package database

import (
	"context"

	"github.com/alfin87aa/go-common/configs"
	"github.com/alfin87aa/go-common/logger"
	"github.com/alfin87aa/go-common/servers/restapi"
	"github.com/elastic/go-elasticsearch/v8"
)

func initElasticsearch(ctx context.Context) {
	config := configs.Configs.DB.ElasticSearch

	var err error
	elasticClient, err = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: config.Address,
		Username:  config.Username,
		Password:  config.Password,
	})

	if err != nil {
		logger.Fatalf(ctx, err, "❌ Elasticsearch client failed to connect")
	}

	logger.Infoln(ctx, "✅ Elasticsearch client connected")

	restapi.AddChecker("elasticsearch", func(ctx context.Context) error {
		_, err := elasticClient.Info()
		return err
	})
}

func GetElasticsearchClient() *elasticsearch.Client {
	return elasticClient
}
