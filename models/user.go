package models

import "github.com/golang-jwt/jwt"

type (
	// User represents movie object from omdb response
	User struct {
		ID        uint64  `gorm:"column:id;primaryKey" json:"id"`
		Email     string  `gorm:"column:email;index:idx_code;unique;not null" json:"email"`
		Password  *string `gorm:"column:password;not null;type:text" json:"password,omitempty"`
		FirstName string  `gorm:"column:firstName;not null" json:"firstName"`
		LastName  *string `gorm:"column:lastName" json:"lastName"`
	}

	// UserPayloadCreate represents json payload schema to create user
	UserPayloadCreate struct {
		Email     string `json:"email" validate:"required,email"`
		Password  string `json:"password" validate:"required,min=8"`
		FirstName string `json:"firstName" validate:"required,min=3"`
		LastName  string `json:"lastName"`
	}

	// UserPayloadLogin represents json payload schema to login
	UserPayloadLogin struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	// UserPayloadUpdate represents json payload schema to update user
	UserPayloadUpdate struct {
		ID        uint64  `json:"-" param:"id" validate:"required"`
		Password  *string `json:"password" validate:"omitempty,min=8"`
		FirstName string  `json:"firstName" validate:"min=3"`
		LastName  *string `json:"lastName"`
	}

	// UserPayloadDeleteById represents payload to delete user by identifier
	UserPayloadDeleteById struct {
		ID uint64 `param:"id" validate:"required"`
	}

	UserJwt struct {
		ID uint64 `json:"id"`
		jwt.StandardClaims
	}
)
