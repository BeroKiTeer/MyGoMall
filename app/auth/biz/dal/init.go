package dal

import (
	"auth/biz/dal/redis"
)

func Init() {
	redis.Init()
	//mysql.Init()
}
