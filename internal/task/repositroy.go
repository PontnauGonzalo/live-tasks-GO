package task

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"	
)

// data manage and database connection
type Repository struct {
	taskCollection *mongo.Collection
}

func NewTaskRepository(db *mongo.Database) *Repository {
	return &Repository {
		taskCollection: db.Collection("tasks"),
	}
}

func (r *Repository) CreateTask(ctx context.Context, task domain.Task) (domain.Task, error) {
	task.ID = uuid.New().String()
	_, err := r.taskCollection.InsertOne(ctx, task)
	// handle error
	if err != nil {
		return domain.Task{}, err
	}
	return task, nil
}