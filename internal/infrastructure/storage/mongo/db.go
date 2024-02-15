package mongodb

import (
	"context"
	"fmt"
	"github.com/bulutcan99/weekly-task-scheduler/internal/infrastructure/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
	"sync"
)

var doOnce sync.Once
var client *mongo.Client

type Mongo struct {
	Client   *mongo.Client
	Context  context.Context
	Database string
}

func NewConnection(ctx context.Context, mongoConfig *config.Mongo) *Mongo {
	url := fmt.Sprintf("mongodb://%s:%d", mongoConfig.Host, mongoConfig.Port)
	dbName := mongoConfig.Name
	doOnce.Do(func() {
		cli, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
		if err != nil {
			panic(err)
		}
		err = cli.Ping(ctx, nil)
		if err != nil {
			panic(err)
		}

		client = cli
	})

	slog.Info("Connected to MongoDB successfully")

	return &Mongo{
		Client:   client,
		Context:  ctx,
		Database: dbName,
	}
}

func (m *Mongo) Close() {
	err := m.Client.Disconnect(m.Context)
	if err != nil {
		slog.Error("Error while disconnecting from MongoDB:", err)
	}
	slog.Info("Connection to MongoDB closed successfully")
}

func (m *Mongo) Stop() {
	mongodbActiveSessionsCount := m.numberSessionsInProgress()
	for mongodbActiveSessionsCount != 0 {
		mongodbActiveSessionsCount = m.numberSessionsInProgress()
	}
}

func (m *Mongo) numberSessionsInProgress() int {
	return m.Client.NumberSessionsInProgress()
}
