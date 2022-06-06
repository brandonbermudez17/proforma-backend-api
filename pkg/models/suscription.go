package models

import (
	"time"

	"gorm.io/gorm"
)

type Subscription struct {
	gorm.Model
	StartDate          time.Time
	EndDate            time.Time
	PaymentMethod      string
	Amount             float32
	Currency           string
	CompanyID          uint
	Company            Company
	SubscriptionTypeID uint
	SubscriptionType   SubscriptionType
}

type SubscriptionType struct {
	Name                   string
	Description            string
	NumberUsersAllowed     uint
	NumberSuppliersAllowed uint
	Active                 bool
}
