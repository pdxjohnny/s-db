package near

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Run is the function to be run for the cli
func Run() {
	data, err := Near(
		viper.GetString("collection"),
		viper.GetFloat64("longitude"),
		viper.GetFloat64("latitude"),
		viper.GetFloat64("radius"),
	)
	if err != nil {
		log.Println("ERROR getting near", err)
		return
	}
	dump, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Println("ERROR dumping near", err)
		return
	}
	fmt.Println(string(dump))
}
