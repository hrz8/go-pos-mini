package models

type (
	// Product represents movie object from omdb response
	Product struct {
		ID          uint64    `gorm:"column:id;primaryKey" json:"id"`
		Name        string    `gorm:"column:name;index:idx_name;unique;not null" json:"name"`
		Description string    `gorm:"column:description;type:text" json:"description"`
		Outlets     []*Outlet `gorm:"many2many:outlets_products" json:"outlets,omitempty"`
	}

	// ProductPayloadCreate represents json payload schema to create user
	ProductPayloadCreate struct {
		Name        string `json:"name" validate:"required,min=8"`
		Description string `json:"description"`
	}

	// ProductPayloadUpdate represents json payload schema to update user
	ProductPayloadUpdate struct {
		ID          uint64 `json:"-" param:"id" validate:"required"`
		Name        string `json:"name" validate:"omitempty,min=8"`
		Description string `json:"description"`
	}

	// ProductPayloadDeleteById represents payload to delete user by identifier
	ProductPayloadDeleteById struct {
		ID uint64 `param:"id" validate:"required"`
	}

	// ProductPayloadGetById represents payload to get user by identifier
	ProductPayloadGetById struct {
		ID uint64 `param:"id" validate:"required"`
	}

	// ProductPayloadGetAll represents payload to get all user
	ProductPayloadGetAll struct {
		Limit *uint64 `query:"limit"`
		Page  *uint64 `query:"page"`
	}
)
