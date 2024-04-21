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

type ErrGenresNotFoundAPI struct {
}

func (e ErrGenresNotFoundAPI) Error() string {
	return "No genres found in API"
}

type ErrDevelopersNotFoundAPI struct {
}

func (e ErrDevelopersNotFoundAPI) Error() string {
	return "No developers found in API"
}

type ErrPublishersNotFoundAPI struct {
}

func (e ErrPublishersNotFoundAPI) Error() string {
	return "No publishers found in API"
}

type ErrGamesNotFoundAPI struct {
}

func (e ErrGamesNotFoundAPI) Error(gameName string) error {
	return fmt.Errorf("no games found in API with name [%s]", gameName)
}
