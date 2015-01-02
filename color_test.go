package term

import "testing"

func TestSColor(t *testing.T) {
	cases := map[string]string{
		"@blue(Brennan is cool)":      "\033[0;34mBrennan is cool\033[0m",
		"@blue(Hello) @green(world)!": "\033[0;34mHello\033[0m \033[0;32mworld\033[0m!",
	}

	for in, want := range cases {
		got := SColor(in)
		if got != want {
			t.Errorf("SColor(%s) == %s, want %s", in, got, want)
		}
	}
}
