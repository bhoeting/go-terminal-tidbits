package term

import (
	"errors"
	"fmt"
	"strings"
)

// colors is a map of color names and their
// corresponding *nix terminal codes
var colors = map[string]string{
	"red":    "0;31",
	"blue":   "0;34",
	"gray":   "0:37",
	"grey":   "0:37",
	"cyan":   "0;36",
	"black":  "0;30",
	"green":  "0;32",
	"purple": "2;35",
	"yellow": "0;33",
}

// SColorf returns a colorized Sprintf
func SColorf(format string, args ...interface{}) string {
	return SColor(fmt.Sprintf(format, args...))
}

// Colorf outputs a colorized Sprintf
func Colorf(format string, args ...interface{}) {
	fmt.Printf("%s", SColorf(format, args...))
}

// Color outputs a colorized string
func Color(str string) {
	fmt.Print(SColor(str))
}

// Colorln outputs a colorized string with a newline
func Colorln(str string) {
	fmt.Println(SColor(str))
}

// SColor returns a colorized string
func SColor(str string) string {
	index := len(str) - 1

	for index >= 0 {
		if str[index] == '@' {
			startIndex := index
			endIndex := strings.Index(str[startIndex:], ")") + 1 + startIndex
			part := str[startIndex:endIndex]
			colorized, err := colorize(part)

			if err != nil {
				fmt.Errorf("%s", err.Error())
				break
			}

			str = str[:startIndex] + colorized + str[endIndex:]
			index = startIndex - 1
		}
		index--
	}
	return str
}

// colorize returns a string formatted to
// support *nix terminal colors
func colorize(str string) (string, error) {
	openParenIndex := strings.Index(str, "(")
	closeParenIndex := strings.Index(str, ")")

	if openParenIndex == -1 {
		return str, errors.New("Could not find open paren (")
	}

	if closeParenIndex == -1 {
		return str, errors.New("Could not find close paren )")
	}

	color := str[1:openParenIndex]
	text := str[openParenIndex+1 : closeParenIndex]

	return fmt.Sprint("\033[" + colors[color] + "m" + text + "\033[0m"), nil
}
