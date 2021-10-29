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
		Email     string `json:"email" validate:"required"`
		Password  string `json:"password" validate:"required"`
		FirstName string `json:"firstName" validate:"required"`
		LastName  string `json:"lastName"`
	}

	// UserPayloadLogin represents json payload schema to login
	UserPayloadLogin struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	UserJwt struct {
		ID uint64 `json:"id"`
		jwt.StandardClaims
	}
)
