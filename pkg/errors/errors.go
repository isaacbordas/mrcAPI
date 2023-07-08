package errors

import (
	"fmt"
)

type ErrItemNotFound struct {
	Entity string
	Slug   string
}

func (e ErrItemNotFound) Error() string {
	return fmt.Sprintf("%s with slug [%s] not found", e.Entity, e.Slug)
}
