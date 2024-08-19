package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		items, err := fctx.ListFiles()
		if err != nil {
			log.Println(err)
		}
		for _, item := range items {
			fmt.Println(item)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

}
