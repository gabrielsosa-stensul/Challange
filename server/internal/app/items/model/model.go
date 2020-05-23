package model

import (
	"errors"
	"log"

	"github.com/MarianoArias/challange-api/internal/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Item struct {
	ID          int    `json:"id" bson:"id,omitempty"`
	Image       string `json:"image" bson:"image,omitempty"`
	Description string `json:"description" bson:"description,omitempty"`
	Order       int    `json:"order" bson:"order,omitempty"`
}

var collection *mongo.Collection

func init() {
	collection = mongodb.GetDatabase().Collection("items")
}

func FindAll() (*[]Item, error) {
	var items []Item

	filter := bson.D{}

	options := options.Find()
	options.SetSort(bson.D{{"order", 1}})

	cur, err := collection.Find(nil, filter, options)

	if err != nil {
		return nil, err
	}

	for cur.Next(nil) {
		var item Item
		err := cur.Decode(&item)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return &items, nil
}

func Find(id int) (*Item, error) {
	var item Item

	filter := bson.D{{"id", id}}

	err := collection.FindOne(nil, filter).Decode(&item)

	if err != nil {
		return nil, err
	}

	return &item, nil
}

func Persist(item *Item) error {
	item.ID = getNextId()
	item.Order = getNextOrder()

	_, err := collection.InsertOne(nil, item)

	if err != nil {
		return err
	}

	return nil
}

func Update(item *Item) error {
	filter := bson.D{{"id", item.ID}}

	document := bson.D{
		{"$set", bson.D{
			{"image", item.Image},
			{"description", item.Description},
			{"order", item.Order},
		}},
	}

	res, err := collection.UpdateOne(nil, filter, document)

	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		errors.New("No document matched")
	}

	return nil
}

func Delete(item Item) error {
	filter := bson.D{{"id", item.ID}}

	res, err := collection.DeleteOne(nil, filter)

	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		errors.New("No document matched")
	}

	// REORDER AFTER DELETE //
	items, err := FindAll()
	if err != nil {
		return err
	}

	for key, item := range *items {
		item.Order = key + 1
		err = Update(&item)
		if err != nil {
			return err
		}
	}
	// REORDER AFTER DELETE //

	return nil
}

func SwitchOrder(from int, to int) error {
	items, err := FindAll()
	if err != nil {
		return err
	}

	for _, item := range *items {
		if item.Order == from {
			item.Order = to
			err = Update(&item)
		} else if from > to && item.Order >= to && item.Order < from {
			item.Order++
			err = Update(&item)
		} else if from < to && item.Order <= to && item.Order > from {
			item.Order--
			err = Update(&item)
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func getNextId() int {
	var item Item

	options := options.Find()
	options.SetSort(bson.D{{"id", -1}})
	options.SetLimit(1)

	cur, err := collection.Find(nil, bson.D{}, options)
	if err != nil {
		return 1
	}

	for cur.Next(nil) {
		err := cur.Decode(&item)
		if err != nil {
			log.Fatal(err)
		}
		return item.ID + 1
	}

	return 1
}

func getNextOrder() int {
	var item Item

	options := options.Find()
	options.SetSort(bson.D{{"order", -1}})
	options.SetLimit(1)

	cur, err := collection.Find(nil, bson.D{}, options)
	if err != nil {
		return 1
	}

	for cur.Next(nil) {
		err := cur.Decode(&item)
		if err != nil {
			log.Fatal(err)
		}
		return item.Order + 1
	}

	return 1
}
