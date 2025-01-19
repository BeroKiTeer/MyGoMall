package dal

import (
	"MyGoMall/app/order/biz/dal/mysql"
	"MyGoMall/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
