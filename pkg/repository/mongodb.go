package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	URI      string
	Database string
	Timeout  time.Duration
}

type MongoDB struct {
	client   *mongo.Client
	database *mongo.Database
}

func NewMongoDB(cfg MongoConfig) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()

	//creating client
	clientOptions := options.Client().ApplyURI(cfg.URI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	//check connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	//access to db
	database := client.Database(cfg.Database)

	return &MongoDB{
		client:   client,
		database: database,
	}, nil
}

func (m *MongoDB) Close(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}
