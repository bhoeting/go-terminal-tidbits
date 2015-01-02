package term

import "testing"

func TestSColor(t *testing.T) {
	cases := map[string]string{
		"@blue(Brennan is cool)":                     "\033[0;34mBrennan is cool\033[0m",
		"@blue(Hello) @green(world)!":                "\033[0;34mHello\033[0m \033[0;32mworld\033[0m!",
		"@green(brennan.hoeting)(@)gmail.com":        "\033[0;32mbrennan.hoeting\033[0m@gmail.com",
		"@green(brennan.hoeting)(@)@blue(gmail.com)": "\033[0;32mbrennan.hoeting\033[0m@\033[0;34mgmail.com\033[0m",
		"(@) Another escaping test (@)":              "@ Another escaping test @",
	}

	for in, want := range cases {
		got := SColor(in)
		if got != want {
			t.Errorf("SColor(%s) == %s, want %s", in, got, want)
		}
	}
}
