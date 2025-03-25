package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Port         string
	Version      string `mapstructure:"version"`
	Mode         string `mapstructure:"mode"`
	MachineID    int64  `mapstructure:"machineID"`
	StartTime    string `mapstructure:"startTime"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
	*LogConfig   `mapstructure:"log"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DB       string `mapstructure:"dbname"`
	MaxOpen  int    `mapstructure:"maxOpen"`
	MaxIdle  int    `mapstructure:"maxIdleCoons"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"poolSize"`
	MaxIdle  int    `mapstructure:"maxIdleCoons"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Path       string `mapstructure:"path"`
	FileName   string `mapstructure:"fileName"`
	MaxAge     int    `mapstructure:"maxAge"`
	MaxBackups int    `mapstructure:"maxBackups"`
	MaxSize    int    `mapstructure:"maxSize"`
}

func Init(fileName string) (err error) {
	// 获取配置文件路径
	viper.SetConfigFile(fileName)
	// 读取配置文件
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig failed, err:%v\n", err)
		return
	}

	// 读取配置文件,反序列化
	err = viper.Unmarshal(Conf)
	if err != nil {
		return
	}

	// 监视
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file changed: %s\n", e.Name)
		err = viper.Unmarshal(Conf)
		if err != nil {
			return
		}
	})
	return
}
