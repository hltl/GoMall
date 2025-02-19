package dal

import (
	"github.com/hltl/GoMall/app/user/biz/dal/mysql"
	"github.com/hltl/GoMall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
