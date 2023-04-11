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
	return option.None[StringLeaf]()
}

func (i Invalid) Number() option.Option[NumberLeaf] {
	return option.None[NumberLeaf]()
}

func (i Invalid) Boolean() option.Option[BooleanLeaf] {
	return option.None[BooleanLeaf]()
}

func (i Invalid) Object() option.Option[ObjectBranch] {
	return option.None[ObjectBranch]()
}

func (i Invalid) Array() option.Option[ArrayBranch] {
	return option.None[ArrayBranch]()
}
