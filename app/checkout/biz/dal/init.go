package dal

import (
	"checkout/biz/dal/mysql"
	"checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
