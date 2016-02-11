package commands

import (
	"github.com/spf13/cobra"

	"github.com/pdxjohnny/s-db/db"
	"github.com/pdxjohnny/s-db/db/get"
	"github.com/pdxjohnny/s-db/db/save"
)

// Commands are the commands that this service can preform
var Commands = []*cobra.Command{
	&cobra.Command{
		Use:   "get",
		Short: "Get number",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			get.Run()
		},
	},
	&cobra.Command{
		Use:   "save",
		Short: "Save a doc in the database",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			save.Run()
		},
	},
	&cobra.Command{
		Use:   "http",
		Short: "Start the http server",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			db.Run()
		},
	},
}

func init() {
	ConfigDefaults(Commands...)
}
