package dal

import (
	"user/biz/dal/mysql"
	"user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
