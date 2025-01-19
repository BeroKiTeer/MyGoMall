package dal

import (
	"product/biz/dal/mysql"
	"product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
