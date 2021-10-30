package models

type (
	// OutletsProductsPayloadGetAll represents payload to get all user
	OutletsProductsPayloadGetAll struct {
		Limit *uint64 `query:"limit"`
		Page  *uint64 `query:"page"`
	}
)
