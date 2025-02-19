package dal

import (
	"github.com/hltl/GoMall/app/auth/biz/dal/mysql"
	"github.com/hltl/GoMall/app/auth/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
