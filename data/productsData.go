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

func (productsData *ProductsData) SetId(id int) {
	productsData.Id = id
}

func (productsData *ProductsData) SetDisplayId(displayId string) {
	productsData.DisplayId = displayId
}

func (productsData *ProductsData) SetName(name string) {
	productsData.Name = name
}

func (productsData *ProductsData) SetDescription(description string) {
	productsData.Description = description
}

func (productsData *ProductsData) ConnectDatabase() {
	productsData.PostgresData.Connect()
}

func (productsData *ProductsData) DisconnectDatabase() {
	productsData.PostgresData.Disconnect()
}

func (productsData *ProductsData) FindProduct(conditions []map[string]interface{}, productChan chan logicinterface.IProductInformation) {
	product := ProductsData{}
	dbInstance := productsData.PostgresData.GetDatabaseInstance()
	queryContext := productsData.PostgresData.PrepareWhereClause(dbInstance.Table("products"), conditions)
	queryContext.Find(&product)
	if product.Id != 0 {
		var productInformation logicinterface.IProductInformation = &product
		productChan <- productInformation
	} else {
		productChan <- nil
	}
}

func (productsData *ProductsData) InsertProduct(productChan chan int) {
	productsData.Id = 0
	queryContext := productsData.PostgresData.GetDatabaseInstance().Table("products")
	queryContext.Create(&productsData)
	productChan <- productsData.Id
}
