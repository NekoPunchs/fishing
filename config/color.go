package config

import (
	"fmt"
	"strconv"
)

type Color struct {
	Red   int
	Green int
	Blue  int
}

func InRange(value int, min int, max int) bool {
	return value >= min && value <= max
}

func Abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

// Range
// Color 与 oC 的RGB相差上下不超过5时返回 Ture
func (c Color) Range(oC Color) bool {
	deviation := 50
	r := Abs(c.Red - oC.Red)
	g := Abs(c.Green - oC.Green)
	b := Abs(c.Blue - oC.Blue)
	return InRange(r+g+b, 0, deviation)
}

func HexToRGB(color string) Color {
	rgbVal, err := strconv.ParseUint(color, 16, 32)
	if err != nil {
		fmt.Println("Error parsing color value:", err)
		return Color{}
	}

	r := int((rgbVal >> 16) & 0xFF)
	g := int((rgbVal >> 8) & 0xFF)
	b := int(rgbVal & 0xFF)

	return Color{r, g, b}
}
