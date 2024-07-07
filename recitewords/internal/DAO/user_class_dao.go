package dao

import "nickygin.com/recitewords/internal/models"

func (d *Dao) GetKnowledgeClass(userid string, parenctid int) ([]*models.UserClass, error) {
	a := d.q.UserClass
	return a.WithContext(d.ctx).Where(a.ParentID.Eq(int32(parenctid)), a.CreatedBy.Eq(userid)).Find()
}

func (d *Dao) CreateClass(userid, title string, parentid int, enableToRemember *bool) (*models.UserClass, error) {
	a := d.q.UserClass
	user := models.UserClass{Title: title, CreatedBy: &userid, ParentID: int32(parentid), EnableToRemember: enableToRemember}
	err := a.WithContext(d.ctx).Create(&user)
	if d.db.RowsAffected < 1 {
		return nil, err
	} else {
		return &user, nil
	}
}
