package model

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
