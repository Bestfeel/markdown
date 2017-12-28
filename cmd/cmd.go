package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/Bestfeel/markdown/markdown"
)

var (
	globalAddr = ":7070"
	globalPath = "."
	globalCss  = "mou"
)

func init() {
	RootCmd.Flags().StringVarP(&globalAddr, "addr", "a", ":7070", "server address")
	RootCmd.Flags().StringVarP(&globalPath, "path", "p", ".", "sever path")
	RootCmd.Flags().StringVarP(&globalCss, "css", "c", "mou", "markdown for css style.example [github|mou|marxico]")

}

var RootCmd = &cobra.Command{
	Use:   "markdown",
	Short: "Powerful markdown online",
	Long:  `This is a powerful online tool about markdown that allows you to run a service online to view markdown files,and can also serve as a static server`,
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		markdown.RunMarkDownServer(globalAddr, globalPath, globalCss)
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
