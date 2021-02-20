package log

import "fmt"

const (
	RESET  = "\033[0m"
	BLACK    = "\033[0;97m"
	WHITE  = "\033[0;30m"
	RED    = "\033[0;31m"
	GREEN  = "\033[0;32m"
	ORANGE = "\033[0;33m"
	BLUE   = "\033[0;34m"
	PURPLE = "\033[0;35m"
	CYAN   = "\033[0;96m"
	GRAY   = "\033[0;37m"
	PINK   = "\033[01;95m"
	YELLOW   = "\033[01;93m"
)

func Off(s string) string { return RESET + s }
func Reset(s string) string { return RESET + s }

func White(s string) string { return WHITE+s+RESET }
func Orange(s string) string { return ORANGE+s+RESET }
func Blue(s string) string { return BLUE+s+RESET }
func Purple(s string) string { return PURPLE+s+RESET }
func Cyan(s string) string { return CYAN+s+RESET }
func Gray(s string) string { return GRAY+s+RESET }
func Green(s string) string { return GREEN+s+RESET }
func Red(s string) string { return RED+s+RESET }

func ColorFull()  {
	for i := 30; i < 101; i++ {
		fmt.Printf("\033[0;%dm 0;%dm \033[01;%dm 01;%dm\n", i,i, i,i)
	}
}