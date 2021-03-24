package logicinterface

type IProductData interface {
	/*
	* Action: Connect with database
	 */
	ConnectDatabase()

	/*
	* Action: Disconnect with database
	 */
	DisconnectDatabase()

	/*
	* Action: Find a Product
	* Return (conditions, productChan) => (Product, channel)
	 */
	FindProduct(conditions []map[string]interface{}, productChan chan IProductInformation)

	/*
	* Action: Insert a Product
	* Return (productChan) => (channel)
	 */
	InsertProduct(productChan chan int)

	// * Attribute value setters
	SetId(int)
	SetName(string)
	SetDisplayId(string)
	SetDescription(string)
}
