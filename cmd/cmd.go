package cmd

import (
	"fmt"
	"github.com/Bestfeel/markdown/mdr"
	"github.com/spf13/cobra"
	"os"
)

var (
	globalAddr = ":7070"
	globalPath = "."
	globalCss  = "marxico"
)

func init() {
	RootCmd.Flags().StringVarP(&globalAddr, "addr", "a", ":7070", "server address")
	RootCmd.Flags().StringVarP(&globalPath, "path", "p", ".", "sever path")
	RootCmd.Flags().StringVarP(&globalCss, "css", "c", "marxico", "markdown for css style.example [github|mou|marxico]")

}

var RootCmd = &cobra.Command{
	Use:   "markdown",
	Short: "Powerful markdown online",
	Long:  `This is a powerful online tool about markdown that allows you to run a service online to view markdown files,and can also serve as a static server`,
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		mdr.RunMarkDownServer(globalAddr, globalPath, globalCss)
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
