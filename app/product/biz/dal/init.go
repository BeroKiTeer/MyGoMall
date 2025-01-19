package dal

import (
	"MyGoMall/app/product/biz/dal/mysql"
	"MyGoMall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
