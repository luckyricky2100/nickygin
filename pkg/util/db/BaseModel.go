package db

import (
	"database/sql"
	"time"
)

type BaseModel struct {
	Id        uint32
	UserId    sql.NullInt32
	CreateOn  *time.Time
	CreateBy  *time.Time
	UpdatedOn sql.NullTime
	DeletedOn sql.NullTime
	IsDel     bool
}

func (model *BaseModel) CompareModel(id uint32) bool {
	return model.Id == id
}
