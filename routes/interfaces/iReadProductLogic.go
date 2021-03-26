package routesinterface

type IReadProductLogic interface {
	/*
	* Action: Read Display Card Data of a Product
	* Return (int) => (Status)
	 */
	ReadDisplayCard() int

	// * Attribute value setters
	SetId(id int)
	// * Attribute value getters
	GetId() int
	GetName() string
	GetDisplayId() string
	GetDescription() string
}
