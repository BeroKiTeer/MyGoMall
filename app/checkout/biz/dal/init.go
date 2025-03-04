package dal

import (
	"checkout/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
