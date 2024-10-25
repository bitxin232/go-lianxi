package main

import (
	"log"

	"github.com/bitxin232/beidan/internal/config"
	"github.com/bitxin232/beidan/internal/handler"
	"github.com/bitxin232/beidan/internal/repository"
	"github.com/bitxin232/beidan/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 加载配置
	cfg := config.Load()

	// 连接数据库
	db, err := gorm.Open(mysql.Open(cfg.DatabaseDSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database")
	}

	// 初始化存储库
	repo := repository.NewBeiDanRepository(db)

	// 初始化服务
	svc := service.NewBeiDanService(repo)

	// 初始化处理器
	h := handler.NewBeiDanHandler(svc)

	// 创建Gin引擎
	r := gin.Default()

	// 加载HTML模板
	r.LoadHTMLGlob("templates/*")

	// 设置路由
	r.GET("/", h.Index)
	r.POST("/getData", h.GetData)

	// 运行服务器
	r.Run(cfg.ServerAddress)
}
