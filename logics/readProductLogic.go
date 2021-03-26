package logics

import (
	logicinterface "product-service/logics/interfaces"
)

type ReadProductLogic struct {
	id          int
	displayId   string
	name        string
	description string
	ProductData logicinterface.IProductData
}

func (readProductLogic *ReadProductLogic) SetId(id int) {
	readProductLogic.id = id
}

func (readProductLogic *ReadProductLogic) GetId() int {
	return readProductLogic.id
}

func (readProductLogic *ReadProductLogic) GetDisplayId() string {
	return readProductLogic.displayId
}

func (readProductLogic *ReadProductLogic) GetName() string {
	return readProductLogic.name
}

func (readProductLogic *ReadProductLogic) GetDescription() string {
	return readProductLogic.description
}

func (readProductLogic *ReadProductLogic) ReadDisplayCard() int {
	readProductLogic.ProductData.ConnectDatabase()
	defer readProductLogic.ProductData.DisconnectDatabase()

	productChan := make(chan logicinterface.IProductInformation)
	conditions := []map[string]interface{}{{
		"id": readProductLogic.id,
	}}
	go readProductLogic.ProductData.FindProduct(conditions, productChan)

	product := <-productChan

	statusCode := 0
	if product != nil {
		readProductLogic.displayId = product.GetDisplayId()
		readProductLogic.name = product.GetName()
	} else {
		statusCode = 1
		readProductLogic.id = 0
	}

	return statusCode
}
