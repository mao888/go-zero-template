package svc

import (
	"go-zero-template/greet/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 启动Gorm支持
	db, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "tech_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})
	// 如果出错就GameOver了
	if err != nil {
		panic(err)
	}
	// 自动同步更新表结构
	//db.AutoMigrate(&models.User{})
	return &ServiceContext{Config: c, DB: db}
}
