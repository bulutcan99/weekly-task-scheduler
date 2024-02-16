package query

import (
	"context"

	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/entity"
	mongodb "github.com/bulutcan99/weekly-task-scheduler/internal/infrastructure/storage/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProviderRepository struct {
	db   *mongodb.Mongo
	coll *mongo.Collection
}

func NewProviderRepository(db *mongodb.Mongo, collectionName string) *ProviderRepository {
	providerColl := db.Client.Database(db.Database).Collection(collectionName)
	return &ProviderRepository{
		db,
		providerColl,
	}
}

func (t *ProviderRepository) InsertProvider(ctx context.Context, provider *entity.Provider) error {
	provider.ID = primitive.NewObjectID()
	_, err := t.coll.InsertOne(ctx, provider)
	if err != nil {
		return err
	}
	return nil
}

func (t *ProviderRepository) FindProviders(ctx context.Context) ([]entity.Provider, error) {
	var providers []entity.Provider
	cursor, err := t.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &providers)
	if err != nil {
		return nil, err
	}
	return providers, nil
}

func (t *ProviderRepository) UpdateProvider(ctx context.Context, provider *entity.Provider) error {
	filter := bson.M{
		"_id": provider.ID,
	}
	_, err := t.coll.ReplaceOne(ctx, filter, provider)
	if err != nil {
		return err
	}
	return nil
}

func (t *ProviderRepository) DeleteProviders(ctx context.Context, name string) error {
	filter := bson.M{
		"name": name,
	}
	_, err := t.coll.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (t *ProviderRepository) GetProviderWithTasks(ctx context.Context, id primitive.ObjectID) (*entity.Provider, error) {
	filter := bson.M{
		"_id": id,
	}
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: filter}},
		{{Key: "$lookup", Value: bson.M{
			"from":         "tasks",
			"localField":   "_id",
			"foreignField": "provider_id",
			"as":           "tasks",
		}}},
	}
	cursor, err := t.coll.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var provider []entity.Provider
	err = cursor.All(ctx, &provider)
	if err != nil {
		return nil, err
	}
	return &provider[0], nil
}

func (t *ProviderRepository) GetProvider(ctx context.Context, filter any) (*entity.Provider, error) {
	var provider entity.Provider
	err := t.coll.FindOne(ctx, filter).Decode(&provider)
	if err != nil {
		return nil, err
	}
	return &provider, nil
}
