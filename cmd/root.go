package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
  )
  
//   func init() {
// 	rootCmd.AddCommand(tryCmd)
//   }
  
  var tryCmd = &cobra.Command{
	Use:   "currency",
	Short: "get support currency using ISO 4217 code as argument",
	Long: `currency --code KES `,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// prints the code given
		someFunc(args[0])
		// fmt.Println(args)
		return nil
	  },
  }

  func Execute() {
	if err := tryCmd.Execute(); err != nil {
	  fmt.Println(err)
	  os.Exit(1)
	}
  }

  func someFunc(x string) {
	var resultCurrency = readcsv()

	for _, item := range resultCurrency {
		// fmt.Println(i,item.isocode)
		if item.isocode == x {
			fmt.Println("Exists the currency is supported")
		}
    }
  }