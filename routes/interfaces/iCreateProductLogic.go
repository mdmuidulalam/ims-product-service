package routesinterface

type ICreateProductLogic interface {
	/*
	* Action: Create a Product
	* Return (int) => (Status)
	* Return => (0) => Product creation successfull
	* Return => (1) => A product with same displayId already exists
	* Return => (2) => A product with sane name already exists
	 */
	CreateProduct() int

	// * Attribute value setters
	SetName(string)
	SetDisplayId(string)
	SetDescription(string)
}
