package dal

import (
	"MyGoMall/app/cart/biz/dal/mysql"
	"MyGoMall/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
