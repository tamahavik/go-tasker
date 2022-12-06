package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	// Id    int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func (c *Customer) tableName() string {
	return "CUSTOMER"
}
