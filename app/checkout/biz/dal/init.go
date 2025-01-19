package dal

import (
	"MyGoMall/app/checkout/biz/dal/mysql"
	"MyGoMall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
