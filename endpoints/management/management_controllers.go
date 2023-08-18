package management

import (
	"brms/endpoints/models"
	"brms/pkg/db"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertOneRule(ruleSet models.RuleSet) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, collectionName, err := db.ConnectDB("rule_engine")
	if err != nil {
		return "", err
	}
	defer client.Disconnect(ctx)

	// check unique
	countRule, err := collectionName.CountDocuments(ctx, bson.M{"name": ruleSet.Name})
	if err != nil {
		return "", err
	}
	if countRule > 0 {
		return "", fmt.Errorf("rule set already exists")
	}

	// insert to mongo
	result, err := collectionName.InsertOne(ctx, ruleSet)
	if err != nil {
		return "", err
	}

	resultID, _ := result.InsertedID.(primitive.ObjectID)

	return resultID.String(), nil
}

func FetchAllRules() ([]models.RuleSet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, collectionName, err := db.ConnectDB("rule_engine")
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)

	var results []models.RuleSet

	cursor, err := collectionName.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}
