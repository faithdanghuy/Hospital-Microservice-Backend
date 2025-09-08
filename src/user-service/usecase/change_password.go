package usecase

import (
	"context"
	"errors"

	"github.com/Hospital-Microservice/hospital-core/log"
	"github.com/Hospital-Microservice/hospital-core/security"
	"github.com/Hospital-Microservice/user-service/repository"
	"go.uber.org/zap"
)

type ChangePasswordUseCase interface {
	Execute(ctx context.Context, userID, oldPwd, newPwd string) error
}

type changePasswordUseCaseImpl struct {
	userRepo repository.UserRepo
}

func (u changePasswordUseCaseImpl) Execute(ctx context.Context, userID, oldPwd, newPwd string) error {

	user, err := u.userRepo.FindUserByID(ctx, userID)
	if err != nil {
		log.Error("User Not Found", zap.Error(err))
		return err
	}

	if !security.VerifyPassword(oldPwd, *user.Password) {
		return errors.New("Old Password Is Incorrect")
	}

	hashPwd, err := security.HashPassword(newPwd)
	if err != nil {
		log.Error("failed to hash new password", zap.Error(err))
		return err
	}

	if err := u.userRepo.ChangePassword(context.Background(), userID, hashPwd); err != nil {
		log.Error("failed to change password", zap.Error(err))
		return err
	}

	return nil
}

func NewChangePasswordUseCase(userRepo repository.UserRepo) ChangePasswordUseCase {
	return &changePasswordUseCaseImpl{userRepo: userRepo}
}
