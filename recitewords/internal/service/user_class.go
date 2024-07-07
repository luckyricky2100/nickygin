package service

import (
	"nickygin.com/recitewords/internal/models"
)

type UserKnowledgeClassRequest struct {
	ParentId int    `form:"parent_id"`
	UserId   string `validate:"required,int"`
}

type CreateUserClass struct {
	Title            string `form:"title"  binding:"required,lt=20"`
	ParentID         int32  `form:"parent_id"  binding:"required,gte=-1"`
	UserId           string `fomr:"userid"`
	EnableToRemember bool   `form:"enable_to_memory"   binding:"required"`
}

func (svc *Service) GetClasses(param *UserKnowledgeClassRequest) ([]*models.UserClass, error) {
	return svc.dao.GetKnowledgeClass(param.UserId, param.ParentId)
}

func (svc *Service) CreateClass(param *CreateUserClass) (*models.UserClass, error) {
	return svc.dao.CreateClass(param.UserId, param.Title, int(param.ParentID), &param.EnableToRemember)
}
