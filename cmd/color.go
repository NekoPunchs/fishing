package cmd

import (
	"fishing/config"
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/spf13/cobra"
	"time"
)

var colorCmd = &cobra.Command{
	Use:   "color",
	Short: "Screen color.",
	Long:  `Get the RGB value of the mouse coordinates.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("鼠标定位倒计时：3\n")
		time.Sleep(1 * time.Second)

		fmt.Printf("鼠标定位倒计时：2\n")
		time.Sleep(1 * time.Second)

		fmt.Printf("鼠标定位倒计时：1\n")
		time.Sleep(1 * time.Second)
		x, y := robotgo.Location()
		fmt.Printf("当前鼠标坐标的RGB为：%v", config.HexToRGB(robotgo.GetPixelColor(x, y)))
	},
}

func init() {
	rootCmd.AddCommand(colorCmd)
}
