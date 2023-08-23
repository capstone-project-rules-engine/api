package management

import (
	"brms/endpoints/models"
	"brms/pkg/db"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func InsertRulestoRuleSet(ruleSetName string, newRules []models.Rule) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, collectionName, err := db.ConnectDB("rule_engine")
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)

	// insert rules
	filterRuleSet := bson.M{
		"endpoint": ruleSetName,
	}

	exist, err := collectionName.CountDocuments(ctx, filterRuleSet)
	if err != nil{
		return err
	}
	// check if rule set exists
	if exist == 0{
		return fmt.Errorf("rule set does not exists")
	}

	// count rules in document
	filterDocument := bson.M{
		"endpoint": ruleSetName,
		"rules": bson.M{
			"$exists": true,
			"$ne":     nil,
		},
	}
	count, err := collectionName.CountDocuments(ctx, filterDocument)
	if err != nil {
		return err
	}

	var filterUpdate bson.M

	if count == 0 { // case where no rules in specific document
		// set filter when it's null
		filterUpdate = bson.M{
			"$set": bson.M{
				"rules": newRules,
			},
		}
	} else { // case where rules already exists
		// insert when rules already exists
		filterUpdate = bson.M{
			"$push": bson.M{
				"rules": bson.M{
					"$each": newRules,
				},
			},
		}
	}
	
	if _, err := collectionName.UpdateOne(ctx, filterRuleSet, filterUpdate); err != nil {
		return err
	}

	return nil
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

func findRuleSetByName(ruleSetName string) (*models.RuleSet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, collectionName, err := db.ConnectDB("rule_engine")
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)

	var ruleSet models.RuleSet

	filter := bson.M{"endpoint": ruleSetName}

	err = collectionName.FindOne(ctx, filter).Decode(&ruleSet)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("rule does not exists")
		}
		return nil, err
	}

	return &ruleSet, nil
}

func UpdateRuleSet(ruleSetName string, updatedRuleSet models.RuleSet) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, collectionName, err := db.ConnectDB("rule_engine")
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)

	filter := bson.M{
		"endpoint": ruleSetName,
	}
	update := bson.M{
		"$set": updatedRuleSet,
	}

	result, err := collectionName.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("no data found")
	}

	return nil
}

func deleteRuleSet(ruleSetName string) error{
	ctx, cancel := context. WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, collectionName, err := db.ConnectDB("rule_engine")
	if err != nil{
		return err
	}
	defer client.Disconnect(ctx)

	filter := bson.M{
		"endpoint": ruleSetName,
	}

	result, err := collectionName.DeleteOne(ctx, filter)
	if err != nil{
		return err
	}
	if result.DeletedCount == 0{
		return fmt.Errorf("no data exists to be deleted")
	}

	return nil
}
