package dal

import (
	"github.com/hltl/GoMall/app/product/biz/dal/mysql"
	"github.com/hltl/GoMall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
