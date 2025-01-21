package dal

import (
	"apis/biz/dal/mysql"
	"apis/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
