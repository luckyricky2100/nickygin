package service

import (
	"errors"

	"nickygin.com/recitewords/internal/models"
)

type AuthRequest struct {
	AppKey    string `form:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" binding:"required"`
}

type RegisterRequest struct {
	AppKey    string `form:"app_key" binding:"required,max=20,min=5"`
	AppSecret string `form:"app_secret" binding:"required,max=20,min=5"`
	Nickname  string `form:"app_nickname" binding:"required,max=20,min=5"`
}

func (svc *Service) GetAuth(param *AuthRequest) (*models.RcUser, error) {
	auth, err := svc.dao.GetAuth(
		param.AppKey,
		param.AppSecret,
	)
	if err != nil {
		return nil, err
	}

	if auth.ID > 0 {
		return auth, nil
	}

	return nil, errors.New("auth info does not exist")
}
func (svc *Service) Register(param *RegisterRequest) error {
	auth, err := svc.dao.Rigister(
		param.AppKey,
		param.AppSecret,
		param.Nickname,
	)

	if err != nil {
		return err
	}

	if auth.ID > 0 {
		return nil
	}

	return errors.New("auth info does not exist")
}
