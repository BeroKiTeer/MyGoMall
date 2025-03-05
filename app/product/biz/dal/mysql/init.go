package mysql

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"product/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		klog.Warnf("mysql连接失败: %v", err)
		panic(err)
	}
	//err := DB.AutoMigrate(&model.Product{})
	//if err != nil {
	//	klog.Fatal(err)
	//	return
	//}
}
