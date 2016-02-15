package variables

const (
	// APIPathGetServer is the path under API
	APIPathGetServer = "/:collection/:id"
	// APIPathGet retrives a record based on a key
	APIPathGet = "/api" + APIPathGetServer
	// APIPathGetSaveServer is the path under API
	APIPathGetSaveServer = "/:collection/:id/:value"
	// APIPathGetSave makes a get request to set a key to a value
	APIPathGetSave = "/api" + APIPathGetSaveServer
	// APIPathSaveServer is the path under API
	APIPathSaveServer = "/:collection/:id"
	// APIPathSave makes a post request to save a document
	APIPathSave = "/api" + APIPathSaveServer
	// APIPathIndexServer is the path under API
	APIPathIndexServer = "/index/:collection/:key/:indexType"
	// APIPathIndex creates an index on a collection
	APIPathIndex = "/api" + APIPathIndexServer
	// APIPathNearServer is the path under API
	APIPathNearServer = "/near/:collection"
	// APIPathNear makes a geospacial query
	APIPathNear = "/api" + APIPathNearServer
)
