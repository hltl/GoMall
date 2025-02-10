package dal

import (
	"github.com/hltl/GoMall/app/order/biz/dal/mysql"
	"github.com/hltl/GoMall/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
