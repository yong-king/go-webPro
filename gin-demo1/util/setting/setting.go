package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Port        int `mapstructure:"port"`
	Vresion     int `mapstructure:"vresion"`
	MysqlConfig `mapstructure:"mysql"`
}

type MysqlConfig struct {
	Host   string `mapstructure:"host"`
	Dbname string `mapstructure:"dbname"`
	Port   int    `mapstructure:"port"`
}

func GetConfig() {
	// 设置默认值，某个配置项没有被用户定义，才会使用默认值。
	viper.SetDefault("ContentDir", "content")

	// 读取配置文件
	// 确定要去哪里去取配置文件
	viper.SetConfigFile("./config/config.yaml")

	// 去其他文件中找适配的配置文件
	//viper.SetConfigName("config") // 配置文件名称(无扩展名)
	//viper.SetConfigType("yaml")   // 如果配置文件的名称中没有扩展名，则需要配置此项
	//viper.AddConfigPath("/etc/appname/")  	// 会按照这些顺序去寻找配置文件
	//viper.AddConfigPath("$HOME/.appname")
	//viper.AddConfigPath("./config/")

	// 查找读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("No config file found")
		} else {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}
	//if err != nil {
	//	panic(fmt.Errorf("Fatal error config file: %s \n", err))
	//}
	//fmt.Println("success!")

	// 实时监控配置文件的变化
	viper.WatchConfig()
	// 当配置变化之后调用的一个回调函数
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
		return
	}
	fmt.Printf("%+v", c)
}
