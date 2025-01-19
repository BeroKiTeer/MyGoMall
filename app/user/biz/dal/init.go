package dal

import (
	"MyGoMall/app/user/biz/dal/mysql"
	"MyGoMall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
