package dal

import (
	"github.com/hltl/GoMall/app/checkout/biz/dal/mysql"
	"github.com/hltl/GoMall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
