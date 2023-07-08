package errors

import (
	"fmt"
)

type ErrPlatformNotFound struct {
	Slug string
}

func (e ErrPlatformNotFound) Error() string {
	return fmt.Sprintf("Platform with slug [%s] not found", e.Slug)
}
