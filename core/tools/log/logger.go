package log

var logger Logger = NewLog()

type internal interface {
	debug(msg string, fields ...Field)
	info(msg string, fields ...Field)
	warn(msg string, fields ...Field)
	error(msg string, fields ...Field)
	dPanic(msg string, fields ...Field)
	panic(msg string, fields ...Field)
	fatal(msg string, fields ...Field)
	debugf(format string, args ...any)
	infof(format string, args ...any)
	warnf(format string, args ...any)
	errorf(format string, args ...any)
	dPanicf(format string, args ...any)
	panicf(format string, args ...any)
	fatalf(format string, args ...any)
}

type Logger interface {
	internal
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	DPanic(msg string, fields ...Field)
	Panic(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
	DPanicf(format string, args ...any)
	Panicf(format string, args ...any)
	Fatalf(format string, args ...any)
}

func Debug(msg string, fields ...Field) {
	logger.debug(msg, fields...)
}

func Info(msg string, fields ...Field) {
	logger.info(msg, fields...)
}

func Warn(msg string, fields ...Field) {
	logger.warn(msg, fields...)
}

func Error(msg string, fields ...Field) {
	logger.error(msg, fields...)
}

func DPanic(msg string, fields ...Field) {
	logger.dPanic(msg, fields...)
}

func Panic(msg string, fields ...Field) {
	logger.panic(msg, fields...)
}

func Fatal(msg string, fields ...Field) {
	logger.fatal(msg, fields...)
}

func Debugf(format string, args ...any) {
	logger.debugf(format, args...)
}

func Infof(format string, args ...any) {
	logger.infof(format, args...)
}

func Warnf(format string, args ...any) {
	logger.warnf(format, args...)
}

func Errorf(format string, args ...any) {
	logger.errorf(format, args...)
}

func DPanicf(format string, args ...any) {
	logger.dPanicf(format, args...)
}

func Panicf(format string, args ...any) {
	logger.panicf(format, args...)
}

func Fatalf(format string, args ...any) {
	logger.fatalf(format, args...)
}

func SetLogger(log Logger) {
	logger = log
}

func GetLogger() Logger {
	return logger
}
