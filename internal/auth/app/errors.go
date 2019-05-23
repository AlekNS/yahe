package app

import "errors"

var (
	//
	// Authentication
	//

	// ErrorUserAlreadyExists .
	ErrorUserAlreadyExists = errors.New("user already exists")

	// ErrorInvalidJwtToken .
	ErrorInvalidJwtToken = errors.New("invalid jwt token")

	// ErrorNotFound .
	ErrorNotFound = errors.New("not found")

	// ErrorUserIsNotActive .
	ErrorUserIsNotActive = errors.New("user is not active")

	// ErrorUserIsActive .
	ErrorUserIsActive = errors.New("user is active")

	// ErrorInternalStorageInconsistent .
	ErrorInternalStorageInconsistent = errors.New("internal storage inconsistent")

	// ErrorPasswordIsVeryBasic .
	ErrorPasswordIsVeryBasic = errors.New("password is very basic")

	// ErrorPasswordMismatch .
	ErrorPasswordMismatch = errors.New("password mismatch")

	//
	// Authorezation
	//
)
