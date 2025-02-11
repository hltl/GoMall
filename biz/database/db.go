package database

import (
	"fmt"
	"time"

	"github.com/hltl/GoMall/biz/model"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 数据库配置参数（根据实际情况修改）
const (
	DBUser     = "lenovo236990" // 确保该用户存在并有权限
	DBPassword = "wangziwen"    // 确认密码正确
	DBHost     = "127.0.0.1"    // 使用IP避免DNS解析问题
	DBPort     = 3306           // 确认MySQL实际端口
	DBName     = "mall"         // 数据库需提前创建
)

func InitDB() *gorm.DB {
	// 构建DSN连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DBUser, DBPassword, DBHost, DBPort, DBName)

	// 添加调试模式
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 改为 logger.Info
	})
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}

	// 自动迁移模型
	err = db.AutoMigrate(
		&model.User{},
		// 可以添加其他模型...
	)
	if err != nil {
		panic("数据库迁移失败: " + err.Error())
	}

	// 测试连接
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)  // 空闲连接数
	sqlDB.SetMaxOpenConns(100) // 最大打开连接
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := sqlDB.Ping(); err != nil {
		panic("数据库连接测试失败: " + err.Error())
	}

	// 在连接成功后添加日志
	logrus.Info("数据库连接成功")

	// 在自动迁移后添加日志
	logrus.WithFields(logrus.Fields{
		"tables": []string{"users"},
	}).Info("数据库表结构已更新")

	return db
}
