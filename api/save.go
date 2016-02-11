package api

import (
	"strings"

	"github.com/pdxjohnny/easyreq"

	"github.com/pdxjohnny/s-db/db/variables"
)

// GetSave saves a record to the database
func GetSave(host, token, collection, id, value string) (*map[string]interface{}, error) {
	path := variables.APIPathGetSave
	path = strings.Replace(path, ":collection", collection, 1)
	path = strings.Replace(path, ":id", id, 1)
	path = strings.Replace(path, ":value", value, 1)
	return easyreq.GenericRequest(host, path, token, nil)
}

// Save saves a record to the database
func Save(host, token, collection, id string, doc map[string]interface{}) (*map[string]interface{}, error) {
	path := variables.APIPathSave
	path = strings.Replace(path, ":collection", collection, 1)
	path = strings.Replace(path, ":id", id, 1)
	return easyreq.GenericRequest(host, path, token, doc)
}
