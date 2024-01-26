package model

import (
	"errors"
	"github.com/liushuojia/go-sdk/mysql"
	"time"

	"gorm.io/gorm"
)

const TableNameFileDefault = "fileDefault"

// FileDefault mapped from table <fileDefault>
type FileDefault struct {
	db              *gorm.DB
	tableNameSuffix string

	ID                   int64          `json:"id" gorm:"column:id; primaryKey; comment:自动编号;"`                                              // 自动编号
	CreatedAt            time.Time      `json:"created_at" gorm:"column:created_at; type:DATETIME; autoCreateTime; not null; comment:创建时间;"` // 创建时间
	UpdatedAt            time.Time      `json:"updated_at" gorm:"column:updated_at; type:DATETIME; autoUpdateTime; not null; comment:更新时间;"` // 更新时间
	DeletedAt            gorm.DeletedAt `json:"-" gorm:"column:deleted_at; index; comment:删除时间;"`
	DataId               int64          `json:"data_id" gorm:"column:data_id; type:bigint; not null; index; comment:业务所属自增id;"`
	DataType             int8           `json:"data_type" gorm:"column:data_type; type:tinyint; not null; index; comment:业务所属类型id;"`
	DataCode             string         `json:"data_code" gorm:"column:data_code; type:char(32); not null; index; comment:业务所属编码;"`
	AreaCode             string         `json:"area_code" gorm:"column:area_code; type:char(32); not null; index; comment:分类所属编码;"`
	AreaType             int8           `json:"area_type" gorm:"column:area_type; type:tinyint; not null; index; comment:分类类型;"`
	AreaTypeAbbreviation string         `json:"area_type_abbreviation" gorm:"column:area_type_abbreviation; type:char(32); not null; index; comment:分类标识;"`
	Module               string         `json:"module" gorm:"column:module; type:char(16); not null; index; comment:所属模块;"`
	FileName             string         `json:"file_name" gorm:"column:file_name; type:char(32); not null; index; comment:文件名称;"`
	FileExtension        string         `json:"file_extension" gorm:"column:file_extension; type:char(32); not null; comment:文件后缀名;"`
	FileSize             int32          `json:"file_size" gorm:"column:file_size; type:int; not null; comment:文件大小;"`
	Uid                  int64          `json:"uid" gorm:"column:uid; type:bigint; not null; index; comment:用户编码;"`
	UidName              string         `json:"uid_name" gorm:"column:uid_name; type:char(32); not null; comment:用户名字;"`
	DepartmentCode       string         `json:"department_code" gorm:"column:department_code; type:char(32); not null; index; comment:用户所属部门;"`
	SavePath             string         `json:"save_path" gorm:"column:save_path; type:char(255); not null; comment:保存路径;"`
	SaveFile             string         `json:"save_file" gorm:"column:save_file; type:char(64); not null; comment:文件名称;"`
	Status               int8           `json:"status" gorm:"column:status; type:tinyint; not null; index; comment:状态 1 正常 0 停用;"`
	Weight               int32          `json:"weight" gorm:"column:weight; type:int; not null; index; comment:排序权重,越大越靠前;"`
	FileHash             string         `json:"file_hash" gorm:"column:file_hash; type:char(32); not null; comment:文件哈希;"`
}

// TableName FileDefault's table name
func (*FileDefault) TableName() string {
	return TableNameFileDefault
}

func (m *FileDefault) QueryDB(searchMap map[string]any) *gorm.DB {
	db := mysqlConn.DefaultField().Query(m, searchMap)
	return db
}
func (m *FileDefault) SetDB(db *gorm.DB) *FileDefault {
	m.db = db
	return m
}
func (m *FileDefault) GetDB() *gorm.DB {
	return mysqlConn.DefaultDB(m.db).
		Set("tableNameSuffix", m.tableNameSuffix).
		Scopes(mysqlConn.TableOfCode(m, m.tableNameSuffix)).
		Model(m).
		Table(m.TableName()).
		Select(mysqlConn.DefaultField().GetFieldList(m))
}
func (m *FileDefault) SetTableNameSuffix(tableNameCode string) *FileDefault {
	m.tableNameSuffix = tableNameCode
	return m
}
func (m *FileDefault) GetTableNameSuffix() string {
	return m.tableNameSuffix
}
func (m *FileDefault) Create() error {
	return m.GetDB().Create(m).Error
}
func (m *FileDefault) Select(id int64) error {
	if id <= 0 {
		return errors.New("ID 为空")
	}
	return m.GetDB().Take(m).Error
}
func (m *FileDefault) Change(updateMap map[string]interface{}) error {
	if m.ID <= 0 {
		return errors.New("ID 为空")
	}

	updateDataMap := mysqlConn.DefaultField().CleanUpdateMap(m, updateMap)
	if len(updateDataMap) <= 0 {
		return errors.New("无更新")
	}
	return m.GetDB().Updates(updateDataMap).Error
}
func (m *FileDefault) Update() error {
	if m.ID <= 0 {
		return errors.New("ID 为空")
	}
	return m.GetDB().Updates(m).Error
}
func (m *FileDefault) Delete() error {
	if m.ID <= 0 {
		return errors.New("ID 为空")
	}
	return m.GetDB().Delete(m).Error
}
func (m *FileDefault) Query(searchMap map[string]any, page int64, pageSize int64) (results []*FileDefault, err error) {
	db := m.QueryDB(searchMap)
	if page > 1 && pageSize > 0 {
		db = db.
			Limit(int(page)).
			Offset(int((page - 1) * pageSize))
	} else {
		if pageSize > 0 {
			db = db.Limit(int(pageSize))
		}
	}
	err = db.Find(&results).Error
	return
}
func (m *FileDefault) Size(searchMap map[string]any) (int64, error) {
	var num int64
	err := m.QueryDB(searchMap).
		Count(&num).
		Error
	return num, err
}
func (m *FileDefault) AfterFind(tx *gorm.DB) error {
	if tableNameCode, ok := tx.Get("tableNameCode"); ok {
		if value, ok := tableNameCode.(string); ok {
			m.SetTableNameSuffix(value)
		}
	}
	m.SetDB(tx)
	return nil
}
