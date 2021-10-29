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
		Err:    errors.New("failed to create product"),
	}
	GetBy = errorMap{
		Status: 404,
		Err:    errors.New("product not found"),
	}
	UpdateById = errorMap{
		Status: 400,
		Err:    errors.New("failed to update product"),
	}
	DeleteById = errorMap{
		Status: 400,
		Err:    errors.New("failed to delete product"),
	}
	GetById = errorMap{
		Status: 400,
		Err:    errors.New("failed to get product by id"),
	}
	GetAll = errorMap{
		Status: 400,
		Err:    errors.New("failed to get all product"),
	}
)
