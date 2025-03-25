package log

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

var Logger *zap.Logger
var SugarLogger *zap.SugaredLogger

// Encoder:编码器(如何写入日志)
func getEncoder() zapcore.Encoder {
	//return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	// 普通格式
	//encodeConfig := zap.NewProductionEncoderConfig()
	////修改时间编码器
	////在日志文件中使用大写字母记录日志级别
	//encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encodeConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	return zapcore.NewConsoleEncoder(encodeConfig)

}

// WriterSyncer ：指定日志将写到哪里去。
func getWriterSyncer() zapcore.WriteSyncer {
	//file, _ := os.OpenFile("./test.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0744)
	// 利用io.MultiWriter支持文件和终端两个输出目标
	//ws := io.MultiWriter(file, os.Stdout)
	//return zapcore.AddSync(file)
	lumberjack := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    10,    // M 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 5,     // 保留旧文件的最大个数
		MaxAge:     30,    // 保留旧文件的最大天数
		Compress:   false, // 是否压缩/归档旧文件
	}
	//return zapcore.AddSync(ws)
	return zapcore.AddSync(lumberjack)
}

// test.err.log记录ERROR级别的日志
func getErrorWriterSyncer() zapcore.WriteSyncer {
	fileError, _ := os.OpenFile("./test.err.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0744)
	return zapcore.AddSync(fileError)
}

// Log Level：哪种级别的日志将被写入。
func InitLogger() {
	writeSyncer := getWriterSyncer()
	encoder := getEncoder()
	errorWriterSyncer := getErrorWriterSyncer()
	core1 := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	// 记录ERROR级别的日志
	core2 := zapcore.NewCore(encoder, errorWriterSyncer, zapcore.ErrorLevel)
	// 使用NewTee将c1和c2合并到core
	core := zapcore.NewTee(core1, core2)
	// 获得准确的调用信息就需要通过AddCallerSkip函数来跳过
	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)) // 添加调用者信息
	SugarLogger = Logger.Sugar()
}

//func InitLogger() {
//	Logger, _ = zap.NewProduction()
//	SugarLogger = Logger.Sugar()
//}

func SimpoleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		Logger.Error(
			"Error fetching url...",
			zap.String("url", url),
			zap.Error(err))
	} else {
		Logger.Info("Success.",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}

func SimpleHttpGet1(url string) {
	SugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		SugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		SugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}

func GinLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		Logger.Info(path,
			zap.Int("statusCode", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

func GinRecovery(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
