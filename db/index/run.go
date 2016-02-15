package index

import (
	"log"

	"github.com/spf13/viper"
)

// Run is the function to be run for the cli
func Run() {
	err := Index(
		viper.GetString("collection"),
		viper.GetString("key"),
		viper.GetString("indexType"),
	)
	if err != nil {
		log.Println("ERROR applying index", err)
		return
	}
}
