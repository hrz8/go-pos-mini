package models

type (
	// Product represents movie object from omdb response
	Product struct {
		ID          uint64    `gorm:"column:id;primaryKey" json:"id"`
		Name        string    `gorm:"column:name;index:idx_code;unique;not null;type:text" json:"name"`
		Description string    `gorm:"column:description;type:text;not null;default:null" json:"description"`
		Outlets     []*Outlet `gorm:"many2many:outlets_products" json:"outlets,omitempty"`
	}
)
