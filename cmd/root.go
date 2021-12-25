package cmd

import (
	"fmt"
	"os"

	"github.com/futuretea/go-expand"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ep pattern",
	Short: "A patterns expander",
	Long:  `A patterns expander`,
	Example: `
eg1:
$ ep 192.168.5.[1-3,8]

192.168.5.1
192.168.5.2
192.168.5.3
192.168.5.8

eg2:
$ ep foo-[a,b,c]-[1,2,3]
foo-a-1
foo-b-1
foo-c-1
foo-a-2
foo-b-2
foo-c-2
foo-a-3
foo-b-3
foo-c-3

`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		expander := expand.NewExpander()
		result, err := expander.Expand(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		for _, s := range result {
			fmt.Println(s)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
