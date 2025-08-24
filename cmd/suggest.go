/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/fatih/color"

	"example.com/cmtr/internal/git"
	"example.com/cmtr/internal/ollama"
	"github.com/spf13/cobra"
)

// suggestCmd represents the suggest command
var suggestCmd = &cobra.Command{
	Use:   "suggest",
	Short: "Suggest a commit message for the current diff",
	Long: `Suggest a commit message for the current diff`,
	Run: func(cmd *cobra.Command, args []string) {
		diff, err := git.GetDiff()
		if err != nil {
			fmt.Println("Error getting diff:", err)
			return
		}

		for {

			message, err := ollama.GetMessage(diff)
			if err != nil {
				fmt.Println("Error getting message:", err)
				return
			}

			color.Cyan("\nCommit message: %s", message)

			fmt.Println("\nWould you like to commit using this message?")
			color.Green("1. Yes")
			color.Yellow("2. Try again")
			color.Red("3. Cancel")

			var choice int
			fmt.Print("\nEnter your choice: ")
			fmt.Scanln(&choice)

			if choice == 1 {
				message, err := git.Commit(message)
				fmt.Println("Committed with message:", message)
				if err != nil {
					fmt.Println("Error committing:", err)
					return
				}
				fmt.Println("Committed with message:", message)
				return
			} else if choice == 2 {
				continue
			} else if choice == 3 {
				return
			}
			
		}


	},
}

func init() {
	rootCmd.AddCommand(suggestCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// suggestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// suggestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
