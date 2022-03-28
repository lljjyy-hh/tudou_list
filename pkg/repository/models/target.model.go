package models

import (
	"tudou_list/pkg/repository/database"
	"tudou_list/pkg/repository/entities"
)

type ITarget interface {
	TargetSelector
	TargetUpdater
	TargetCreater
}
type TargetSelector interface {
	// 获取一条记录
	GetOneTarget(id uint) (*entities.Target, error)
	// 获取多条记录
	GetAnyTargets(condition string, data ...interface{}) ([]entities.Target, error)
}

type TargetUpdater interface {
	// 更新一条记录
	SetOneTarget(id uint, values map[string]interface{}) error
}

type TargetCreater interface {
	AddOneTarget(tPtr *entities.Target) error
}

type target struct {
}

func GetTargetModel() ITarget {
	return &target{}
}

// 获取一条记录
func (t *target) GetOneTarget(id uint) (*entities.Target, error) {
	// 获取数据库
	d, err := database.GetDb(nil)
	if err != nil {
		return nil, err
	}
	var et entities.Target
	// 根据主键找target
	result := d.Db.Find(&et, id)

	if result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {

		return nil, database.ErrNotFound
	}

	// result.Row().Scan(&et)

	return &et, nil
}

// 获取多条记录
func (t *target) GetAnyTargets(condition string, data ...interface{}) ([]entities.Target, error) {
	// 获取数据库
	d, err := database.GetDb(nil)
	if err != nil {
		return nil, err
	}
	var ets []entities.Target
	result := d.Db.Where(condition, data...).Find(&ets)
	if result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, database.ErrNotFound
	}
	return ets, nil
}

func (t *target) SetOneTarget(id uint, values map[string]interface{}) error {
	// 获取数据库
	d, err := database.GetDb(nil)
	if err != nil {
		return err
	}

	// 校验state
	if state, found := values["state"]; found {
		if state != entities.TargetDone && state != entities.TargetFail && state != entities.TargetPending {
			return database.ErrInvaildState
		}
	}

	// 根据主键更新值
	et := entities.Target{Id: id}
	result := d.Db.Model(&et).Updates(values)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return database.ErrNotFound
	} else {
		return nil
	}
}

// 插入一条记录
func (t *target) AddOneTarget(et *entities.Target) error {
	d, err := database.GetDb(nil)
	if err != nil {
		return err
	}
	result := d.Db.Create(et)
	return result.Error
}
