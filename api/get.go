package api

import (
	"strings"

	"github.com/pdxjohnny/easyreq"

	"github.com/pdxjohnny/s-db/db/variables"
)

// Get retirives a record from the database
func Get(host, token, collection, id string) (*map[string]interface{}, error) {
	path := variables.APIPathGet
	path = strings.Replace(path, ":collection", collection, 1)
	path = strings.Replace(path, ":id", id, 1)
	return easyreq.GenericRequest(host, path, token, nil)
}
