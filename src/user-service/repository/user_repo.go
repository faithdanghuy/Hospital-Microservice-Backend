package repository

import (
	"context"

	"github.com/Hospital-Microservice/hospital-core/db"
	"github.com/Hospital-Microservice/hospital-core/record"

	"github.com/Hospital-Microservice/user-service/entity"
)

type UserRepo interface {
	InsertUser(ctx context.Context, user *entity.UserEntity) error
	UpdateUser(ctx context.Context, user entity.UserEntity) error
	FindUserByPhone(ctx context.Context, user entity.UserEntity) (*entity.UserEntity, error)
	FindUserByID(ctx context.Context, ID string) (*entity.UserEntity, error)
	CreateEmptyPatientProfile(ctx context.Context, userID *string) error
	CreateEmptyDoctorProfile(ctx context.Context, userID *string) error
	ChangePassword(ctx context.Context, oldPwd, newPwd string) error
	FilterUsers(ctx context.Context, pagination *record.Pagination) (*record.Pagination, error)
	GetUserDetail(ctx context.Context, ID string) (*entity.UserEntity, error)
}

type userRepoImpl struct {
	DB *db.Database
}

func NewUserRepo(db *db.Database) UserRepo {
	return &userRepoImpl{
		DB: db,
	}
}
