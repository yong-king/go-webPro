package main

import (
	"bubble/controller"
	"bubble/dao/mysql"
	"bubble/dao/redis"
	"bubble/logger"
	"bubble/pkg/snowflake"
	"bubble/router"
	"bubble/setting"
	"context"
	"flag"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 配置文件路径
	var fileName string
	flag.StringVar(&fileName, "configPath", "./conf/config.yaml", "配置文件路径！")
	flag.Parse()

	// 1.获取配置信息 vipper
	err := setting.Init(fileName)
	if err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
		return
	}
	// 2.设置日志文件
	err = logger.Init(setting.Conf.LogConfig)
	if err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	// 3.mysql初始化
	err = mysql.Init(setting.Conf.MysqlConfig)
	if err != nil {
		zap.L().Error("init mysql failed", zap.Error(err))
		return
	}
	defer mysql.Close()
	// 4.redis初始化
	err = redis.Init(setting.Conf.RedisConfig)
	if err != nil {
		zap.L().Error("init redis failed", zap.Error(err))
		return
	}
	defer redis.Close()
	// 雪花算法生成userid初始化
	snowflake.Init(setting.Conf.MachineID, setting.Conf.StartTime)
	// 初始化gin框架内置的校验器使用的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed, err:%v\n", err)
		return
	}
	// 5.gin框架初始化
	r := router.Setup(setting.Conf.Mode)
	// 6.启动服务
	srv := &http.Server{
		Addr:    setting.Conf.Port,
		Handler: r,
	}
	// 监听
	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Error("listen failed", zap.Error(err))
			return
		}
	}()
	// 7.优雅关机
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = srv.Shutdown(ctx)
	if err != nil {
		zap.L().Error("Server Shutdown Failed", zap.Error(err))
		return
	}
}
