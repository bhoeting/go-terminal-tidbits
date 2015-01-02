package main

import "./terminal_tidbits"

func main() {
	term.Colorf("@green(%s), @yellow(%s)%s\n", "Hello", "world", "!")
	term.Colorln("@blue(Brennan)")

	// w := new(io.Writer)
	// fmt.Fprint(, ...)
}
