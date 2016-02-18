package save

import (
	"errors"
	"fmt"
	"log"

	"github.com/pdxjohnny/s-db/db/shared"
	"github.com/pdxjohnny/s-db/db/variables"

	"gopkg.in/mgo.v2/bson"
)

// Save tries to insert then tries to save
func Save(collectionName string, doc interface{}) (interface{}, error) {
	connection := shared.Connection()
	if connection == nil {
		log.Println("MongoConn", connection)
		return nil, errors.New("No connection to database server")
	}
	collection := connection.DB(variables.DBName).C(collectionName)

	var asMap map[string]interface{}
	switch value := doc.(type) {
	case *map[string]interface{}:
		asMap = *(value)
	case map[string]interface{}:
		asMap = value
	default:
		return nil, errors.New("Must be a map")
	}

	id, ok := asMap["_id"]
	if !ok {
		id = bson.NewObjectId()
		asMap["_id"] = id
	} else {
		asMap["_id"] = bson.ObjectIdHex(id.(string))
	}
	fmt.Println("Saving", collectionName, doc)
	return collection.UpsertId(asMap["_id"], doc)
}

// Insert creates a document
func Insert(collectionName string, doc interface{}) error {
	connection := shared.Connection()
	if connection == nil {
		log.Println("MongoConn", connection)
		return errors.New("No connection to database server")
	}
	collection := connection.DB(variables.DBName).C(collectionName)
	err := collection.Insert(doc)
	if err != nil {
		return err
	}

	return nil
}

// Update updates a document
func Update(collectionName string, doc interface{}) error {
	connection := shared.Connection()
	if connection == nil {
		log.Println("MongoConn", connection)
		return errors.New("No connection to database server")
	}
	var asMap map[string]interface{}
	switch value := doc.(type) {
	case *map[string]interface{}:
		asMap = *(value)
	case map[string]interface{}:
		asMap = value
	default:
		return errors.New("Must be a map")
	}
	_, ok := asMap["_id"]
	if !ok {
		return errors.New("Doc needs to have _id to be saved")
	}
	findDoc := bson.M{"_id": asMap["_id"].(string)}

	collection := connection.DB(variables.DBName).C(collectionName)
	err := collection.Update(findDoc, doc)
	if err != nil {
		return err
	}

	return nil
}
