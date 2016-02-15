package api

import (
	"strings"

	"github.com/pdxjohnny/easyreq"

	"github.com/pdxjohnny/s-db/db/variables"
)

// Index rcreates an index on a collection
func Index(host, token, collection, key, indexType string) (*map[string]interface{}, error) {
	path := variables.APIPathIndex
	path = strings.Replace(path, ":collection", collection, 1)
	path = strings.Replace(path, ":key", key, 1)
	path = strings.Replace(path, ":indexType", indexType, 1)
	return easyreq.GenericRequest(host, path, token, nil)
}
