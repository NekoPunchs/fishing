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

// Range
// Color 与 oC 的RGB相差上下不超过5时返回 Ture
func (c Color) Range(oC Color) bool {
	deviation := 20
	// fmt.Println(c, oC) // 调试
	if oC.Red-deviation < c.Red && c.Red < oC.Red+deviation {
		if oC.Green-deviation < c.Green && c.Green < oC.Green+deviation {
			if oC.Blue-deviation < c.Blue && c.Blue < oC.Blue+deviation {
				return true
			}
		}
	}
	return false
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
