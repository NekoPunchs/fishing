package common

import (
	"fishing/config"
	"github.com/go-vgo/robotgo"
	"time"
)

func KeyClick(key string) {
	_ = robotgo.KeyToggle(key, "down")
	time.Sleep(200 * time.Millisecond)
	_ = robotgo.KeyToggle(key, "up")
}

// GetRGBbyLocation 获取指定位置RGB
func GetRGBbyLocation(x, y int) config.Color {
	return config.HexToRGB(robotgo.GetPixelColor(x, y))
}

// 绝对值
func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
