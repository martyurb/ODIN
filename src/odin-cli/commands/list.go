package commands

import (
    "fmt"
    "github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
    Use:   "list",
    Short: "lists the user's current Odin jobs",
    Long:  `This subcommand lists the user's current Odin jobs`,
    Run: func(cmd *cobra.Command, args []string) {
            listJob()
    },
}

func init() {
    RootCmd.AddCommand(ListCmd)
}

func listJob() {
    fmt.Println("list job")
}
