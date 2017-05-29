package models

import (
	"errors"
	"gopkg.in/mgo.v2"
)

var (
	ErrNotFound  = mgo.ErrNotFound
	ErrInvalidId = errors.New("Invalid BSON ID")
)
