package dal

import (
	"MyGoMall/app/payment/biz/dal/mysql"
	"MyGoMall/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
