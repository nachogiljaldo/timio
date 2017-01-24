package main

import "testing"

func TestArguments(t *testing.T) {
	cases := []struct {
		in []string
		want string
	}{
		{[]string{"--something=xxx"}, DefaultFile},
		{[]string{"--file=xxx"}, "xxx"},
		{[]string{"--filexxx"}, DefaultFile},
		{[]string{"--files=xxx"}, DefaultFile},
		{[]string{"asdfasdf", "--files=xxx"}, DefaultFile},
	}
	for _, c := range cases {
		got := GetConfigFile(c.in)
		if got != c.want {
			t.Errorf("GetConfigFile(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}