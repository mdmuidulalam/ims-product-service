package data

import (
	datainterface "product-service/data/interfaces"
	logicinterface "product-service/logics/interfaces"
)

type ProductsData struct {
	Id           int                         `gorm:"column:id"`
	DisplayId    string                      `gorm:"column:displayid"`
	Name         string                      `gorm:"column:name"`
	Description  string                      `gorm:"column:description"`
	PostgresData datainterface.IPostgresData `gorm:"-"`
}

func (productsData *ProductsData) GetId() int {
	return productsData.Id
}

func (productsData *ProductsData) GetDisplayId() string {
	return productsData.DisplayId
}

func (productsData *ProductsData) GetName() string {
	return productsData.Name
}

func (productsData *ProductsData) GetDescription() string {
	return productsData.Description
}

func (productsData *ProductsData) FindProduct(conditions []map[string]interface{}) *logicinterface.IProductInformation {
	queryContext := productsData.PostgresData.PrepareWhereClause(productsData.PostgresData.GetDatabaseInstance().Table("products"), conditions)
	queryContext.Find(&productsData)

	var productInformation logicinterface.IProductInformation = productsData
	return &productInformation
}
