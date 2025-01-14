package fish

import (
	"context"
	"fishing/config"
	"fmt"
	"github.com/go-vgo/robotgo"
	"time"
)

// 状态判断 控制甩勾
// sCh chan bool
func fishStatus(over *int, fCh chan bool) {
	// 211 218 215
	var fish = config.Conf.FishColor // 目标颜色
	// 1220, 224
	var fishLocation = config.Conf.FishLocation

	var c int
	fmt.Println("fishStatus: 开始监控钓鱼生命周期！")

	defer func() {
		fmt.Printf("fishStatus：over, fCh通道关闭~\n")
		close(fCh)
	}()

	for {
		color := getRGBbyLocation(fishLocation[0], fishLocation[1])
		// fmt.Printf("fishStatus: %v\n", color)
		if color.Range(fish) {
			fmt.Println("fishStatus：拉杆中~")
		} else {
			fmt.Printf("fishStatus：停止拉杆~ %d\n", c)
			c++
			if c > 3 {
				fCh <- true
				return
			}

		}

		if *over == 1 { // 关闭
			return
		}

		time.Sleep(1 * time.Second)
	}
}

// 鱼在挣扎 控制按键
func fishStruggle(over *int, sCh chan bool) {
	// 110, 64, 85
	// 1159, 463
	fmt.Println("fishStruggle: 开始监控鱼挣扎情况！！")
	var fish = config.Conf.StruggleColor
	var fishLocation = config.Conf.StruggleLocation

	oldColor := getRGBbyLocation(fishLocation[0], fishLocation[1]) // 初始值
	for {
		time.Sleep(10 * time.Millisecond)
		color := getRGBbyLocation(fishLocation[0], fishLocation[1])
		if color.Range(fish) { // 颜色范围
			fmt.Println("fishStruggle：检测三元色差 鱼挣扎~")
			sCh <- true
		}

		if abs(oldColor.Red-color.Red) > 15 { // R的差值
			fmt.Println("fishStruggle：检测红元色差 鱼挣扎~")
			sCh <- true
		}

		if *over == 1 { // 关闭
			fmt.Printf("fishStruggle：over, sCh通道关闭~\n")
			close(sCh)
			return
		}
	}
}

// KeyboardSimulation 模拟按键
// 接收到第一个拉杆进入钓鱼循环
func KeyboardSimulation(sCH chan bool) {
	for {
		fmt.Println("Key：拉勾")

		_ = robotgo.KeyToggle("space", "down")
		v, ok := <-sCH
		if !ok {
			break
		}
		println(v)

		// 50 有点鬼畜 效果还可以
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Key：松手")
		_ = robotgo.KeyToggle("space", "up")
		time.Sleep(100 * time.Millisecond)
	}

}

func Fishing() {
	var over int
	sCh := make(chan bool, 2)
	fCh := make(chan bool)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute) // 函数超时
	defer cancel()

	go KeyboardSimulation(sCh)  // 拉或者松
	go fishStruggle(&over, sCh) // 挣扎
	go fishStatus(&over, fCh)   // 控制一次钓鱼的生命周期

	select {
	case <-ctx.Done():
		fmt.Println("Fishing: 本次钓鱼操作超时!")
	case <-fCh:
		fmt.Println("Fishing：停止本次钓鱼！")
	}

	over = 1 // 控制关闭通道 避免 panic: send on closed channel
}

func Run(times int) {
	time.Sleep(5 * time.Second)
	fmt.Println("开始钓鱼！")

	for i := 0; i < times; i++ {
		// 甩杆
		fmt.Printf("Run: 第 %d 次甩杆！\n", i)

		_ = robotgo.KeyToggle("space", "down")
		time.Sleep(500 * time.Millisecond)
		_ = robotgo.KeyToggle("space", "up")

		_ = robotgo.KeyToggle("space", "down")
		time.Sleep(500 * time.Millisecond)
		_ = robotgo.KeyToggle("space", "up")

		time.Sleep(20 * time.Second)
		fmt.Println("开始拉！")
		_ = robotgo.KeyToggle("space", "down")
		time.Sleep(200 * time.Millisecond)
		_ = robotgo.KeyToggle("space", "up")

		Fishing()
		time.Sleep(10 * time.Second)
	}
}

// 获取指定位置RGB
func getRGBbyLocation(x, y int) config.Color {
	return config.HexToRGB(robotgo.GetPixelColor(x, y))
}

// 绝对值
func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
