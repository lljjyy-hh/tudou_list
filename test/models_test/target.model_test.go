package models_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
	"tudou_list/pkg/repository/database"
	"tudou_list/pkg/repository/entities"
	"tudou_list/pkg/repository/models"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var target models.ITarget

// var dbRecorder = &database.Clogger.Recorder
var dbLogger = database.Default.LogMode(logger.Error)

func TestTarget(t *testing.T) {
	t.Run("test get one target", testGetOneTarget)
	// t.Run("test add one target", testAddOneTarget)
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
			continue
		} else if err != nil {
			t.Errorf("Test GetOne() Error[%s]", err)
		} else if rt.Id != tt.result.Id {
			t.Errorf("Want %s, get %s!, Error[%s]", toJson(tt.result, leading), toJson(rt, leading), err)
		}
		fmt.Print("-----------------------")
	}
}

func testAddOneTarget(t *testing.T) {
	fmt.Printf("\t\ttestAddOneTarget...\n")

	var testTargets []entities.Target
	testTargets = append(testTargets, entities.Target{
		Detail:    "test",
		Feedback:  "test",
		CreatedBy: "test",
		CreatedAt: time.Now(),
		DoneAt:    time.Now(),
		State:     0,
		Deadline:  time.Now(),
	})
	for _, tt := range testTargets {
		err := target.AddOneTarget(&tt)
		leading := "\t\t  "
		if err != nil {
			t.Errorf("\t\tAdd one target[%s] \n\t\tError: %s\n", toJson(tt, leading), err)
		} else {
			fmt.Printf("\t\tSucceed to add one target[%s]\n", toJson(tt, leading))
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

	target = &models.Target{}

	return func() {

	}
}

func toJson(s interface{}, leading string) string {
	sJSON, err := json.MarshalIndent(s, "", leading)
	if err != nil {
		panic(fmt.Sprintf("%s", err))
	}
	return string(sJSON)
}

func TestMain(m *testing.M) {
	defer pkgSetUp()()
	m.Run()
}
