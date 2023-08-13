package settings

// 使用viper来做配置管理工具

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 全局变量,用来保存程序的所有配置信息
var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"db_name"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

// 使用 viper 管理配置
func InIt() (err error) {
	viper.SetConfigFile("config.yaml") // 指定配置文件(包括文件名与类型)
	// viper.SetConfigName("config") // 指定配置文件名称(不需要带后缀) 有重复名称 不是与下面type搭配使用的
	// viper.SetConfigType("yaml")   // 指定配置文件类型 (专用于从远程配置信息时指定配置)
	viper.AddConfigPath(".")   // 指定查找配置文件路径(这里使用相对路径)
	err = viper.ReadInConfig() // 读取配置信息
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig() failed err: %v\n", err)
		return
	}

	// 把读取到的配置信息反序列化到Conf变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed err: %v\n", err)
	}

	viper.WatchConfig() // 热加载 配置修改自动加载 时刻监控配置文件的变化
	viper.OnConfigChange(
		func(in fsnotify.Event) {
			fmt.Println("配置文件修改了...")
			if err := viper.Unmarshal(Conf); err != nil {
				fmt.Printf("viper.Unmarshal failed, err: %v\n", err)
			}
		},
	)
	return
}
