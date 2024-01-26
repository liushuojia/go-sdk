package mysqlConn

import "gorm.io/gorm"

var db *gorm.DB

func DefaultDB(oldGormDB *gorm.DB) *gorm.DB {
	if oldGormDB != nil {
		return oldGormDB.Session(
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
func UseDB(gormDB *gorm.DB) {
	db = gormDB
}
