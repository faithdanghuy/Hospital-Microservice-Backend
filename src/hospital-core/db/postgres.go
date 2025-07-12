package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Database struct {
	Executor *gorm.DB
}

func New(conn Connection) *Database {
	db, err := gorm.Open(postgres.Open(conn.ToConnectionString()), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	settingDb, err := db.DB()
	if err != nil {
		panic(err)
	}

	settingDb.SetMaxOpenConns(conn.MaxOpenConnections)
	settingDb.SetMaxIdleConns(conn.MaxIdleConnections)
	settingDb.SetConnMaxIdleTime(conn.ConnectionMaxIdleTime)
	settingDb.SetConnMaxLifetime(conn.ConnectionMaxLifeTime)

	return &Database{
		Executor: db,
	}
}
