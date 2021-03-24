package datainterface

import "gorm.io/gorm"

type IPostgresData interface {
	GetDatabaseInstance() *gorm.DB
	Connect()
	Disconnect()
	PrepareWhereClause(context *gorm.DB, conditions []map[string]interface{}) *gorm.DB
}
