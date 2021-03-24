package logics

import (
	logicinterface "product-service/logics/interfaces"
)

type CreateProductLogic struct {
	displayId   string
	name        string
	description string
	ProductData logicinterface.IProductData
}

func (createProductLogic *CreateProductLogic) SetDisplayId(displayId string) {
	createProductLogic.displayId = displayId
}

func (createProductLogic *CreateProductLogic) SetName(name string) {
	createProductLogic.name = name
}

func (createProductLogic *CreateProductLogic) SetDescription(description string) {
	createProductLogic.description = description
}

func (createProductLogic *CreateProductLogic) CreateProduct() (int, []int) {
	createProductLogic.ProductData.ConnectDatabase()
	defer createProductLogic.ProductData.DisconnectDatabase()

	duplicateDisplayIdProductChan := make(chan logicinterface.IProductInformation)
	conditions := []map[string]interface{}{{
		"displayId": createProductLogic.displayId,
	}}
	go createProductLogic.ProductData.FindProduct(conditions, duplicateDisplayIdProductChan)

	duplicateNameProductChan := make(chan logicinterface.IProductInformation)
	conditions = []map[string]interface{}{{
		"name": createProductLogic.name,
	}}
	go createProductLogic.ProductData.FindProduct(conditions, duplicateNameProductChan)

	duplicateDisplayIdProduct, duplicateNameProduct := <-duplicateDisplayIdProductChan, <-duplicateNameProductChan

	errorTypes := []int{}
	statusCode := 0
	if duplicateDisplayIdProduct != nil {
		statusCode = 1
		errorTypes = append(errorTypes, 1)
	}
	if duplicateNameProduct != nil {
		statusCode = 1
		errorTypes = append(errorTypes, 2)
	}

	if statusCode == 0 {
		createProductLogic.ProductData.SetId(0)
		createProductLogic.ProductData.SetDisplayId(createProductLogic.displayId)
		createProductLogic.ProductData.SetName(createProductLogic.name)
		createProductLogic.ProductData.SetDescription(createProductLogic.description)

		productIdChan := make(chan int)
		go createProductLogic.ProductData.InsertProduct(productIdChan)
		<-productIdChan
	}

	return statusCode, errorTypes
}
