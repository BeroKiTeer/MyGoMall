package dal

import (
	"stock/biz/dal/mysql"
	"stock/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
