package log
import (
	"fmt"
	"io"
	"os"
	"runtime"
	"time"
)

const (
	PREFIX      = "[FOUR-SEASONS]"
	TIME_FORMAT = "2006-01-02 15:04:05"
)

var (
	Verbose, NonColor bool
	Output            io.Writer = os.Stdout

	LEVEL_FLAGS = [...]string{"DEBUG", " INFO", " WARN", "ERROR", "FATAL"}
)

func init() {
	if runtime.GOOS == "windows" {
		NonColor = true
	}
}

const (
	DEBUG = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func Print(level int, format string, args ...interface{}) {

	var logFormat = "%s %s [%s] %s\n"
	if !NonColor {
		switch level {
		case DEBUG:
			logFormat = "%s \033[36m%s\033[0m [\033[34m%s\033[0m] %s\n"
		case INFO:
			logFormat = "%s \033[36m%s\033[0m [\033[32m%s\033[0m] %s\n"
		case WARNING:
			logFormat = "%s \033[36m%s\033[0m [\033[33m%s\033[0m] %s\n"
		case ERROR:
			logFormat = "%s \033[36m%s\033[0m [\033[31m%s\033[0m] %s\n"
		case FATAL:
			logFormat = "%s \033[36m%s\033[0m [\033[35m%s\033[0m] %s\n"
		}
	}

	fmt.Fprintf(Output, logFormat, PREFIX, time.Now().Format(TIME_FORMAT),
		LEVEL_FLAGS[level], fmt.Sprintf(format, args...))

	if level == FATAL {
		os.Exit(1)
	}
}

func Debug(format string, args ...interface{}) {
	Print(DEBUG, format, args...)
}

func Warn(format string, args ...interface{}) {
	Print(WARNING, format, args...)
}

func Info(format string, args ...interface{}) {
	Print(INFO, format, args...)
}

func Error(format string, args ...interface{}) {
	Print(ERROR, format, args...)
}

func Fatal(format string, args ...interface{}) {
	Print(FATAL, format, args...)
}