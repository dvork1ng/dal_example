package core

import "errors"

var (
	ErrEcosystemNotFound   = errors.New("ecosystem not found")
	ErrRecordNotFound      = errors.New("record not found")
	ErrRecordAlreadyExists = errors.New("record already exists")
)
