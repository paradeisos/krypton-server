package models

import (
	"errors"

	"gopkg.in/mgo.v2"
)

var (
	ErrNotFound      = mgo.ErrNotFound
	ErrInvalidId     = errors.New("Invalid BSON ID")
	ErrInvalidParams = errors.New("Invalid Params")
	ErrDB            = errors.New("Unknow DB Error")
)
