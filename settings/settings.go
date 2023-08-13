package settings

// 使用viper来做配置管理工具

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 使用 viper 管理配置
func InIt() (err error) {
	viper.SetConfigName("config") // 指定配置文件名称(不需要带后缀)
	viper.SetConfigType("yaml")   // 指定配置文件类型
	viper.AddConfigPath(".")      // 指定查找配置文件路径(这里使用相对路径)
	err = viper.ReadInConfig()    // 读取配置信息
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig() failed err: %v\n", err)
		return
	}
	viper.WatchConfig() // 热加载 配置修改自动加载
	viper.OnConfigChange(
		func(in fsnotify.Event) {
			fmt.Println("配置文件修改了")
		},
	)
	return
}
