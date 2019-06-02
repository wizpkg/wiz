package log

import (
	"fmt"
)

// Log logs the message to the configured backend
func Log(msgs ...interface{}) {
	fmt.Println(msgs...)
}
