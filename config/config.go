package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var Conf Config

type Config struct {
	// 钓鱼状态定位
	Wkk               string
	BeginFishLocation []int
	FishLocation      []int
	FishColor         Color
	// 挣扎定位及颜色
	StruggleLocation []int
	StruggleColor    Color
}

func init() {
	viper.SetConfigFile("./config.yaml") // 指定配置文件路径
	viper.AddConfigPath(".")             // 还可以在工作目录中查找配置
	err := viper.ReadInConfig()          // 读取配置信息
	if err != nil {                      // 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 将读取的配置信息保存至全局变量Conf
	if err := viper.UnmarshalKey("fish", &Conf); err != nil {
		panic(fmt.Errorf("Unmarshal conf failed, err:%s \n", err))
	}
}
