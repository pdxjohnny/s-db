package index

import (
	"errors"
	"fmt"

	"github.com/pdxjohnny/s-db/db/shared"
	"github.com/pdxjohnny/s-db/db/variables"

	"gopkg.in/mgo.v2"
)

// Index creates an index on an a collection
func Index(collectionName, key, indexType string) error {
	connection := shared.Connection()
	if connection == nil {
		return errors.New("No connection to database server")
	}
	collection := connection.DB(variables.DBName).C(collectionName)

	setIndex := fmt.Sprintf("$%s:%s", indexType, key)

	// Define the index
	index := mgo.Index{
		Key: []string{setIndex},
	}

	// Apply the index
	return collection.EnsureIndex(index)
}
