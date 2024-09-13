package admin

import "Gwen/model"

type TagForm struct {
	Id     uint   `json:"id"`
	Name   string `json:"name" validate:"required"`
	Color  uint   `json:"color" validate:"required"`
	UserId uint   `json:"user_id" validate:"required"`
}

func (f *TagForm) FromTag(group *model.Tag) *TagForm {
	f.Id = group.Id
	f.Name = group.Name
	f.Color = group.Color
	f.UserId = group.UserId
	return f
}

func (f *TagForm) ToTag() *model.Tag {
	i := &model.Tag{}
	i.Id = f.Id
	i.Name = f.Name
	i.Color = f.Color
	i.UserId = f.UserId
	return i
}

type TagQuery struct {
	UserId int `form:"user_id"`
	IsMy   int `form:"is_my"`
	PageQuery
}
