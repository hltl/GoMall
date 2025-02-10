package dal

import (
	"github.com/hltl/GoMall/app/cart/biz/dal/mysql"
	"github.com/hltl/GoMall/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
