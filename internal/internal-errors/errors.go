package internalerrors

import (
	"errors"

	"gorm.io/gorm"
)

var ErrInternal error = errors.New("internal server error")

var NotFound error = errors.New("record not found")

func GetError(err error, msg string) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return NotFound
	}

	if msg != "" {
		return errors.New(msg)
	}

	return ErrInternal
}
