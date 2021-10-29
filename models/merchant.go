package models

type (
	// Merchant represents movie object from omdb response
	Merchant struct {
		ID          uint64   `gorm:"column:id;primaryKey" json:"id"`
		Name        string   `gorm:"column:name;index:idx_name;unique;not null" json:"name"`
		Description string   `gorm:"column:description;type:text" json:"description"`
		Outlets     []Outlet `gorm:"foreignKey:MerchantID" json:"outlets,omitempty"`
	}

	// MerchantPayloadGetAll represents payload to get all user
	MerchantPayloadGetAll struct {
		Limit *uint64 `query:"limit"`
		Page  *uint64 `query:"page"`
	}
)
