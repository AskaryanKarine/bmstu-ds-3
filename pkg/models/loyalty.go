package models

type LoyaltyStatusType string

const (
	BRONZE LoyaltyStatusType = "BRONZE"
	SILVER LoyaltyStatusType = "SILVER"
	GOLD   LoyaltyStatusType = "GOLD"
)

type LoyaltyInfoResponse struct {
	// Status - статус в программе лояльности
	Status LoyaltyStatusType `json:"status,omitempty"`
	// Discount - скидка по программе лояльности
	Discount int `json:"discount,omitempty"`
	// ReservationCount - количество бронирований
	ReservationCount *int `json:"reservationCount,omitempty"`
}
