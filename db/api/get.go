package api

import (
	"fmt"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"

	"github.com/pdxjohnny/s-db/db/get"
	"github.com/pdxjohnny/s/restQuickReply"
)

// GetDoc uses get to retrive a document
func GetDoc(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	collection := r.PathParam("collection")
	doc, err := get.Get(collection, id)
	if err != nil {
		rest.Error(w, "Not Found", 404)
		return
	}
	w.WriteHeader(http.StatusOK)
	if doc == nil {
		w.(http.ResponseWriter).Write(restQuickReply.BlankResponse)
		return
	}
	fmt.Println(doc)
	w.WriteJson(doc)
}
