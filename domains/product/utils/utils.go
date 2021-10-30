package utils

import (
	"github.com/hrz8/go-pos-mini/models"
)

// GetPrice get price from given outletID and productID
func GetPrice(objs *[]models.OutletsProducts, outletID *uint64, productID *uint64) uint64 {
	for _, v := range *objs {
		if v.OutletID == *outletID && v.ProductID == *productID {
			return v.Price
		}
	}
	return 0
}
