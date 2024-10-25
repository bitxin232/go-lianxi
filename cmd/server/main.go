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
	"gorm.io/gorm/logger"
)

func main() {
	// 加载配置
	cfg := config.Load()

	// 连接数据库
	db, err := gorm.Open(mysql.Open(cfg.DatabaseDSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 设置 GORM 日志级别为 Info
	})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	// 打印数据库连接信息
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database connection:", err)
	}
	log.Println("Database connected successfully. Max open connections:", sqlDB.Stats().MaxOpenConnections)

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
	log.Println("Server starting on", cfg.ServerAddress)
	if err := r.Run(cfg.ServerAddress); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
