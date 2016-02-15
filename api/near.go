package api

import (
	"strings"

	"github.com/pdxjohnny/easyreq"

	"github.com/pdxjohnny/s-db/db/variables"
)

// Near preforms a geospacial query to find things near a point
func Near(host, token, collection string, doc map[string]interface{}) (*map[string]interface{}, error) {
	path := variables.APIPathNear
	path = strings.Replace(path, ":collection", collection, 1)
	return easyreq.GenericRequest(host, path, token, doc)
}
