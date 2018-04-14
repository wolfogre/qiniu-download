package log

import "time"

var (
	statusErr error
	errTime time.Time
)

func Status() (bool, string) {
	if time.Now().Add(10 * time.Minute).Before(errTime) {
		if statusErr == nil {
			return true, "ok"
		}
		return false, statusErr.Error()
	}
	return true, "ok"
}

func ChangeStatus(err error) {
	statusErr = err
	errTime = time.Now()
}
