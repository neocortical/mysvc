package mysvc

import "errors"

// ErrNotFound signifies that a single requested object was not found.
var ErrNotFound = errors.New("not found")

// User is a user business object.
type User struct {
	ID   int64
	Name string
}

// Service defines the interface exposed by this package.
type Service interface {
	GetUser(id int64) (User, error)
	GetUsers(ids []int64) (map[int64]User, error)
}
