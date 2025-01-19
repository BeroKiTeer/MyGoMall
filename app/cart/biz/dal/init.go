package dal

import (
	"cart/biz/dal/mysql"
	"cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
