package cli

import "fmt"

const (
	InfoColor    = "\033[1;34m"
	NoticeColor  = "\033[1;36m"
	WarningColor = "\033[1;33m"
	ErrorColor   = "\033[1;31m"
	DebugColor   = "\033[0;36m"
	ColorTail = "\033[0m"
)

func Info(s string) {
	fmt.Print(InfoColor + s + ColorTail)
}

func Notice(s string) {
	fmt.Print(NoticeColor + s + ColorTail)
}

func Warning(s string) {
	fmt.Print(WarningColor + s + ColorTail)
}

func Error(s string) {
	fmt.Print(ErrorColor + s + ColorTail)
}

func Debug(s string) {
	fmt.Print(DebugColor + s + ColorTail)
}

func Warningf(format string, args ...interface{}) {
	Warning(fmt.Sprintf(format, args))
}

func Errorf(format string, args ...interface{}) {
	Error(fmt.Sprintf(format, args))
}