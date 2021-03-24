package routesinterface

type ICreateProductLogic interface {
	/*
	* Action: Create a Product
	* Return (int, []string) => (Status, ErrorTypes)
	* Return => (0, []) => Product creation successfull
	* Return => (1, [1,2]) => A product with same displayId or name already exists. ErrorType{1 => displayId, 2 => name}
	 */
	CreateProduct() (int, []int)

	// * Attribute value setters
	SetName(string)
	SetDisplayId(string)
	SetDescription(string)
}
