package mysqlConn

import (
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

func NewDatabase() *Database {
	return &Database{}
}

type Database struct {
	username string
	password string
	host     string
	port     int
	database string

	db *gorm.DB

	TableList []string
	TableMap  map[string]*DatabaseTable
}
type DatabaseTable struct {
	createSqlString string
	FieldList       []string
	FieldMap        map[string]*DatabaseTableField
}
type DatabaseTableField struct {
	Column        string `json:"column"`
	Type          string `json:"type"`
	TypeLength    string `json:"type_length"`
	Comment       string `json:"comment"`
	Default       string `json:"default"`
	PrimaryKey    bool   `json:"primary_key"`
	Unique        bool   `json:"unique"`
	NotNull       bool   `json:"not_null"`
	AutoIncrement bool   `json:"auto_increment"`
	Index         bool   `json:"index"`
}

func (d *Database) SetUsername(username string) *Database {
	d.username = username
	return d
}
func (d *Database) SetPassword(password string) *Database {
	d.password = password
	return d
}
func (d *Database) SetHost(host string) *Database {
	d.host = host
	return d
}
func (d *Database) SetPort(port int) *Database {
	d.port = port
	return d
}
func (d *Database) SetDatabase(database string) *Database {
	d.database = database
	return d
}

func (d *Database) DB() (*gorm.DB, error) {
	if d.db == nil {
		gormDB, err := GormDB(d.username, d.password, d.host, d.port, d.database)
		if err != nil {
			return nil, err
		}
		d.db = gormDB
	}
	return d.db.Session(&gorm.Session{NewDB: true}), nil
}
func (d *Database) Get() error {
	gormDB, err := d.DB()
	if err != nil {
		return err
	}

	rows, err := gormDB.Raw("show tables").Rows()
	if err != nil {
		return err
	}
	defer rows.Close()

	var tableList []string
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

		for i := range cols {
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
			tableList = append(tableList, value)
		}
	}

	d.TableList = tableList
	d.TableMap = make(map[string]*DatabaseTable)
	for _, tableName := range d.TableList {
		databaseTable, err := d.GetTable(tableName)
		if err != nil {
			return err
		}
		d.TableMap[tableName] = databaseTable
	}
	return nil
}
func (d *Database) GetTable(tableName string) (*DatabaseTable, error) {
	databaseTable := &DatabaseTable{}

	gormDB, err := d.DB()
	if err != nil {
		return nil, err
	}
	rows, err := gormDB.Raw("SHOW CREATE table " + tableName).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
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
			case "Create Table":
				databaseTable.createSqlString = value
			}
		}
	}

	databaseTable.FieldList, databaseTable.FieldMap, err = d.GetTableField(tableName)
	if err != nil {
		return nil, err
	}

	return databaseTable, nil
}
func (d *Database) GetTableField(tableName string) ([]string, map[string]*DatabaseTableField, error) {
	gormDB, err := d.DB()
	if err != nil {
		return nil, nil, err
	}

	var fieldList []string
	fieldMap := make(map[string]*DatabaseTableField)

	rows, err := gormDB.Raw(fmt.Sprintf("desc %s", tableName)).Rows()
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

		gormField := &DatabaseTableField{}
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
		fieldMap[gormField.Column] = gormField
		fieldList = append(fieldList, gormField.Column)
	}

	return fieldList, fieldMap, nil
}

