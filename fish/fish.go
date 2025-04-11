package fish

import (
	"context"
	"fishing/common"
	"fishing/config"
	"fmt"
	"github.com/go-vgo/robotgo"
	"time"
)

type Fish struct {
	keyUpCh        chan bool
	finishCh       chan bool
	pause          int
	struggleColors []config.Color
}

func (fish *Fish) Init() {
	fish.keyUpCh = make(chan bool)
	fish.finishCh = make(chan bool)
	fish.struggleColors = make([]config.Color, 2)
	fish.struggleColors[0] = config.Color{Red: 110, Green: 100, Blue: 100} // 白天
	fish.struggleColors[1] = config.Color{Red: 90, Green: 60, Blue: 60}    // 夜晚
	fish.pause = 1
	go fish.KeyboardSimulation() // 拉或者松
	go fish.fishStruggle()       // 挣扎
	go fish.fishStatusListener() // 控制一次钓鱼的生命周期
}

// 状态判断 控制甩勾
// keyUpCh chan bool
func (fish *Fish) fishStatusListener() {
	// 211 218 215
	fishColor := config.Conf.FishColor // 目标颜色
	// 1220, 224
	fishLocation := config.Conf.FishLocation
	fmt.Println("fishStatusListener: 开始监控钓鱼生命周期！")
	const MAX_COUNT = 5
	c := 0
	for {
		time.Sleep(1 * time.Second)
		if fish.pause == 1 {
			continue
		}

		color := common.GetRGBbyLocation(fishLocation[0], fishLocation[1])
		if !color.Range(fishColor) {
			c++
			if c > MAX_COUNT {
				fish.finishCh <- true
				c = 0
			}
		}
	}
}

// 鱼在挣扎 控制按键
func (fish *Fish) fishStruggle() {
	// 110, 64, 85
	// 1159, 463
	fmt.Println("fishStruggle: 开始监控鱼挣扎情况！！")
	var struggleLocation = config.Conf.StruggleLocation
	color := common.GetRGBbyLocation(struggleLocation[0], struggleLocation[1])
	for {
		time.Sleep(5 * time.Millisecond)

		if fish.pause == 1 { // 关闭
			continue
		}
		
		for i := 0; i < len(fish.struggleColors); i++ {
			c := fish.struggleColors[i]
			if color.Red > c.Red && color.Green < c.Green && color.Blue < c.Blue {
				fmt.Println("fishStruggle：检测到鱼挣扎~")
				fish.keyUpCh <- true
				break
			}
		}
	}
}

// KeyboardSimulation 模拟按键
// 接收到第一个拉杆进入钓鱼循环
func (fish *Fish) KeyboardSimulation() {

	pressState := false

	for {
		if fish.pause == 1 {
			pressState = false
			continue
		}

		//fmt.Println("Key：拉勾")
		if !pressState {
			_ = robotgo.KeyDown("space")
			pressState = true
		}

		select {
		case _, ok := <-fish.keyUpCh:
			if !ok {
				continue
			}
		default:
			continue
		}
		// 50 有点鬼畜 效果还可以
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Key：松手")
		_ = robotgo.KeyUp("space", "up")
		pressState = false
		time.Sleep(100 * time.Millisecond)
	}
}

func (fish *Fish) Fishing() {
	fish.pause = 0

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute) // 函数超时
	defer cancel()

	select {
	case <-ctx.Done():
		fmt.Println("Fishing: 本次钓鱼操作超时!")
	case <-fish.finishCh:
		fmt.Println("Fishing：停止本次钓鱼！")
	}

	fish.pause = 1 // 控制关闭通道 避免 panic: send on closed channel
}

func (fish *Fish) CheckBegin() {
	c := 0
	for {
		var location = config.Conf.BeginFishLocation
		color := common.GetRGBbyLocation(location[0], location[1])
		var targetColor = config.Conf.FishColor
		if color.Range(targetColor) {
			return
		} else {
			c++
			if c > 10 {
				return
			}
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func (fish *Fish) Run(times int) {
	fmt.Printf("开始钓鱼！总次数:%d\n", times)

	for i := 0; i < times; i++ {
		// 甩杆
		fmt.Printf("Run: 第 %d 次甩杆！\n", i)
		fish.CheckBegin()
		_ = robotgo.KeyUp(robotgo.Space)
		fmt.Println("甩杆")
		common.KeyClick(robotgo.Space)

		fmt.Println("甩杆等待")
		time.Sleep(5 * time.Second) // 甩杆等待
		fmt.Println("上钩等待")
		time.Sleep(10 * time.Second) // 上钩等待

		fmt.Println("开始拉！")
		common.KeyClick(robotgo.Space)

		fish.Fishing()
		time.Sleep(5 * time.Second)
	}
}
