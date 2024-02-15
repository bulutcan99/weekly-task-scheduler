package query

import (
	"context"
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/valueobject"
	mongodb "github.com/bulutcan99/weekly-task-scheduler/internal/infrastructure/storage/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepository struct {
	db   *mongodb.Mongo
	coll *mongo.Collection
}

func NewTaskRepository(db *mongodb.Mongo, collectionName string) *TaskRepository {
	taskCollection := db.Client.Database(db.Database).Collection(collectionName)
	return &TaskRepository{
		db,
		taskCollection,
	}
}

func (t *TaskRepository) CreateOrUpdate(ctx context.Context, filter any, update any) *mongo.SingleResult {
	update = bson.M{
		"$set":         update,
		"$setOnInsert": bson.M{},
	}

	upsertOption := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	res := t.coll.FindOneAndUpdate(ctx, filter, update, upsertOption)
	return res
}

func (t *TaskRepository) GetTasks(ctx context.Context, opt *options.FindOptions) ([]valueobject.Task, error) {
	filter := bson.M{}
	cursor, err := t.coll.Find(ctx, filter, opt)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []valueobject.Task
	cursor.All(ctx, &tasks)

	return tasks, nil
}
