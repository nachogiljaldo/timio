package timer

import "time"

type Timer struct {
	ID string
	Expiration int64
	Content string
}

func (timer Timer) IsExpired() bool {
	return timer.Expiration < (time.Now().UnixNano() / int64(time.Millisecond))
}
