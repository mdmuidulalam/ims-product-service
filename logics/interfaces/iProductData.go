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
	* Action: Find Products
	* Return (productsChan) => (channel)
	 */
	FindProducts(productsChan chan []IProductInformation, pageNumber int, pageSize int, orderBy string, isOrderbyIncreasing bool)

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
