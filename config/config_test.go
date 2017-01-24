package config

import (
	"io"
	"testing"
)

func TestConfigReadSuccess(t *testing.T) {
	cases := []struct {
		in string
		want Config
	}{
		{"../timio.yml", Config{HttpPort:8080}},
	}
	for _, c := range cases {
		got, err := LoadConfiguration(c.in)
		if got != c.want {
			t.Errorf("LoadConfiguration(%q) == %q, want %q with err %q", c.in, got, c.want, err)
		}
	}
}

func TestConfigReadErrors(t *testing.T) {
	cases := []struct {
		in string
		want error
	}{
		{"idontexist.yml", io.EOF},
	}
	for _, c := range cases {
		config, got := LoadConfiguration(c.in)
		if got == nil {
			t.Errorf("LoadConfiguration(%q) == %q, want %q with err %q", c.in, config, c.want, got)
		}
	}
}