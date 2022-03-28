package entities

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/gorm"

	"tudou_list/pkg/repository/database"
)

const (
	TargetFail = iota - 1
	TargetPending
	TargetDone
)

type Target struct {
	Id        uint           `json:"id" gorm:"primaryKey"`
	Detail    string         `json:"detail"`
	Feedback  sql.NullString `json:"feedback"`
	CreatedBy string         `json:"created_by"`
	CreatedAt time.Time      `json:"created_at"`
	DoneAt    sql.NullTime   `json:"done_at"`
	State     int            `json:"state"`
	Deadline  sql.NullTime   `json:"deadline"`
	GapOrder  int            `json:"gap_order"`
}

// TableName 会将 Target 的表名重写为 `td_f_target`
func (t *Target) TableName() string {
	return "td_f_target"
}

// hooks
// 创建target前执行操作
func (t *Target) BeforeCreate(tx *gorm.DB) error {
	fmt.Print("BeforeCreate]")
	var maxOrder int
	gap := 100
	row := tx.Model(&t).Select("max(gap_order)").Row()
	row.Scan(&maxOrder)
	// gap_order为上一个顺序最大的target加100
	t.GapOrder = maxOrder + gap
	// state为0: pending
	t.State = TargetPending
	// 创建时间为当前时间
	t.CreatedAt = time.Now()
	// 完成时间为空
	t.DoneAt = sql.NullTime{Valid: false}
	// 完成说明为空
	t.Feedback = sql.NullString{Valid: false}
	return nil
}

// 更新target前执行操作
func (t *Target) BeforeUpdate(tx *gorm.DB) error {
	if t.State != TargetFail && t.State != TargetPending && t.State != TargetDone {
		return database.ErrInvaildState
	}
	return nil
}
