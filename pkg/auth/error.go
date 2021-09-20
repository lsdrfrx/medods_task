package auth

import "errors"

var ErrInvalidToken = errors.New("invalid token")
var ErrUnexpectedSingingMethod = errors.New("unexpected singing method")
var ErrEmptyToken = errors.New("token is empty")
var ErrUnableToDecodeToken = errors.New("unable to decode token")
var ErrNoTokenInDatabase = errors.New("unable to find token in database")
var ErrInvalidTokenPair = errors.New("invalid token pair")
