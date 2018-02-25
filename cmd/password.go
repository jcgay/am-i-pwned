package cmd

import (
	"fmt"
	"github.com/jcgay/am-i-pwned/pwned"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var logger = log.New(os.Stderr, "", 0)

// passwordCmd represents the password command
var passwordCmd = &cobra.Command{
	Use:   "password [candidate-password]",
	Short: "Find if your password is Pwned",
	Long: `Password will be searched by a partial hash (the 5 first characters). For example:

    am-i-pwned password hello-world

or you can type your password instead of passing it as parameter with:

    am-i-pwned password`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		result, err := pwned.CheckPassword(pwned.SelectPassword(args, os.Stdin))
		if err != nil {
			logger.Fatal(err)
		}
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(passwordCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// passwordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// passwordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
