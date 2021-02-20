package log

import "fmt"

type Level struct {
	name string
	level int
	prefix string
}

var (
	level int = 10

	DEBUG = Level{name: "DEBUG", level: 100, prefix: CYAN+"DEBUG"}
	INFO  = Level{name: "INFO", level: 10, prefix: YELLOW+"INFO"}
	WARN  = Level{name: "WARN", level: 2, prefix: ORANGE+"WARN"}
	FATAL = Level{name: "FATAL", level: 1, prefix: RED+"FATAL"}
	ERROR = Level{name: "ERROR", level: 1, prefix: RED+"ERROR"}
)

type Formatter struct {
	level Level
	text string
}

func (f Formatter) build() string {
	return fmt.Sprintf("%s: %s%s", f.level.prefix, WHITE, f.text)
}

func Log(level Level, text string) {
	fmt.Print(Formatter{level: level, text: fmt.Sprintf(text)}.build())
}


func Format(level Level, text string, o ...interface{}) {
	if len(o) > 0 {
		fmt.Println(Formatter{level: level, text: fmt.Sprintf(text, o...)}.build())
	} else {
		fmt.Println(Formatter{level: level, text: fmt.Sprintf(text)}.build())
	}
}

func Info(s string, o ...interface{}) { Format(INFO, s, o...) }
func Warn(s string, o ...interface{}) { Format(WARN, s, o...) }
func Error(s string, err error) { Format(ERROR, s, err.Error()) }
func Fatal(s string, o ...interface{}) { Format(FATAL, s, o...) }
func Debug(s string, o ...interface{}) { Format(DEBUG, s, o...) }

func SetLevel(i int)  { level = i }

