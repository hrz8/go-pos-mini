package models

type (
	// Merchant represents movie object from omdb response
	Merchant struct {
		ID          uint64   `gorm:"column:id;primaryKey" json:"id"`
		Name        string   `gorm:"column:name;index:idx_code;unique;not null;type:text" json:"name"`
		Description string   `gorm:"column:description;type:text;not null;default:null" json:"description"`
		Outlets     []Outlet `gorm:"foreignKey:OutletID" json:"outlets,omitempty"`
	}

	// MerchantPayloadGetAll represents payload to get all user
	MerchantPayloadGetAll struct {
		Limit *uint64 `query:"limit"`
		Page  *uint64 `query:"page"`
	}
)
