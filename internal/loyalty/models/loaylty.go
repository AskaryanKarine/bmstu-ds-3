package models

import "github.com/AskaryanKarine/bmstu-ds-3/pkg/models"

type ExpandedLoyalty struct {
	models.LoyaltyInfoResponse
	Discount int
}
