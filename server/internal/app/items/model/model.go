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

// init sets up the collection in mongodb's database.
func init() {
	collection = mongodb.GetDatabase().Collection("items")
}

// FindAll returns all the items from the database or an error otherwise.
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

// Find returns an item from the database or an error otherwise
func Find(id int) (*Item, error) {
	var item Item

	filter := bson.D{{"id", id}}

	err := collection.FindOne(nil, filter).Decode(&item)

	if err != nil {
		return nil, err
	}

	return &item, nil
}

// Persist persists an item in the database.
// It returns write error encountered.
func Persist(item *Item) error {
	item.ID = getNextId()
	item.Order = getNextOrder()

	_, err := collection.InsertOne(nil, item)

	if err != nil {
		return err
	}

	return nil
}

// Update updates an item in the database.
// It returns write error encountered.
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

// Delete deletes an item from the database.
// It returns write error encountered.
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

// SwitchOrder changes the order of an item in the database and also of those 
// affected in the middle.
// It returns write error encountered.
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

// getNextId returns the number of the next id to use in the database.
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

// getNextOrder return the number of the next order to use in the database.
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
