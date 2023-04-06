package write

import (
  "os"

  "github.com/spf13/cobra"
)

var writeCmd = &cobra.Command{
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
  rootCmd.AddCommand(writeCmd)

  writeCmd.Flags().StringVarP(&name, "name", "n", "", "Project name")
  writeCmd.MarkFlagRequired("name")

  writeCmd.Flags().StringVarP(&provider, "provider", "p", "", "Cloud provider")
  writeCmd.MarkFlagRequired("provider")

  writeCmd.Flags().StringVarP(&region, "region", "r", "", "Cloud region")
  writeCmd.MarkFlagRequired("region")

  writeCmd.Flags().StringVarP(&backend, "backend", "b", "", "Project backend")
  writeCmd.MarkFlagRequired("backend")
}

func writeProviderFile(name string, provider string, region string, backend string){
    // validate provider and region

    // call function to write provider file
}
