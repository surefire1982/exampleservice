package entity

import "errors"

// ErrNotFound not found
var ErrNotFound = errors.New("Not found")

// ErrCannotBeDeleted user cannot be deleted
var ErrCannotBeDeleted = errors.New("Cannot be deleted")

// ErrAlreadyExists user record already exists
var ErrAlreadyExists = errors.New("User already exists")
