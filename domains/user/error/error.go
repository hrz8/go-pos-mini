package domain_error

import "errors"

type (
	errorMap struct {
		Status int
		Err    error
	}
)

var (
	Create = errorMap{
		Status: 400,
		Err:    errors.New("failed to create user"),
	}
	WrongPassword = errorMap{
		Status: 403,
		Err:    errors.New("username or password is wrong"),
	}
	Login = errorMap{
		Status: 403,
		Err:    errors.New("failed to login the user"),
	}
	GetBy = errorMap{
		Status: 404,
		Err:    errors.New("user is not found"),
	}
	UpdateById = errorMap{
		Status: 400,
		Err:    errors.New("failed to update user"),
	}
	DeleteById = errorMap{
		Status: 403,
		Err:    errors.New("failed to delete user"),
	}
)
