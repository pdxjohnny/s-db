package get

import (
	"errors"

	"github.com/pdxjohnny/s-db/db/shared"
	"github.com/pdxjohnny/s-db/db/variables"

	"gopkg.in/mgo.v2/bson"
)

// Get reteives a record given the id
func Get(collectionName, id string) (*map[string]interface{}, error) {
	connection := shared.Connection()
	if connection == nil {
		return nil, errors.New("No connection to database server")
	}
	var result map[string]interface{}
	collection := connection.DB(variables.DBName).C(collectionName)
	if !bson.IsObjectIdHex(id) {
		return nil, errors.New("Invalid id: " + id)
	}
	err := collection.FindId(bson.ObjectIdHex(id)).One(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
