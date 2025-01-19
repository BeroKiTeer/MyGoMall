package dal

import (
	"MyGoMall/app/auth/biz/dal/mysql"
	"MyGoMall/app/auth/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
