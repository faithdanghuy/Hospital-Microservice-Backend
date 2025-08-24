package usecase

import (
	"context"
	"fmt"
	"strings"

	"github.com/Hospital-Microservice/hospital-core/log"
	"github.com/Hospital-Microservice/hospital-core/security"
	"go.uber.org/zap"

	"github.com/Hospital-Microservice/user-service/entity"
	"github.com/Hospital-Microservice/user-service/repository"
)

type RegisterUseCase interface {
	Execute(ctx context.Context, user *entity.UserEntity) error
}

type registerUseCaseImpl struct {
	UserRepo repository.UserRepo
}

func (r registerUseCaseImpl) Execute(ctx context.Context, user *entity.UserEntity) error {
	birthYear := user.Birthday.Year()
	rawPwd := *user.Phone + "@" + fmt.Sprintf("%d", birthYear)
	hashPwd, err := security.HashPassword(rawPwd)
	if err != nil {
		log.Error("failed to hash password", zap.Error(err))
	}
	user.Password = &hashPwd

	if err := r.UserRepo.InsertUser(ctx, user); err != nil {
		log.Error("failed to register user", zap.Error(err))
		return err
	}

	switch strings.ToLower(*user.Role) {
	case "patient":
		if err := r.UserRepo.CreateEmptyPatientProfile(ctx, user.ID); err != nil {
			log.Error("failed to create empty patient profile", zap.Error(err))
			return err
		}
	case "doctor":
		if err := r.UserRepo.CreateEmptyDoctorProfile(ctx, user.ID); err != nil {
			log.Error("failed to create empty doctor profile", zap.Error(err))
			return err
		}
	case "admin":
	default:
		log.Warn("unknown role, skipping profile creation", zap.String("role", *user.Role))
	}
	return nil
}

func NewRegisterUseCase(UserRepo repository.UserRepo) RegisterUseCase {
	return &registerUseCaseImpl{
		UserRepo: UserRepo,
	}
}
