package cmd

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/spf13/cobra"
	"time"
)

var locationCmd = &cobra.Command{
	Use:   "location",
	Short: "Get mouse coordinates.",
	Long:  `Get mouse coordinates.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("鼠标定位倒计时：3\n")
		time.Sleep(1 * time.Second)

		fmt.Printf("鼠标定位倒计时：2\n")
		time.Sleep(1 * time.Second)

		fmt.Printf("鼠标定位倒计时：1\n")
		time.Sleep(1 * time.Second)
		x, y := robotgo.Location()

		fmt.Printf("当前鼠标的坐标为： x %d, y %d\n", x, y)
	},
}

func init() {
	rootCmd.AddCommand(locationCmd)
}
