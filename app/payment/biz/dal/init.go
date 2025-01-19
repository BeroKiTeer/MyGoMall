package dal

import (
	"payment/biz/dal/mysql"
	"payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
