package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"

	"github.com/pdxjohnny/s-db/db/save"
	"github.com/pdxjohnny/s/restQuickReply"
)

// GetSaveDoc uses save to save a document
func GetSaveDoc(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	value := r.PathParam("value")
	collection := r.PathParam("collection")
	_, ok := r.Env["JWT_PAYLOAD"]
	// Make sure that the user is allowed to save this
	if ok {
		// Now we know that auth is enabled so
		_, ok := r.Env["JWT_PAYLOAD"].(map[string]interface{})["backend"]
		if !ok {
			// FIXME: Add call to access control service
			rest.Error(w, "Only backend services can call this service", http.StatusInternalServerError)
			return
		}
	}
	doc := map[string]interface{}{
		"_id": id,
		id:    value,
	}
	fmt.Println(doc)
	err := save.Save(collection, doc)
	if err != nil {
		log.Println(err)
		rest.Error(w, "Could not save", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.(http.ResponseWriter).Write(restQuickReply.BlankResponse)
}

// PostSaveDoc uses save to save a document
func PostSaveDoc(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	collection := r.PathParam("collection")
	// Make sure that the user owns what they are trying to save
	_, ok := r.Env["JWT_PAYLOAD"].(map[string]interface{})["backend"]
	if !ok && r.Env["REMOTE_USER"] != id {
		rest.Error(w, "Cannot modify another users data", http.StatusInternalServerError)
		return
	}
	var doc map[string]interface{}
	err := r.DecodeJsonPayload(&doc)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(doc)
	err = save.Save(collection, doc)
	if err != nil {
		log.Println(err)
		rest.Error(w, "Could not save", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.(http.ResponseWriter).Write(restQuickReply.BlankResponse)
}
