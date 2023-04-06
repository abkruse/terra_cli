package cmd

import (
  "os"

  "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use: "terra_cli"
  Short: "Write terraform from default files"
}

func Execute() {
  err:= writeCmd.Execute()
  if err != nil {
    os.Exit(1)
  }
}

func init() {
  rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
  rootCmd.AddCommand(writeCmd)
}
