package gorm

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const TableNameUserClass = "user_class"

type UserClass struct {
	ID         int32      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedOn  time.Time  `gorm:"column:created_on;default:CURRENT_TIMESTAMP;comment:新建时间" json:"created_on"` // 新建时间
	CreatedBy  string     `gorm:"column:created_by;comment:创建人" json:"created_by"`                            // 创建人
	ModifiedOn *time.Time `gorm:"column:modified_on;comment:修改时间" json:"modified_on"`                         // 修改时间
	ModifiedBy string     `gorm:"column:modified_by;comment:修改人" json:"modified_by"`                          // 修改人
	DeletedOn  *time.Time `gorm:"column:deleted_on;comment:删除时间" json:"deleted_on"`                           // 删除时间
	IsDel      int32      `gorm:"column:is_del;comment:是否删除 0为未删除、1为已删除" json:"is_del"`                       // 是否删除 0为未删除、1为已删除
	Title      string     `gorm:"column:title;not null" json:"title"`
	ParentID   int32      `gorm:"column:parent_id;not null" json:"parent_id"`
}

func (*UserClass) TableName() string {
	return TableNameUserClass
}

func TestCreate(t *testing.T) {
	user := UserClass{Title: "Title" + strconv.Itoa(rand.Intn(100)), ParentID: 0}

	dsn := "root:123456@tcp(127.0.0.1:3306)/recite_english?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	result := db.Create(&user) // 通过数据的指针来创建
	fmt.Println(user.ID)
	fmt.Println(result.Error)        // 返回 error
	fmt.Println(result.RowsAffected) // 返回插入记录的条数

	if user.ID == 0 {
		t.Errorf("it is zero.")
	}
	if result.Error != nil {
		t.Errorf(result.Error.Error())
	}
	if result.RowsAffected == 0 {
		t.Errorf("created fail.")
	}
}

func TestUpdate(t *testing.T) {
	theId := 6
	user := UserClass{Title: "Title33.36", ParentID: 0}

	dsn := "root:123456@tcp(127.0.0.1:3306)/recite_english?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.First(&user, theId) // 通过数据的指针来创建

	if user.ID != int32(theId) {
		t.Errorf("selected fail")
	}
	now := time.Now()
	db.Model(&user).Updates(UserClass{ModifiedOn: &now, Title: "modify_title"})
	db.First(&user, theId)
	if user.ID != int32(theId) {
		t.Errorf("selected fail")
	}

	if user.ModifiedOn == nil {
		t.Errorf("Updated time fail.")
	}
}
