package database

import (
	"encoding/json"
	"fmt"
	"project-tasker/models"
	"testing"
)

func TestGormInsert(t *testing.T) {
	db := GetConnection()
	data := models.Customer{
		Name:  "test",
		Phone: "09928392839",
	}

	result := db.Create(&data)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println(data)
}

func TestGormBatchInsert(t *testing.T) {
	db := GetConnection()
	data1 := models.Customer{
		Name:  "test1",
		Phone: "09928392831",
	}

	data2 := models.Customer{
		Name:  "test2",
		Phone: "09928392832",
	}

	data3 := models.Customer{
		Name:  "test3",
		Phone: "09928392833",
	}

	var result []models.Customer
	result = append(result, data1, data2, data3)

	db.Create(&result)
	fmt.Println(result)
}

func TestGormFind(t *testing.T) {
	db := GetConnection()
	var data models.Customer
	db.First(&data, 3)
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

}
