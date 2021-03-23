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

func (createProductLogic *CreateProductLogic) CreateProduct() int {
	conditions := []map[string]interface{}{
		map[string]interface{}{
			"displayId": createProductLogic.displayId,
		},
		map[string]interface{}{
			"name": createProductLogic.name,
		},
	}

	product := createProductLogic.ProductData.FindProduct(conditions)
	if product != nil {
		return 0
	} else {
		return 1
	}
}
