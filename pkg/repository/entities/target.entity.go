package entities

import (
	"time"

	"gorm.io/gorm"
)

type Target struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Detail    string    `json:"detail"`
	Feedback  string    `json:"feedback"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	DoneAt    time.Time `json:"done_at"`
	State     int       `json:"state"`
	Deadline  time.Time `json:"deadline"`
	GapOrder  int       `json:"gap_order"`
}

// TableName 会将 Target 的表名重写为 `td_f_target`
func (t *Target) TableName() string {
	return "td_f_target"
}

// hooks
// 创建target时，gap_order为上一个顺序最大的target加100
func (t *Target) BeforeCreate(tx *gorm.DB) error {
	var maxOrder int
	gap := 100
	row := tx.Model(&t).Select("max(gap_order)").Row()
	row.Scan(&maxOrder)
	t.GapOrder = maxOrder + gap
	return nil
}
