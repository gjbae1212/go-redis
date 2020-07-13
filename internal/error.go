package internal

import (
	"errors"
)

var (
	ErrorRedisEmpty    = errors.New("redis empty")
	ErrorRedisNotFound = errors.New("redis not found")
	ErrorRedisUnknown  = errors.New("redis unknown")
	ErrorRedisInvalid  = errors.New("redis invalid")
	ErrorRedisAlready  = errors.New("redis already")
)
