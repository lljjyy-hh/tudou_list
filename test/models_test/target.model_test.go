package models_test

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"tudou_list/pkg/repository/database"
	"tudou_list/pkg/repository/entities"
	"tudou_list/pkg/repository/models"
	"tudou_list/pkg/utils"
)

var target models.ITarget

// var dbRecorder = &database.Clogger.Recorder
var dbLogger = database.Default.LogMode(logger.Error)

func TestTarget(t *testing.T) {
	// t.Run("test get one target", testGetOneTarget)
	t.Run("test get any targets", testGetAnyTargets)
	// t.Run("test add one target", testAddOneTarget)
	// t.Run("test set one target", testSetOneTarget)
}

func testGetOneTarget(t *testing.T) {
	getOneTests := []struct {
		tid    uint
		result *entities.Target
		err    error
	}{
		{1, &entities.Target{Id: 1}, nil},
		{100, &entities.Target{Id: 100}, database.ErrNotFound},
	}

	leading := "\t\t  "
	// _db, _ := database.GetDb(nil)
	for _, tt := range getOneTests {
		rt, err := target.GetOneTarget(tt.tid)

		if err == tt.err {
			if err == nil && rt.Id != tt.result.Id {
				t.Errorf("Want %s, get %s!", utils.ToJsonString(tt.result, leading), utils.ToJsonString(rt, leading))
			} else {
				continue
			}
		} else if err != nil {
			t.Errorf("Test GetOne() Error[%s]", err)
		}
		fmt.Print("-----------------------")
	}
}

func testGetAnyTargets(t *testing.T) {
	getAnyTests := []struct {
		condition   string
		data        []interface{}
		resultState int
		err         error
	}{
		{"state = ?", []interface{}{0}, 0, nil},
		{"state = ?", []interface{}{1}, 1, nil},
		{"state = ?", []interface{}{-1}, -1, nil},
		{"state = ?", []interface{}{3}, 3, database.ErrNotFound},
	}

	// _db, _ := database.GetDb(nil)
	for _, tt := range getAnyTests {
		rts, err := target.GetAnyTargets(tt.condition, tt.data...)

		if err == tt.err {
			if err == nil {
				for _, rt := range rts {
					if rt.State != tt.resultState {
						t.Errorf("Want state[%d], get state[%d]!", tt.resultState, rt.State)
						break
					}
				}
			} else {
				continue
			}
		} else if err != nil {
			t.Errorf("Test GetAnyTargets() Error[%s]", err)
		}
		fmt.Print("-----------------------")
	}
}

func testAddOneTarget(t *testing.T) {
	fmt.Printf("\t\ttestAddOneTarget...\n")

	var testTargets []entities.Target
	testTargets = append(testTargets, entities.Target{
		Detail:    "test",
		Feedback:  sql.NullString{String: "test", Valid: true},
		CreatedBy: "test",
		CreatedAt: time.Now(),
		DoneAt:    sql.NullTime{Time: time.Now(), Valid: true},
		State:     0,
		Deadline:  sql.NullTime{Time: time.Now(), Valid: true},
	})
	for _, tt := range testTargets {
		err := target.AddOneTarget(&tt)
		leading := "\t\t  "
		if err != nil {
			t.Errorf("\t\tAdd one target[%s] \n\t\tError: %s\n", utils.ToJsonString(tt, leading), err)
		} else {
			// t.Errorf("\t\tAdd one target[%s] \n\t\tError: %s\n", utils.ToJsonString(tt, leading), err)
			fmt.Printf("\t\tSucceed to add one target[%s]\n", utils.ToJsonString(tt, leading))
		}
	}
}

func testSetOneTarget(t *testing.T) {
	fmt.Printf("\t\ttestSetOneTarget...\n")

	getOneTests := []struct {
		tid    uint
		values map[string]interface{}

		err error
	}{
		{1, map[string]interface{}{"done_at": sql.NullTime{Time: time.Now(), Valid: true}, "feedback": "test done1", "state": entities.TargetDone}, nil},
		{2, map[string]interface{}{"done_at": sql.NullTime{Time: time.Now(), Valid: true}, "feedback": "test fail1", "state": entities.TargetFail}, nil},
		{4, map[string]interface{}{"done_at": sql.NullTime{Time: time.Now(), Valid: true}, "feedback": "test fail1", "state": 10}, database.ErrInvaildState},
		{3, map[string]interface{}{"done_at": sql.NullTime{Time: time.Now(), Valid: true}, "feedback": "test fail1", "state": 0}, nil},
		{100, map[string]interface{}{"done_at": sql.NullTime{Time: time.Now(), Valid: true}, "feedback": "test fail1", "state": entities.TargetFail}, database.ErrNotFound},
	}
	for _, tt := range getOneTests {
		err := target.SetOneTarget(tt.tid, tt.values)
		leading := "\t\t  "
		if err == tt.err {
			if err == nil {
				et, err := target.GetOneTarget(tt.tid)
				// 查询出错
				if err != nil {
					t.Errorf("\t\tSet one target[%s] \n\t\tError: %s\n", utils.ToJsonString(tt, leading), err)
					continue
				}
				// 查询更新后的结果与预期不符
				if et.State != tt.values["state"] || et.Feedback.String != tt.values["feedback"] {
					t.Errorf("\t\tSet fail! Want state[%d], feedback[%s], get state[%d], feedback[%s]. Error[%s]", tt.values["state"], tt.values["feedback"], et.State, et.Feedback.String, err)
				}
			}
		} else {
			t.Errorf("\t\tSet one target[%s] \n\t\tError: %s\n", utils.ToJsonString(tt, leading), err)
		}
	}
}

func pkgSetUp() func() {

	_, err := database.GetDb(&gorm.Config{
		PrepareStmt: true,
		Logger:      dbLogger,
	})
	if err != nil {
		panic(fmt.Sprintf("Create database Error[%s]", err))
	}

	target = models.GetTargetModel()

	return func() {

	}
}

func TestMain(m *testing.M) {
	defer pkgSetUp()()
	m.Run()
}
