package db

import (
	"log"
	"net/http"

	"github.com/spf13/viper"

	"github.com/pdxjohnny/s-db/db/api"
	"github.com/pdxjohnny/s/service"
)

// Run starts the http(s) server for the cli
func Run() {
	mux := http.NewServeMux()
	mux.Handle("/api/", http.StripPrefix("/api", *api.MakeHandler()))
	err := service.ServeMux(
		mux,
		viper.GetString("addr"),
		viper.GetString("port"),
		viper.GetString("cert"),
		viper.GetString("key"),
	)
	if err != nil {
		log.Println(err)
		return
	}
}
