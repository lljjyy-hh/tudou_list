package service

import (
	"fmt"
	"tudou_list/pkg/repository/entities"
	"tudou_list/pkg/repository/models"
)

type Target struct {
}

func (t *Target) NewTarget(et *entities.Target) (fb string, err error) {
	mt := models.GetTargetModel()
	err = mt.AddOneTarget(et)
	if err != nil {
		fb = fmt.Sprint(err)
	} else {
		fb = "创建新Target成功！"
	}
	return
}

func (t *Target) GetDone() ([]entities.Target, error) {
	tm := models.GetTargetModel()
	return tm.GetAnyTargets("state = ?", []interface{}{1})
}

func (t *Target) GetFail() ([]entities.Target, error) {
	tm := models.GetTargetModel()
	return tm.GetAnyTargets("state = ?", []interface{}{-1})
}

func (t *Target) GetPending() ([]entities.Target, error) {
	tm := models.GetTargetModel()
	return tm.GetAnyTargets("state = ?", []interface{}{0})
}
