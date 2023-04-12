package cjson

import (
	"errors"

	option "github.com/kerelape/option/pkg"
)

var ErrInvalidJSON = errors.New("invalid JSON value")

// Invalid is an invalid JSON value.
type Invalid struct{}

// NewInvalid creates a new Invalid.
func NewInvalid() Invalid {
	return Invalid{}
}

func (i Invalid) MarshalJSON() ([]byte, error) {
	return nil, ErrInvalidJSON
}

func (i Invalid) String() option.Option[StringLeaf] {
	return option.NewNoneReason[StringLeaf](ErrInvalidJSON)
}

func (i Invalid) Number() option.Option[NumberLeaf] {
	return option.NewNoneReason[NumberLeaf](ErrInvalidJSON)
}

func (i Invalid) Boolean() option.Option[BooleanLeaf] {
	return option.NewNoneReason[BooleanLeaf](ErrInvalidJSON)
}

func (i Invalid) Object() option.Option[ObjectBranch] {
	return option.NewNoneReason[ObjectBranch](ErrInvalidJSON)
}

func (i Invalid) Array() option.Option[ArrayBranch] {
	return option.NewNoneReason[ArrayBranch](ErrInvalidJSON)
}
