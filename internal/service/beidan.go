package service

import (
	"github.com/bitxin232/beidan/internal/model"
	"github.com/bitxin232/beidan/internal/repository"
)

type BeiDanService struct {
	repo *repository.BeiDanRepository
}

func NewBeiDanService(repo *repository.BeiDanRepository) *BeiDanService {
	return &BeiDanService{repo: repo}
}

func (s *BeiDanService) GetBeiDansByPeriod(period string) ([]model.BeiDan, error) {
	return s.repo.FindByPeriod(period)
}
