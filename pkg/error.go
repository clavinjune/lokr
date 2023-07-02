package pkg

import (
	"errors"
	"fmt"
)

var (
	ErrBase         = errors.New("error lokr-proto")
	ErrNilDB        = fmt.Errorf("%w: database is nil", ErrBase)
	ErrZeroAffected = fmt.Errorf("%w: zero rows affected", ErrBase)
	ErrEmptyLockKey = fmt.Errorf("%w: empty lock key", ErrBase)
)
