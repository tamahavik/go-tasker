package database

import (
	"fmt"
	"project-tasker/models"
	"testing"
)

func TestExecQuery(t *testing.T) {
	db := GetConnection()
	data := models.Customer{
		Name:  "test",
		Phone: "09928392839",
	}

	result := db.Create(&data)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println(result.RowsAffected)
}
