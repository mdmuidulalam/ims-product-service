package logicinterface

type IProductInformation interface {
	GetId() int
	GetDisplayId() string
	GetName() string
	GetDescription() string
}
