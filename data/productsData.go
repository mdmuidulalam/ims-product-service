package data

import (
	datainterface "product-service/data/interfaces"
	logicinterface "product-service/logics/interfaces"
)

type ProductsData struct {
	Id           int                         `gorm:"column:id"`
	DisplayId    string                      `gorm:"column:displayId"`
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
	productsData.PostgresData.Connect()
	defer productsData.PostgresData.Disconnect()

	product := ProductsData{}
	dbInstance := productsData.PostgresData.GetDatabaseInstance()
	queryContext := productsData.PostgresData.PrepareWhereClause(dbInstance.Table("products"), conditions)
	queryContext.Find(&product)
	if product.Id != 0 {
		var productInformation logicinterface.IProductInformation = &product
		return &productInformation
	} else {
		return nil
	}
}
