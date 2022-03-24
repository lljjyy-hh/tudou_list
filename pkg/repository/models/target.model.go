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
	GetAnyTarget(condition map[string]interface{}) ([]*entities.Target, error)
}

type TargetUpdater interface {
	// 更新一条记录
	SetOneTarget(condition map[string]interface{}) error
}

type TargetCreater interface {
	AddOneTarget(tPtr *entities.Target) error
}

type Target struct {
}

// 获取一条记录
func (t *Target) GetOneTarget(id uint) (*entities.Target, error) {
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

	result.Row().Scan(&et)

	return &et, nil
}

// 获取多条记录
func (t *Target) GetAnyTarget(condition map[string]interface{}) ([]*entities.Target, error) {
	return nil, nil
}
func (t *Target) SetOneTarget(condition map[string]interface{}) error {
	return nil
}

// 插入一条记录
func (t *Target) AddOneTarget(et *entities.Target) error {
	d, err := database.GetDb(nil)
	if err != nil {
		return err
	}
	result := d.Db.Create(et)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
