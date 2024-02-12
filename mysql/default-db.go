package mysqlConn

import (
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var db *gorm.DB

func UseDB(gormDB *gorm.DB) {
	db = gormDB
}

func DefaultDB(dbList ...*gorm.DB) *gorm.DB {

	if len(dbList) > 0 && dbList[0] != nil {
		return dbList[0].Session(
			&gorm.Session{},
		)
	}

	return db.Session(
		&gorm.Session{
			QueryFields: true,
			PrepareStmt: true,
			NewDB:       true,
		},
	)
}

func WriteDB(dbList ...*gorm.DB) *gorm.DB {
	if len(dbList) > 0 && dbList[0] != nil {
		return dbList[0].Clauses(dbresolver.Write).Session(
			&gorm.Session{},
		)
	}

	return db.Clauses(dbresolver.Write).Session(
		&gorm.Session{
			QueryFields: true,
			PrepareStmt: true,
			NewDB:       true,
		},
	)
}
