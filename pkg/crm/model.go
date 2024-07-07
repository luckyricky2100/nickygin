package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/logging/logrus"
	"gorm.io/plugin/opentelemetry/tracing"
	"nickygin.com/global"
	"nickygin.com/pkg/setting"
)

const (
	STATE_OPEN  = 1
	STATE_CLOSE = 0
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	logger := logger.New(
		logrus.NewWriter(),
		logger.Config{
			SlowThreshold: time.Millisecond,
			Colorful:      false,
			LogLevel: func() logger.LogLevel {
				//debug := os.Getenv("DEBUG"); debug == "true"
				if global.ServerSetting.RunMode == "debug" {
					//db.Logger = db.Logger.LogMode(logger.Info)
					return logger.Info
				} else {
					//db.Logger = db.Logger.LogMode(logger.Silent)
					return logger.Silent
				}
			}(),
		},
	)
	db, err := gorm.Open(databaseSetting.MySQLDialector(), &gorm.Config{Logger: logger})
	if err != nil {
		return nil, err
	}
	//set tracing
	if err := db.Use(tracing.NewPlugin()); err != nil {
		panic(err)
	}
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)

	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}

func updateTimeStampForCreateCallback(db *gorm.DB) {
	// Get the current time
	now := time.Now().Unix()
	if createdAtField, ok := db.Statement.Schema.FieldsByName["CreatedOn"]; ok {
		if !createdAtField.NotNull {
			// Set the value of the CreatedAt field
			db.Statement.SetColumn(createdAtField.DBName, now)
		}
	}
	if updatedAtField, ok := db.Statement.Schema.FieldsByName["ModifiedOn"]; ok {
		if !updatedAtField.NotNull {
			// Set the value of the UpdatedAt field
			db.Statement.SetColumn(updatedAtField.DBName, now)
		}
	}
}

func updateTimeStampForUpdateCallback(db *gorm.DB) {
	if _, ok := db.Get("gorm:update_column"); !ok {
		db.Statement.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

func deleteCallback(db *gorm.DB) {
	//实现硬删除和软删除
	if db.Statement.Error == nil {
		var extraOption string
		if str, ok := db.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := db.Statement.Schema.FieldsByName["DeletedOn"]
		isDelField, hasIsDelField := db.Statement.Schema.FieldsByName["IsDel"]
		if hasDeletedOnField && hasIsDelField {
			now := time.Now().Unix()

			db.Exec(
				"UPDATE ? SET ?=?,?=?",
				db.Statement.Table,
				deletedOnField.DBName,
				now,
				isDelField.DBName,
				1,
				"",
				//addExtraSpaceIfExist(db.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)
		} else {
			db.Exec(
				"DELETE FROM ?",
				db.Statement.Table,
				addExtraSpaceIfExist(extraOption),
			)
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
