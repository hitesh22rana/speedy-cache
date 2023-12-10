package cache

import "errors"

var (
	ErrKeyNotFound = errors.New("error: key not found")
)

type Cache interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Delete(key string) error
}
