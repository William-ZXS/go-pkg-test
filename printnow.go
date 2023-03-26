package go_pkg_test

import (
	"time"
)

func PrintNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
