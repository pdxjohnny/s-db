package near

import (
	"errors"
	"fmt"
	"log"

	"github.com/pdxjohnny/s-db/db/shared"
	"github.com/pdxjohnny/s-db/db/variables"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Near reteives records near the given longitude and latitude
func Near(collectionName string, longitude, latitude, radius float64) (*map[string]interface{}, error) {
	connection := shared.Connection()
	if connection == nil {
		return nil, errors.New("No connection to database server")
	}
	collection := connection.DB(variables.DBName).C(collectionName)

	setIndex := fmt.Sprintf("$%s:%s", "2dsphere", "location")

	// Define the index
	index := mgo.Index{
		Key: []string{setIndex},
	}

	// Apply the index
	err := collection.EnsureIndex(index)
	if err != nil {
		log.Println("ERROR creating 2dsphere index", err)
	}

	var found []map[string]interface{}
	err = collection.Find(bson.M{
		"location": bson.M{
			"$nearSphere": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{longitude, latitude},
				},
				"$maxDistance": radius,
			},
		},
	}).All(&found)
	if err != nil {
		return nil, err
	}
	result := map[string]interface{}{
		"rows": found,
	}
	return &result, nil
}
