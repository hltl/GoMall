package dal

import (
	"github.com/hltl/GoMall/gomall/app/frontend/biz/dal/mysql"
	"github.com/hltl/GoMall/gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
