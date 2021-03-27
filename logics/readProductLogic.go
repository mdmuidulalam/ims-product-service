package logics

import (
	logicinterface "product-service/logics/interfaces"
	routeinterface "product-service/routes/interfaces"
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

func (readProductLogic *ReadProductLogic) ReadBulk(pageNumber int, pageSize int, orderBy string, isOrderbyIncreasing bool) []*routeinterface.IReadProductLogic {
	readProductLogic.ProductData.ConnectDatabase()
	defer readProductLogic.ProductData.DisconnectDatabase()

	productsChan := make(chan []logicinterface.IProductInformation)
	go readProductLogic.ProductData.FindProducts(productsChan, pageNumber, pageSize, orderBy, isOrderbyIncreasing)

	productsData := <-productsChan

	var products []*routeinterface.IReadProductLogic

	for _, productData := range productsData {
		var product ReadProductLogic
		product.id = productData.GetId()
		product.displayId = productData.GetDisplayId()
		product.name = productData.GetName()
		product.description = productData.GetDescription()

		var productInterface routeinterface.IReadProductLogic = &product
		products = append(products, &productInterface)
	}

	return products
}
