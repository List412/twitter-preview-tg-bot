package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func New(db *mongo.Collection) Repository {
	return Repository{db: db}
}

type Repository struct {
	db *mongo.Collection
}

func (r Repository) Count(ctx context.Context) (int64, error) {
	return r.db.CountDocuments(ctx, bson.D{})
}

func (r Repository) IsExist(ctx context.Context, userName string) (bool, error) {
	count, err := r.db.CountDocuments(ctx, bson.D{{Key: "userName", Value: userName}})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r Repository) Add(ctx context.Context, userName string) error {
	_, err := r.db.InsertOne(ctx, bson.D{{Key: "userName", Value: userName}})
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) Delete(ctx context.Context, userName string) error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) Get(ctx context.Context, userName string) error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) All(ctx context.Context, limit int, offset int) error {
	//TODO implement me
	panic("implement me")
}
