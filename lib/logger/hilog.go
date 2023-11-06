package logger

import (
	"flag"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"hibar/cmd/comet/conf"
	"io"
	"log/syslog"
	"os"
	"strings"
)

var (
	loggerLevel      = flag.String("loggerLevel", "INFO", "最小日志级别. 可能的值: DEBUG, INFO, WARN, ERROR, FATAL")
	logCollectorAddr = flag.String("logCollectorAddr", ":514", "异步日志收集器网络地址位置") // 例 rsyslog.xxxx-fat.svc.cluster.local:514
)

func NewHiLogger() *hiLogger {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, logLevel(*loggerLevel))

	logger := zap.New(core)
	sugarLogger := logger.Sugar()

	hiLog := &hiLogger{SugaredLogger: sugarLogger}

	return hiLog
}

type hiLogger struct {
	*zap.SugaredLogger
}

var _ Logger = hiLogger{}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	if conf.SerENV == conf.Prod {
		encoderConfig = zap.NewProductionEncoderConfig()
	}
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	var ws io.Writer
	if conf.SerENV == conf.Prod {
		sysLog, err := syslog.Dial("udp", *logCollectorAddr, syslog.LOG_WARNING|syslog.LOG_DAEMON, "")
		if err != nil { //syslog.Priority(loglevel)
			// 这个根据自己的实际情况来，添加输出端。比如现在用的多的ELK日志平台，向日志采集器发送日志。(推荐使用UDP输送日志，吞吐高。服务端内部网络环境可控，能避免丢包)
			ws = io.MultiWriter(sysLog)
		}
	} else {
		ws = io.MultiWriter(os.Stdout)
	}
	return zapcore.AddSync(ws)
}

func logLevel(lv string) zapcore.Level {
	lv = strings.ToUpper(lv)
	switch lv {
	case "INFO":
		return zapcore.InfoLevel
	case "WARN":
		return zapcore.WarnLevel
	case "ERROR":
		return zapcore.ErrorLevel
	case "FATAL":
		return zapcore.FatalLevel
	case "DEBUG":
		return zapcore.PanicLevel
	default: // 默认debug
		return zapcore.InfoLevel
	}
}
