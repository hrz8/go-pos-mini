package models

type (
	// Outlet represents movie object from omdb response
	Outlet struct {
		ID          uint64     `gorm:"column:id;primaryKey" json:"id"`
		Name        string     `gorm:"column:name;index:idx_name;unique;not null" json:"name"`
		Description string     `gorm:"column:description;type:text" json:"description"`
		MerchantID  uint64     `gorm:"not null" json:"-"`
		Merchant    *Merchant  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"merchant,omitempty"`
		Products    []*Product `gorm:"many2many:outlets_products" json:"products,omitempty"`
		OutletsProducts
	}

	// PartnersPartnerTypes represents join table schema for partner -> partner_type
	OutletsProducts struct {
		OutletID  uint64 `json:"-"`
		ProductID uint64 `json:"-"`
		Price     uint64 `gorm:"column:price;not null" json:"price"`
	}

	// OutletPayloadGetAll represents payload to get all user
	OutletPayloadGetAll struct {
		Limit *uint64 `query:"limit"`
		Page  *uint64 `query:"page"`
	}
)
