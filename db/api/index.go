package api

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"

	"github.com/pdxjohnny/s-db/db/index"
	"github.com/pdxjohnny/s/restQuickReply"
)

// GetIndex creates an index on a key
func GetIndex(w rest.ResponseWriter, r *rest.Request) {
	collection := r.PathParam("collection")
	key := r.PathParam("key")
	indexType := r.PathParam("indexType")
	err := index.Index(collection, key, indexType)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.(http.ResponseWriter).Write(restQuickReply.OKResponse)
}
