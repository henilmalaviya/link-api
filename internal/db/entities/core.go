package entities

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Insert[T any](ctx context.Context, collection *mongo.Collection, entity *T) (primitive.ObjectID, error) {
	result, err := collection.InsertOne(ctx, entity)
	if err != nil {
		return primitive.NilObjectID, err
	}

	// Extract and return the inserted ID
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, mongo.ErrInvalidIndexValue
	}

	return insertedID, nil
}

func GetByID[T any](ctx context.Context, collection *mongo.Collection, id string) (*T, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return GetOneWithFilter[T](ctx, collection, map[string]any{"_id": objId})
}

func GetWithFilter[T any](ctx context.Context, collection *mongo.Collection, filter map[string]any) ([]*T, error) {
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []*T
	for cursor.Next(ctx) {
		var elem T
		if err := cursor.Decode(&elem); err != nil {
			return nil, err
		}
		results = append(results, &elem)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

func GetOneWithFilter[T any](ctx context.Context, collection *mongo.Collection, filter map[string]any) (*T, error) {
	var result T
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
