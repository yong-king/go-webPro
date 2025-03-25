package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	//err := dao.InitMySQL()
	//if err != nil {
	//	fmt.Printf("init mysql failed, err:%v\n", err)
	//	return
	//}
	//defer dao.DB.Close()
	//fmt.Println("init mysql success")

	//modules.QueryRowDemo()
	//modules.QueryMuitiRowDemo()
	//modules.InsertRowDemo()
	//modules.PrepareQueryDemo()
	//modules.SqlInjectDemo("yk")
	// sql注入
	//modules.SqlInjectDemo("xxx ' or 1=1 #'")
	//modules.SqlInjectDemo("xxx' union select * from user #")
	// 事务
	//modules.TransactionDemo()

	//r := gin.Default()
	//
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//
	//r.Run(":8080")

	//sqlx
	//err := dao.InitSqlx()
	//if err != nil {
	//	fmt.Printf("sqlx connetct failed, err:%v\n", err)
	//}
	//defer dao.DBs.Close()
	//fmt.Println("sqlx connetct success")

	////modules.QueryRaw()
	//modules.QueryMultiRows()
	//u1 := modules.User1{Name: "dlrb", Age: 18}
	//u2 := modules.User1{Name: "yk", Age: 28}
	//u3 := modules.User1{Name: "zrl", Age: 38}
	//users := []*modules.User1{&u1, &u2, &u3}
	//err = modules.BunchInsertUSer(users)
	//if err != nil {
	//	fmt.Printf("bunchInsertUser failed, err:%v\n", err)
	//	return
	//}
	//fmt.Println("bunchInsertUser success")

	//users := []interface{}{&u1, &u2, &u3}
	//err = modules.BunchInsertUSer2(users)
	//if err != nil {
	//	fmt.Printf("bunch insert failed, err:%v\n", err)
	//	return
	//}
	//fmt.Println("bunch insert success")
	//

	//user := []*modules.User1{&u1, &u2, &u3}
	//err = modules.BatchInsertUsers3(user)
	//if err != nil {
	//	fmt.Printf("BatchInsertUsers3 failed, err:%v\n", err)
	//}
	//fmt.Println("bunch insert success")

	//ids := []int{11, 14, 13, 15}
	//users, err := modules.QueryByIDs(ids)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//for _, user := range users {
	//	fmt.Println(user)
	//}

	//err := dao.InitRedis()
	//if err != nil {
	//	fmt.Printf("redis init failed, err:%v\n", err)
	//	return
	//}
	//fmt.Println("redis init success")
	//defer dao.Rdb.Close()

	//modules.RedisExample()
	//modules.HgetDemo()
	//modules.RedisExample2()
	//err := dao.InitRedis9()
	//if err != nil {
	//	fmt.Printf("init redis failed, err:%v\n", err)
	//	return
	//}
	//fmt.Println("init redis success")
	//defer dao.Rdb9.Close()

	//modules.DoCommand()
	//modules.DoDemo()
	//modules.ScanKEyDemo()
	//modules.PipelineDemo()
	//modules.TransactionRedisDemo()
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	//err = modules.WatchDemo(ctx, "pipeline_counter")
	//if err != nil {
	//	fmt.Printf("watch demo failed, err:%v\n", err)
	//	return
	//}
	//fmt.Println("watch demo success")
	//modules.Example()
	//log.SetupLogger()
	//log.SimpleHttpGet("www.baidu.com")
	//log.SimpleHttpGet("http://www.google.com")
	//log.InitLogger()
	//defer log.Logger.Sync()
	//log.SimpoleHttpGet("www.baidu.com")
	//log.SimpoleHttpGet("http://www.baidu.com")
	//log.SimpleHttpGet1("www.baidu.com")
	//log.SimpleHttpGet1("http://www.baidu.com")

	//router := gin.Default()
	//log.InitLogger()
	//defer log.Logger.Sync()
	//router := gin.New()
	//router.Use(log.GinLogger(log.Logger), log.GinRecovery(log.Logger, true))
	//router.GET("/hello", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "hello world",
	//	})
	//})
	//router.Run(":8080")

	//setting.GetConfig()
	//router := gin.Default()
	//router.GET("/version", func(c *gin.Context) {
	//	c.String(http.StatusOK, viper.GetString("version"))
	//})
	//router.Run(":8080")

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Hello World")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit // 阻塞在此，当接收到上述两种信号时才会往下执行
	log.Println("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
