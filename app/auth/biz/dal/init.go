package dal

import (
	"auth/biz/dal/mysql"
	"auth/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
