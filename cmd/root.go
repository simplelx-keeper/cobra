package cmd

import "github.com/spf13/cobra"

var rootCmd := &cobra.Command{
	Use: "split_xlsx",
	Short: "",
	TraverseChildren:true,
	Long: ``,
}

func init()  {
	rootCmd.AddCommand(NewSplitCmd())
}

func Execute()  {
	return rootCmd.Execute()
}
