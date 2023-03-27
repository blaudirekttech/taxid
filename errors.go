package taxid

import (
	"errors"
)

var (
	ErrInvalidTaxID       = errors.New("invalid tax id")
	ErrUnsupportedCountry = errors.New("country not supported")
)
