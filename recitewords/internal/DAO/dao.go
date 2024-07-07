package dao

import (
	"context"

	"gorm.io/gorm"
	"nickygin.com/global"
	"nickygin.com/recitewords/internal/query"
)

type Dao struct {
	q   *query.Query
	db  *gorm.DB
	ctx context.Context
}

func New(c context.Context, engine *gorm.DB) *Dao {
	return &Dao{query.Use(global.DBEngine), engine, c}
}
