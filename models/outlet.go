package models

type (
	// Outlet represents movie object from omdb response
	Outlet struct {
		ID          uint64     `gorm:"column:id;primaryKey" json:"id"`
		Name        string     `gorm:"column:name;index:idx_code;unique;not null;type:text" json:"name"`
		Description string     `gorm:"column:description;type:text;not null;default:null" json:"description"`
		Merchant    *Merchant  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"merchant,omitempty"`
		Products    []*Product `gorm:"many2many:outlets_products" json:"products,omitempty"`
	}

	// PartnersPartnerTypes represents join table schema for partner -> partner_type
	OutletsProducts struct {
		OutletID  uint64
		ProductID uint64
		Price     uint64
	}

	// OutletPayloadGetAll represents payload to get all user
	OutletPayloadGetAll struct {
		Limit *uint64 `query:"limit"`
		Page  *uint64 `query:"page"`
	}
)
