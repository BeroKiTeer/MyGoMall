package dal

import (
	"order/biz/dal/mysql"
	"order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
