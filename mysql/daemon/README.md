# 由golang建立表，并同步表数据，方便分表分库
## 需自行编写model
    admin := &model.Admin{}
    admin.SetDB(db)
    admin2 := &model.Admin{}
    admin2.SetTableNameSuffix("you").SetDB(db)

