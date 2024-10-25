package handler

import (
	"net/http"

	"github.com/bitxin232/beidan/internal/service"
	"github.com/gin-gonic/gin"
)

type BeiDanHandler struct {
	svc *service.BeiDanService
}

func NewBeiDanHandler(svc *service.BeiDanService) *BeiDanHandler {
	return &BeiDanHandler{svc: svc}
}

func (h *BeiDanHandler) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

// 获取beidan数据
func (h *BeiDanHandler) GetData(c *gin.Context) {
	// 从表单中获取期数
	period := c.PostForm("period")
	beidans, err := h.svc.GetBeiDansByPeriod(period)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}

	c.HTML(http.StatusOK, "data.html", gin.H{
		"beidans": beidans,
	})
}
