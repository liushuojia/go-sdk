package mysqlConn

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gen"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

type CommonMethod struct {
	ID int64
}

func (m *CommonMethod) IsEmpty() bool {
	if m == nil {
		return true
	}
	return m.ID == 0
}

// TableName
//func (m CommonMethod) TableName() string {
//	return "@@table"
//}

// TableName table name with gorm NamingStrategy
//func (m CommonMethod) TableName(namer schema.Namer) string {
//	if namer == nil {
//		return "@@table"
//	}
//	return namer.TableName("@@table")
//}

type Querier interface {
	// SELECT * FROM @@table WHERE id=@id AND `deleted_at` IS NULL LIMIT 1
	//Get(id int64) (*gen.T, error) // GetByID query data by id and return it as *struct*
}

//var (
//	username = flag.String("u", "root", "-u username")
//	password = flag.String("p", "liushuojia", "-p password")
//	host     = flag.String("h", "127.0.0.1", "-h host")
//	port     = flag.Int("P", 3306, "-P port")
//	db       = flag.String("db", "my_test", "-db database")
//	out      = flag.String("out", "/Volumes/work/app-aliyun/sdk/example/gen", "-out outPath")
//)

func Build(username, password, host string, port int, db, out string) {
	// go run gen.go -c ../app.ini
	gormDB, err := GormDB(
		username,
		password,
		host,
		port,
		db,
	)
	if err != nil {
		log.Fatalln("connect fail", err)
	}

	basePath, err := filepath.Abs(filepath.Dir(out))
	if err != nil {
		log.Fatalln("our path fail", err)
	}

	os.RemoveAll(basePath + "/gen/query")
	os.RemoveAll(basePath + "/gen/model")

	g := gen.NewGenerator(gen.Config{
		OutPath:      basePath + "/query",
		ModelPkgPath: "model",
		//Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		Mode: gen.WithoutContext | gen.WithDefaultQuery, // generate mode
	})
	g.UseDB(gormDB) // reuse your gorm db

	// 自定义字段的数据类型
	// 统一数字类型为int64,兼容protobuf
	dataMap := map[string]func(detailType gorm.ColumnType) (dataType string){
		"tinyint":   func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"smallint":  func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"mediumint": func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"bigint":    func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"int":       func(detailType gorm.ColumnType) (dataType string) { return "int64" },
	}
	// 要先于`ApplyBasic`执行
	g.WithDataTypeMap(dataMap)

	var allModel []interface{}
	tables, _ := gormDB.Migrator().GetTables()
	for _, v := range tables {
		allModel = append(allModel, g.GenerateModel(v, gen.WithMethod(CommonMethod{})))
	}

	//allModel := g.GenerateAllTable()

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	//g.ApplyInterface(func(Querier) {}, CommonMethod{})

	// Generate structs from all tables of current database
	//g.GenerateAllTable()

	//g.GenerateModel("user", gen.WithMethod(CommonMethod{}))

	// Generate structs from all tables of current database
	//g.GenerateModel("user", gen.WithMethod(CommonMethod{}))
	//g.GenerateModelAs("user", "Employee")

	//g.Execute()

	//g.GenerateModel("user", gen.WithMethod(CommonMethod{}))

	// Generate structs from all tables of current database
	//g.GenerateModel("user", gen.WithMethod(CommonMethod{}))
	//g.GenerateModelAs("user", "Employee")
	//

	//g.ApplyBasic(allModel...)
	for _, v := range allModel {
		g.ApplyInterface(func(Querier) {}, v)
	}

	// Generate the code
	g.Execute()

}
