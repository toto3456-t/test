package entity

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	PaymentType string

}