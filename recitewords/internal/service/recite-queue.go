package service

type ReciteQueue struct {
	UserId         uint32 `form:"id" binding:"required,gte=1"`
	State          uint8  `form:"state,default=1" binding:"oneof=0 1"`
	ContentClassId uint32 `form:"classid" binding:"reguired,gte=1"`
}
