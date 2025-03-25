package main

import (
	"context"
	"flag"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/logger"
	"web_app/routers"
	"web_app/settings"
)

func main() {

	//if len(os.Args) < 2 {
	//	fmt.Printf("需要输入yaml配置文件位置！")
	//	return
	//}

	//filename := flag.String("fileParth", "./conf/config.yaml", "配置文件地址!")
	var filename string
	flag.StringVar(&filename, "Path", "./config.yaml", "配置文件地址!")
	flag.Parse()

	//1. 加载配置
	err := settings.Init(filename) //settings.Init(os.Args[1])
	if err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
	}

	//2.初始化日志
	err = logger.Init(settings.Conf.LogConfig)
	if err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
	}
	defer zap.L().Sync()

	//3.初始化mysql
	err = mysql.Init(settings.Conf.MySQLConfig)
	if err != nil {
		zap.L().Error("init mysql failed", zap.Error(err))
	}
	defer mysql.Close()

	//4.初始化redis
	err = redis.Init(settings.Conf.RedisConfig)
	if err != nil {
		zap.L().Error("init redis failed", zap.Error(err))
	}
	defer redis.Close()

	//5.注册路由
	r := routers.Setup()

	//6.启动服务(优雅关机)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.Port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Error("listen failed", zap.Error(err))
		}
	}()

	// 优雅关机
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Error("Server Shutdown", zap.Error(err))
	}
	zap.L().Info("Server exiting")
}
