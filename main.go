package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BeiDan struct {
	ID          uint   `gorm:"primaryKey"`
	Period      string `gorm:"column:period"`
	MatchNumber string `gorm:"column:match_number"`
	Event       string `gorm:"column:event"`
	Deadline    string `gorm:"column:deadline"`
	HomeTeam    string `gorm:"column:home_team"`
	Handicap    string `gorm:"column:handicap"`
	AwayTeam    string `gorm:"column:away_team"`
}

func main() {
	// 连接数据库
	dsn := "root:123@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败")
	}

	// 连接数据库是否成功
	fmt.Println("数据库连接成功")

	// 创建Gin引擎
	r := gin.Default()

	// 加载HTML模板
	r.LoadHTMLGlob("templates/*")

	// 处理首页请求
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	// 处理获取数据请求
	r.POST("/getData", func(c *gin.Context) {
		period := c.PostForm("period")

		var beidans []BeiDan
		result := db.Where("period = ?", period).Find(&beidans)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
			return
		}

		c.HTML(http.StatusOK, "data.html", gin.H{
			"beidans": beidans,
		})
	})

	// 运行服务器
	r.Run(":8080")
}
