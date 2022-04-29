/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "FolderFlooder",
	Short: "It's a lot of folders!",
	Long:  `Flood your directories with directories, why? For fun!`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		amount, _ := cmd.Flags().GetInt("amount")
		index, _ := cmd.Flags().GetBool("indexed")
		c, _ := cmd.Flags().GetBool("caution")

		mk(amount, index, c)
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

func init() {
	rootCmd.Flags().BoolP("caution", "c", false, "If you want to make more than 50 folders at a time")
	rootCmd.Flags().BoolP("indexed", "i", true, "If you want to name your folders by numbering them")
	rootCmd.Flags().IntP("amount", "a", 10, "The amount of folders you want")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func mk(amount int, numbered, cautious bool) {
	if cautious == false {
		if amount > 50 {
			fmt.Println("In order that you dont acidentally flood the wrong folder with too many folders, you cannot make more than 50 at a time. use the -c flag to avoid this, defaulting to 10")
			amount = 10
		}
	}
	for i := 0; i < amount; i++ {
		var name string
		if numbered == true {
			name = fmt.Sprint(i + 1)
		} else {
			name = fmt.Sprintf("%d", time.Now().UnixMilli())
		}
		os.Mkdir(name, os.ModeDir)
		time.Sleep(10 * time.Millisecond)
	}
}
