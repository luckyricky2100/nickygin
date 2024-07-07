package service

import (
	"context"

	"nickygin.com/global"
	dao "nickygin.com/recitewords/internal/DAO"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(svc.ctx, global.DBEngine)
	return svc
}
