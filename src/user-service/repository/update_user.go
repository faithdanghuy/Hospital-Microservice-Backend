package repository

import (
	"context"
	"reflect"

	. "github.com/Hospital-Microservice/hospital-core/gorm"
	"github.com/Hospital-Microservice/user-service/entity"
)

func (u *userRepoImpl) UpdateUser(
	ctx context.Context,
	user entity.UserEntity,
) error {
	var omitFields = OmitFields(user, func(fieldValue reflect.Value) bool {
		return fieldValue.Kind() == reflect.Ptr && fieldValue.IsNil()
	})

	return u.DB.Executor.WithContext(ctx).
		Omit(omitFields...).
		Updates(&user).Error
}
