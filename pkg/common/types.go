package common

import "fmt"

type Error struct {
	ErrorCode        string
	ErrorDescription string
}

func (err Error) Error() string {
	return fmt.Sprintf("%s: %s", err.ErrorCode, err.ErrorDescription)
}
