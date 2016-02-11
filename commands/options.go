package commands

// ConfigOptions is used to set viper defaults
var ConfigOptions = map[string]interface{}{
	"get": map[string]interface{}{
		"collection": map[string]interface{}{
			"value": "",
			"help":  "The collection to look in",
		},
		"id": map[string]interface{}{
			"value": "",
			"help":  "The id's doc to return",
		},
	},
	"save": map[string]interface{}{
		"collection": map[string]interface{}{
			"value": "",
			"help":  "The collection to put in",
		},
	},
	"http": map[string]interface{}{
		"addr": map[string]interface{}{
			"value": "0.0.0.0",
			"help":  "Address to bind to",
		},
		"port": map[string]interface{}{
			"value": 8080,
			"help":  "Port to bind to",
		},
		"cert": map[string]interface{}{
			"value": "",
			"help":  "Certificate for https server, for example - keys/http/cert.pem",
		},
		"key": map[string]interface{}{
			"value": "",
			"help":  "Key for https server, for example - keys/http/key.pem",
		},
	},
}
