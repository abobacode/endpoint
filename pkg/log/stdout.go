package log

import (
	"log"
)

const (
	Reset  Color = "\033[0m"
	Red    Color = "\033[31m"
	Green  Color = "\033[32m"
	Yellow Color = "\033[33m"
	Cyan   Color = "\033[36m"
)

type Color string

func Colored(content string, color Color) string {
	return string(color) + content + string(Reset)
}

// nolint:gochecknoglobals // it's OK
var (
	cyan   = Colored("[PUDGE]", Cyan)
	yellow = Colored("[PUDGE]", Yellow)
	red    = Colored("[PUDGE]", Red)
)

func Info(message interface{}) {
	log.Printf("%s: %v \n", cyan, message)
}

func Warning(message interface{}) {
	log.Printf("%s: %v \n", yellow, message)
}
