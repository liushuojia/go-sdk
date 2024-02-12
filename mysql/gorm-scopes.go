package mysqlConn

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type TableIndexField struct {
	Table       string `gorm:"column:Table"`
	KeyName     string `gorm:"column:Key_name"`
	ColumnName  string `gorm:"column:Column_name"`
	NonUnique   int64  `gorm:"column:Non_unique"`
	SeqInIndex  int64  `gorm:"column:Seq_in_index"`
	Collation   string `gorm:"column:Collation"`
	Cardinality int64  `gorm:"column:Cardinality"`
	IndexType   string `gorm:"column:Index_type"`
}

func TableSeparator() string {
	return "_"
}
func TableOfCode(u m) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		tableName := u.TableName()

		field := DefaultField()
		if !field.GetAutoMigrate(tableName) {
			field.SetAutoMigrate(tableName, true)
			rows, err := db.Session(&gorm.Session{NewDB: true}).
				Raw(fmt.Sprintf("desc %s", tableName)).
				Rows()

			if err != nil {
				db.Session(&gorm.Session{NewDB: true}).
					Table(tableName).
					AutoMigrate(u)
			} else {
				defer rows.Close()

				gormFieldMap := make(map[string]GormField)

				cols, _ := rows.Columns()
				for rows.Next() {
					columns := make([]interface{}, len(cols))
					columnPointers := make([]interface{}, len(cols))
					for i := range columns {
						columnPointers[i] = &columns[i]
					}
					if err := rows.Scan(columnPointers...); err != nil {
						continue
					}

					gormField := GormField{}
					for i, colName := range cols {
						val := columnPointers[i].(*interface{})

						value := ""
						if fmt.Sprintf("%T", *val) == "[]uint8" {
							value = string((*val).([]uint8))
						}
						if fmt.Sprintf("%T", *val) == "int64" {
							value = strconv.Itoa(int((*val).(int64)))
						}
						if value == "" {
							continue
						}

						switch colName {
						case "Extra":
							if value == "auto_increment" {
								gormField.AutoIncrement = true
							}
						case "Field":
							gormField.Column = value
						case "Key":
							switch value {
							case "PRI":
								gormField.AutoIncrement = true
							case "UNI":
								gormField.Unique = true
							case "MUL":
								gormField.Index = true
							}
						case "Null":
							if value == "NO" {
								gormField.NotNull = true
							}
						case "Type":
							typeString := strings.ToLower(value)
							typeLength := ""
							tmp1 := strings.Split(typeString, "(")
							if len(tmp1) > 1 {
								typeString = tmp1[0]
								typeLength = strings.Trim(tmp1[1], ")")
							}
							gormField.Type = typeString
							gormField.TypeLength = typeLength
						case "Default":
							gormField.Default = value
						}
					}
					gormFieldMap[gormField.Column] = gormField
				}

				tableField := field.GetAndSet(u)

				var sqlExec []string

				for _, gormField := range tableField.GormMap {
					sqlGormField, ok := gormFieldMap[gormField.Column]
					if !ok {
						sqlString := "ALTER TABLE " + tableName + " ADD " + gormField.Column + " " + gormField.Type
						if gormField.TypeLength != "" {
							sqlString += "(" + gormField.TypeLength + ")"
						}
						if gormField.NotNull {
							sqlString += " NOT NULL"
						} else {
							sqlString += " NULL"
						}
						if gormField.Default != "" {
							sqlString += " DEFAULT '" + gormField.Default + "'"
						}
						if gormField.Comment != "" {
							sqlString += " COMMENT '" + gormField.Comment + "'"
						}
						sqlExec = append(sqlExec, sqlString)

						if gormField.Index {
							sqlExec = append(sqlExec, "ALTER TABLE "+tableName+" ADD INDEX ("+gormField.Column+")")
						}
						if gormField.Unique {
							sqlExec = append(sqlExec, "ALTER TABLE "+tableName+" ADD UNIQUE INDEX ("+gormField.Column+")")
						}
						continue
					}

					changeFlag := false
					if gormField.Type != sqlGormField.Type {
						changeFlag = true
					}
					if gormField.Type == sqlGormField.Type && gormField.TypeLength != sqlGormField.TypeLength {
						switch gormField.Type {
						case "int":
							if !(gormField.TypeLength == "" || gormField.TypeLength == "11" ||
								sqlGormField.TypeLength == "" || sqlGormField.TypeLength == "11") {
								changeFlag = true
							}
						case "tinyint":
							if !(gormField.TypeLength == "" || gormField.TypeLength == "4" ||
								sqlGormField.TypeLength == "" || sqlGormField.TypeLength == "4") {
								changeFlag = true
							}
						case "bigint":
							if !(gormField.TypeLength == "" || gormField.TypeLength == "20" ||
								sqlGormField.TypeLength == "" || sqlGormField.TypeLength == "20") {
								changeFlag = true
							}
						default:
							changeFlag = true
						}
					}
					if gormField.NotNull != sqlGormField.NotNull && !gormField.PrimaryKey {
						changeFlag = true
					}
					if changeFlag {
						sqlString := "ALTER TABLE " + tableName + " CHANGE " + sqlGormField.Column + " " + gormField.Column + " " + gormField.Type
						if gormField.TypeLength != "" {
							sqlString += "(" + gormField.TypeLength + ")"
						}
						if gormField.NotNull {
							sqlString += " NOT NULL"
						} else {
							sqlString += " NULL"
						}
						if gormField.Default != "" {
							sqlString += " DEFAULT '" + gormField.Default + "'"
						}
						if gormField.Comment != "" {
							sqlString += " COMMENT '" + gormField.Comment + "'"
						}
						sqlExec = append(sqlExec, sqlString)
					}

					////SHOW INDEXES FROM admin;
					if gormField.Unique != sqlGormField.Unique {
						if sqlGormField.Unique {
							var dataMap TableIndexField
							if err := db.Session(&gorm.Session{NewDB: true}).
								Raw("SHOW INDEXES FROM "+tableName+" where column_name = ? ", sqlGormField.Column).
								First(&dataMap).Error; err == nil {
								sqlExec = append(sqlExec, "ALTER TABLE "+tableName+" DROP INDEX "+dataMap.KeyName)
							}
						}
						if gormField.Unique {
							sqlExec = append(sqlExec, "ALTER TABLE "+tableName+" ADD UNIQUE INDEX ("+gormField.Column+")")
						}
					}
					if gormField.Index != sqlGormField.Index {
						if sqlGormField.Index {
							var dataMap TableIndexField
							if err := db.Session(&gorm.Session{NewDB: true}).
								Raw("SHOW INDEXES FROM "+tableName+" where column_name = ? ", sqlGormField.Column).
								First(&dataMap).Error; err == nil {
								sqlExec = append(sqlExec, "ALTER TABLE "+tableName+" DROP INDEX "+dataMap.KeyName)
							}
						}
						if gormField.Index {
							sqlExec = append(sqlExec, "ALTER TABLE "+tableName+" ADD INDEX ("+gormField.Column+")")
						}
					}

					delete(gormFieldMap, gormField.Column)
				}

				for _, v := range gormFieldMap {
					sqlExec = append(sqlExec, "ALTER TABLE "+tableName+" DROP "+v.Column)
				}

				for _, v := range sqlExec {
					db.Session(&gorm.Session{NewDB: true}).Exec(v)
				}
			}
		}
		return db.Table(tableName)
	}
}
func TableOfDB(u m, dbName string, code string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		tableName := u.TableName()
		if dbName != "" {
			tableName = dbName + "." + tableName
		}
		if code != "" {
			tableName += "_" + code
		}
		return db.Table(tableName)
	}
}
func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.Query("page"))
		if page <= 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(c.Query("pageSize"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func QueryData(db *gorm.DB, fileList []string, groupBy []string, orderBy []string, data interface{}) error {
	if db == nil {
		return errors.New("请传递已链接数据库的 DB")
	}

	if len(groupBy) > 0 {
		strGroupBy := ""
		for _, v := range groupBy {
			if strGroupBy != "" {
				strGroupBy += ","
			}
			strGroupBy += v
		}
		if strGroupBy != "" {
			db = db.Group(strGroupBy)
		}
	}

	if len(orderBy) > 0 {
		for _, v := range orderBy {
			db = db.Order(v)
		}
	}
	db = db.Select(fileList)

	return db.Find(data).Error
}
func QueryPage(db *gorm.DB, page int, pageSize int) *gorm.DB {
	if page > 1 && pageSize > 0 {
		db = db.Limit(pageSize).
			Offset((page - 1) * pageSize)
	} else {
		if pageSize > 0 {
			db = db.Limit(pageSize)
		}
	}
	return db
}
