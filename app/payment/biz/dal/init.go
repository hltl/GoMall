package dal

import (
	"github.com/hltl/GoMall/app/payment/biz/dal/mysql"
	// "github.com/hltl/GoMall/app/payment/biz/dal/redis"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
