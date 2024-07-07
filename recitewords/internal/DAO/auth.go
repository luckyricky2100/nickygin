package dao

import (
	"errors"
	"time"

	"nickygin.com/pkg/errcode"
	"nickygin.com/recitewords/internal/models"
)

func (d *Dao) GetAuth(appKey, appSecret string) (*models.RcUser, error) {
	a := d.q.RcUser
	return a.WithContext(d.ctx).Where(a.LoginName.Eq(appKey), a.LoginPwd.Eq(appSecret)).First()
}

func (d *Dao) Rigister(appKey, appSecret string, nickname string) (*models.RcUser, error) {
	var now = time.Now()
	m := &models.RcUser{LoginName: appKey, LoginPwd: appSecret, Nickname: nickname, CreatedOn: &now, ModifiedOn: &now, DeletedOn: &now}
	a := d.q.RcUser
	_, err := a.WithContext(d.ctx).Where(a.LoginName.Eq(appKey)).First()
	if err == nil {
		return nil, errors.New(errcode.RigisterAlreadyExist.Msg)
	}
	err = a.WithContext(d.ctx).Create(m)
	if d.db.RowsAffected <= 0 {
		return nil, err
	}
	m, err = a.WithContext(d.ctx).Where(a.LoginName.Eq(appKey)).First()
	if err != nil {
		return nil, err
	}

	return m, nil
}
