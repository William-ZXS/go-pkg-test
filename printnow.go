package go_pkg_test

import (
	"fmt"
	"time"
)

func PrintNow() string {
	t := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("v2.0.0 %s", t)
}
