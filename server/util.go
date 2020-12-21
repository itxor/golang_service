package server

import (
	"time"
)

// CallTimeFormat is a call time format
const CallTimeFormat = "15:04"

func validateCallTime(t string) error {
	_, err := time.Parse(CallTimeFormat, t)

	return err
}
