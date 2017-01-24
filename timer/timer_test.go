package timer

import "testing"
import (
	"time"
)

func TestTimer(t *testing.T) {
	cases := []struct {
		in int64
		want bool
	}{
		{nowInMillis() - int64(time.Millisecond), true},
		{nowInMillis() + int64(time.Millisecond), false},
	}
	for _, c := range cases {
		timer := Timer{}
		timer.Expiration = c.in
		got := timer.IsExpired()
		if got != c.want {
			t.Errorf("Timer(%q).IsExpired() == %q, want %q", c.in, got, c.want)
		}
	}
}

func nowInMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}