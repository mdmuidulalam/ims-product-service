package routeinterface

type IReadProductLogic interface {
	/*
	* Action: Read Display Card Data of a Product
	* Return (int) => (Status)
	 */
	ReadDisplayCard() int

	/*
	* Action: Read Bulk Data of Products
	* Return ([]*IReadProductLogic) => (Array of product read interface)
	 */
	ReadBulk(pageNumber int, pageSize int, orderBy string, isOrderbyIncreasing bool) []*IReadProductLogic

	// * Attribute value setters
	SetId(id int)
	// * Attribute value getters
	GetId() int
	GetName() string
	GetDisplayId() string
	GetDescription() string
}
