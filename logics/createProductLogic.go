package logics

type CreateProductLogic struct {
	displayId   string
	name        string
	description string
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

	return 0
}
