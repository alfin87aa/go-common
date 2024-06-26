package database

import (
	"context"
	"fmt"

	"github.com/alfin87aa/go-common/configs"
	"github.com/alfin87aa/go-common/logger"
	"github.com/alfin87aa/go-common/servers/restapi"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initMongo(ctx context.Context) {
	config := configs.Configs.DB.Mongo

	var err error

	mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%s@%s:%d/%s",
			config.Username,
			config.Password,
			config.Host,
			config.Port,
			config.DataBase,
		),
	))

	if err != nil {
		logger.Fatalf(ctx, err, "❌ MongoDB client failed to connect")
	}

	logger.Infoln(ctx, "✅ MongoDB client connected")

	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			logger.Fatalf(ctx, err, "❌ MongoDB client failed to disconnect")
		}
	}()

	restapi.AddChecker("mongo", func(ctx context.Context) error {
		return mongoClient.Ping(ctx, nil)
	})
}

func GetMongoClient() *mongo.Client {
	return mongoClient
}
