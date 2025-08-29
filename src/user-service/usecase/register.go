package usecase

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Hospital-Microservice/hospital-core/log"
	"github.com/Hospital-Microservice/hospital-core/security"
	"go.uber.org/zap"

	"github.com/Hospital-Microservice/user-service/entity"
	"github.com/Hospital-Microservice/user-service/repository"

	"github.com/Hospital-Microservice/hospital-core/model/req"
	rabbit "github.com/Hospital-Microservice/hospital-core/provider"
)

type RegisterUseCase interface {
	Execute(ctx context.Context, user *entity.UserEntity) error
}

type registerUseCaseImpl struct {
	UserRepo  repository.UserRepo
	Publisher *rabbit.RabbitPublisher
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

	if user.Email != nil {
		notify := req.NotificationReq{
			ToEmails: []string{*user.Email},
			Subject:  "Welcome to Hospital System",
			Body:     fmt.Sprintf("Hello %s, your account has been created successfully. Your login is %s and your password is %s", *user.FullName, *user.Phone, rawPwd),
		}
		body, _ := json.Marshal(notify)
		if err := r.Publisher.Publish(ctx, body); err != nil {
			log.Error("failed to publish notification", zap.Error(err))
		}
	}
	return nil
}

func NewRegisterUseCase(UserRepo repository.UserRepo, Publisher *rabbit.RabbitPublisher) RegisterUseCase {
	return &registerUseCaseImpl{
		UserRepo:  UserRepo,
		Publisher: Publisher,
	}
}
