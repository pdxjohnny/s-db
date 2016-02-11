package api

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/pdxjohnny/go-json-rest-middleware-jwt"

	dbVariables "github.com/pdxjohnny/s-db/db/variables"
	"github.com/pdxjohnny/s/token"
	"github.com/pdxjohnny/s/variables"
)

// CreateAuthMiddleware creates the middleware for authtication
func CreateAuthMiddleware() (*jwt.Middleware, error) {
	err := token.LoadTokenVerifyKey()
	if err != nil {
		log.Println("Error loading TokenVerifyKey:", err)
	}

	authMiddleware := &jwt.Middleware{
		Realm:            token.AuthRealm,
		SigningAlgorithm: token.SigningAlgorithm,
		VerifyKey:        token.TokenVerifyKey,
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour * 24,
		Authenticator: func(username string, password string) error {
			return errors.New("This message should never be seen")
		},
	}
	return authMiddleware, nil
}

// MakeHandler creates the api request handler
func MakeHandler() *http.Handler {
	api := rest.NewApi()

	// Make sure we want to enable auth
	if variables.EnableAuth != false {
		authMiddleware, err := CreateAuthMiddleware()
		if err != nil {
			panic(err)
		}

		api.Use(&rest.IfMiddleware{
			// Only authenticate non login or register requests
			Condition: func(request *rest.Request) bool {
				return true
			},
			IfTrue: authMiddleware,
		})
	}
	api.Use(rest.DefaultProdStack...)
	router, err := rest.MakeRouter(
		rest.Get(dbVariables.APIPathGetServer, GetDoc),
		rest.Get(dbVariables.APIPathGetSaveServer, GetSaveDoc),
		rest.Post(dbVariables.APIPathSaveServer, PostSaveDoc),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	handler := api.MakeHandler()
	return &handler
}