func GetUpdateSql(newDB, oldDB *Database) []string {
	var sqlStringList []string
	for _, tableName := range newDB.TableList {
		tableNew := newDB.TableMap[tableName]
		tableOld, ok := oldDB.TableMap[tableName]
		if !ok {
			sqlStringList = append(sqlStringList, tableNew.createSqlString)
			continue
		}

		for k, newField := range tableNew.FieldMap {
			oldField, ok := tableOld.FieldMap[k]
			if !ok {
				sqlString := "ALTER TABLE " + tableName + " ADD " + newField.Column + " " + newField.Type
				if newField.TypeLength != "" {
					sqlString += "(" + newField.TypeLength + ")"
				}
				if newField.NotNull {
					sqlString += " NOT NULL"
				} else {
					sqlString += " NULL"
				}
				if newField.Default != "" {
					sqlString += " DEFAULT '" + newField.Default + "'"
				}
				if newField.Comment != "" {
					sqlString += " COMMENT '" + newField.Comment + "'"
				}
				sqlStringList = append(sqlStringList, sqlString)

				if newField.Index {
					sqlStringList = append(sqlStringList, "ALTER TABLE "+tableName+" ADD INDEX ("+newField.Column+")")
				}
				if newField.Unique {
					sqlStringList = append(sqlStringList, "ALTER TABLE "+tableName+" ADD UNIQUE INDEX ("+newField.Column+")")
				}
				continue
			}

			changeFlag := false
			if newField.Type != oldField.Type {
				changeFlag = true
			}
			if newField.Type == oldField.Type && newField.TypeLength != oldField.TypeLength {
				switch newField.Type {
				case "int":
					if !(newField.TypeLength == "" || newField.TypeLength == "11" ||
						oldField.TypeLength == "" || oldField.TypeLength == "11") {
						changeFlag = true
					}
				case "tinyint":
					if !(newField.TypeLength == "" || newField.TypeLength == "4" ||
						oldField.TypeLength == "" || oldField.TypeLength == "4") {
						changeFlag = true
					}
				case "bigint":
					if !(newField.TypeLength == "" || newField.TypeLength == "20" ||
						oldField.TypeLength == "" || oldField.TypeLength == "20") {
						changeFlag = true
					}
				default:
					changeFlag = true
				}
			}
			if newField.NotNull != oldField.NotNull && !newField.PrimaryKey {
				changeFlag = true
			}
			if changeFlag {
				sqlString := "ALTER TABLE " + tableName + " CHANGE " + newField.Column + " " + newField.Column + " " + newField.Type
				if newField.TypeLength != "" {
					sqlString += "(" + newField.TypeLength + ")"
				}
				if newField.NotNull {
					sqlString += " NOT NULL"
				} else {
					sqlString += " NULL"
				}
				if newField.Default != "" {
					sqlString += " DEFAULT '" + newField.Default + "'"
				}
				if newField.Comment != "" {
					sqlString += " COMMENT '" + newField.Comment + "'"
				}
				sqlStringList = append(sqlStringList, sqlString)
			}

			////SHOW INDEXES FROM admin;
			if newField.Unique != oldField.Unique {
				if oldField.Unique {
					gormDB, _ := oldDB.DB()
					var dataMap TableIndexField
					if err := gormDB.Raw("SHOW INDEXES FROM "+tableName+" where column_name = ? ", oldField.Column).
						First(&dataMap).Error; err == nil {
						sqlStringList = append(sqlStringList, "ALTER TABLE "+tableName+" DROP INDEX "+dataMap.KeyName)
					}
				}
				if newField.Unique {
					sqlStringList = append(sqlStringList, "ALTER TABLE "+tableName+" ADD UNIQUE INDEX ("+newField.Column+")")
				}
			}
			if newField.Index != oldField.Index {
				if oldField.Index {
					gormDB, _ := oldDB.DB()
					var dataMap TableIndexField
					if err := gormDB.
						Raw("SHOW INDEXES FROM "+tableName+" where column_name = ? ", oldField.Column).
						First(&dataMap).Error; err == nil {
						sqlStringList = append(sqlStringList, "ALTER TABLE "+tableName+" DROP INDEX "+dataMap.KeyName)
					}
				}
				if newField.Index {
					sqlStringList = append(sqlStringList, "ALTER TABLE "+tableName+" ADD INDEX ("+newField.Column+")")
				}
			}
		}
		//fmt.Println(tableOld)
		//panic("")
	}
	return sqlStringList
}
