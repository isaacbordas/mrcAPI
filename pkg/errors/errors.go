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

type ErrEntityNotAllowed struct {
	Entity string
}

func (e ErrEntityNotAllowed) Error() string {
	return fmt.Sprintf("Entity [%s] not allowed", e.Entity)
}

type ErrPlatformsNotFoundAPI struct {
}

func (e ErrPlatformsNotFoundAPI) Error() string {
	return "No platforms found in API"
}
