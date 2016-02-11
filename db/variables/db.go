package variables

import (
	"os"
)

var (
	// DBName is the name of the database in the database
	// Change this to access another database
	DBName = "s-db"
	// EnvDBAddress is the env variable that the database address will be in
	// Change this if it is not the env variable that points to your mongo server
	EnvDBAddress = "MONGO_PORT_27017_TCP_ADDR"
	// DBAddress is the address of database server it is set by getting EnvDBAddress
	DBAddress string
)

func init() {
	DBAddress = os.Getenv(EnvDBAddress)
}
