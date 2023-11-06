package logger

var glogger Logger

// Logger 使用interface解耦. 调用方不关心log的具体实现，可自由接入
type Logger interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Debug(args ...interface{})
	Fatal(args ...interface{})

	Infof(fmt string, args ...interface{})
	Warnf(fmt string, args ...interface{})
	Errorf(fmt string, args ...interface{})
	Debugf(fmt string, args ...interface{})
	Fatalf(fmt string, args ...interface{})
}

// SetLogger 设置日志，初始化的时候调用。(不要并发调用，不安全)
func SetLogger(l Logger) {
	if l == nil {
		panic("nil logger")
	}
	//
	glogger = l
}

func Info(args ...interface{}) {
	glogger.Info(args...)
}
func Warn(args ...interface{}) {
	glogger.Warn(args...)
}
func Error(args ...interface{}) {
	glogger.Error(args...)
}
func Debug(args ...interface{}) {
	glogger.Debug(args...)
}
func Fatal(args ...interface{}) {
	glogger.Fatal(args...)
}

func Infof(fmt string, args ...interface{}) {
	glogger.Infof(fmt, args...)
}
func Warnf(fmt string, args ...interface{}) {
	glogger.Warnf(fmt, args...)
}
func Errorf(fmt string, args ...interface{}) {
	glogger.Errorf(fmt, args...)
}
func Debugf(fmt string, args ...interface{}) {
	glogger.Debugf(fmt, args...)
}
func Fatalf(fmt string, args ...interface{}) {
	glogger.Fatalf(fmt, args...)
}
