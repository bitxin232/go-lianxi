package repository

import (
	"github.com/bitxin232/beidan/internal/model"
	"gorm.io/gorm"
)

type BeiDanRepository struct {
	db *gorm.DB
}

func NewBeiDanRepository(db *gorm.DB) *BeiDanRepository {
	return &BeiDanRepository{db: db}
}

// 根据期数查找贝单
func (r *BeiDanRepository) FindByPeriod(period string) ([]model.BeiDan, error) {
	var beidans []model.BeiDan
	result := r.db.Model(&model.BeiDan{}).Where("period = ?", period).Find(&beidans)
	return beidans, result.Error
}
