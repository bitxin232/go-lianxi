package model

import "gorm.io/gorm"

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

// TableName 指定模型对应的数据库表名
func (BeiDan) TableName() string {
	return "beidan"
}

// BeforeCreate GORM 的钩子函数，在创建记录前执行
func (b *BeiDan) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.Table = "beidan"
	return nil
}

// BeforeUpdate GORM 的钩子函数，在更新记录前执行
func (b *BeiDan) BeforeUpdate(tx *gorm.DB) error {
	tx.Statement.Table = "beidan"
	return nil
}

// BeforeDelete GORM 的钩子函数，在删除记录前执行
func (b *BeiDan) BeforeDelete(tx *gorm.DB) error {
	tx.Statement.Table = "beidan"
	return nil
}
