package repository

import (
	"context"

	"github.com/Hospital-Microservice/hospital-core/record"
	"github.com/Hospital-Microservice/user-service/entity"
)

func (u *userRepoImpl) FilterUsers(ctx context.Context, pagination *record.Pagination) (*record.Pagination, error) {
	var users []entity.UserEntity
	var totalRows int64

	if err := u.DB.Executor.WithContext(ctx).
		Model(&entity.UserEntity{}).
		Count(&totalRows).Error; err != nil {
		return nil, err
	}

	offset := pagination.GetOffset()
	limit := pagination.GetLimit()
	sort := pagination.GetSort()

	if err := u.DB.Executor.WithContext(ctx).
		Order(sort).
		Limit(limit).
		Offset(offset).
		Find(&users).Error; err != nil {
		return nil, err
	}

	pagination.Rows = users
	pagination.TotalRows = totalRows
	pagination.TotalPages = int((totalRows + int64(limit) - 1) / int64(limit))

	return pagination, nil
}
