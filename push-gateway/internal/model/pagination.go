package model

import (
	"push-gateway/internal/service"

	"gorm.io/gorm"
)

type Pagination struct {

	// @inject_tag: json:"total"
	Total int64 `protobuf:"varint,1,opt,name=Total,proto3" json:"Total,omitempty"`
	// @inject_tag: json:"current"
	Current uint32 `protobuf:"varint,2,opt,name=Current,proto3" json:"Current,omitempty"`
	// @inject_tag: json:"pageSize"
	PageSize uint32 `protobuf:"varint,3,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
}

func NewPagination(Current, PageSize uint32) *Pagination {
	return &Pagination{
		Current:  Current,
		PageSize: PageSize,
	}
}

func (v *Pagination) GetCurrent() int {
	return int(v.Current)
}
func (v *Pagination) GetPageSize() int {
	// if v.PageSize == 0 {
	// 	v.PageSize = 20
	// }
	return int(v.PageSize)
}
func (v *Pagination) GetTotal() int {
	return int(v.Total)
}

func (v *Pagination) Build() *service.Pagination {
	return &service.Pagination{
		Current:  v.Current,
		PageSize: v.PageSize,
		Total:    v.Total,
	}
}
func (v *Pagination) Paginate() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if v.PageSize == 0 {
			v.PageSize = 20
		}
		if v.Current == 0 {
			v.Current = 1
		}
		offset := (v.Current - 1) * v.PageSize
		if offset < 0 {
			offset = 0
		}
		return db.Offset(int(offset)).Limit(int(v.PageSize))
	}
}
