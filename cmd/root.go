package cmd

import (
	"fishing/fish"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var times int
var rootCmd = &cobra.Command{
	Use:   "fishing",
	Short: "Zhu Xian World Fishing Script.",
	Long: `A script that can automatically fish
Let you free your hands and enjoy fishing in the world of Zhu Xian.`,
	Run: func(cmd *cobra.Command, args []string) {
		fish.Run(times)
	},
}

func Execute() {
	rootCmd.PersistentFlags().IntVarP(&times, "times", "t", 50, "fishing times")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
