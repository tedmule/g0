package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// var rootCmd = &cobra.Command{
// 	Use:   "ted",
// 	Short: "Ted is a mighty bear",
// 	Long: `Ted: fuck you thunder
// you can suck my dick
// you can't do that cuz you're just a shit`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Println("cobra running")
// 	},
// }

var rootCmd = &cobra.Command{}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Long:  `This is my version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("My version is 1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
