package data

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresData struct {
	databaseInstance *gorm.DB
}

func (postgresData *PostgresData) GetDatabaseInstance() *gorm.DB {
	return postgresData.databaseInstance
}

func (postgresData *PostgresData) Connect() {
	dsn := "host=localhost user=postgres password=DefaultPassword dbname=ps-defaultClient port=5432 sslmode=disable TimeZone=Asia/Dhaka"
	var err error
	if postgresData.databaseInstance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		panic(err)
	}
}

func (postgresData *PostgresData) Disconnect() {
	sqlDB, err := postgresData.databaseInstance.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
}

func (postgresData *PostgresData) PrepareWhereClause(context *gorm.DB, conditions []map[string]interface{}) *gorm.DB {
	if len(conditions) != 0 {
		context = context.Where(conditions[0])
		for i := 1; i < len(conditions); i++ {
			context = context.Or(conditions[i])
		}
	}

	return context
}
