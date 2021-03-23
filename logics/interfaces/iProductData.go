package logicinterface

type IProductData interface {
	/*
	* Action: Find a Product
	* Return (conditions) => (Product)
	 */
	FindProduct(conditions []map[string]interface{}) *IProductInformation
}
