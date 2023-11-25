package wework

import "errors"

var ErrNilStoreToken = errors.New("store token is nil or expired")

var ErrInvalidInternalCorp = errors.New("invalid internal corp config")

var ErrInvalidOpenCorp = errors.New("invalid open corp config")
