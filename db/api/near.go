package api

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"

	"github.com/pdxjohnny/s-db/db/near"
	"github.com/pdxjohnny/s/restQuickReply"
)

// PostNear retives locations near the point specified
func PostNear(w rest.ResponseWriter, r *rest.Request) {
	collection := r.PathParam("collection")

	var recvDoc map[string]interface{}
	err := r.DecodeJsonPayload(&recvDoc)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var longitude float64
	var latitude float64
	var radius float64
	var ok bool
	if longitude, ok = recvDoc["longitude"].(float64); !ok {
		rest.Error(w, "longitude needs to be a float", http.StatusInternalServerError)
		return
	}
	if latitude, ok = recvDoc["latitude"].(float64); !ok {
		rest.Error(w, "latitude needs to be a float", http.StatusInternalServerError)
		return
	}
	if radius, ok = recvDoc["radius"].(float64); !ok {
		rest.Error(w, "radius needs to be a float", http.StatusInternalServerError)
		return
	}

	doc, err := near.Near(collection, longitude, latitude, radius)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	if doc == nil {
		w.(http.ResponseWriter).Write(restQuickReply.BlankResponse)
		return
	}
	w.WriteJson(doc)
}
