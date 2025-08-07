package handler

import (
	"github.com/Hospital-Microservice/user-service/usecase"
	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	HandleLogin(c echo.Context) error
	HandleRegister(c echo.Context) error
	HandleProfile(c echo.Context) error
	HandleUpdate(c echo.Context) error
}

type userHandlerImpl struct {
	loginUseCase    usecase.LoginUseCase
	registerUseCase usecase.RegisterUseCase
	profileUseCase  usecase.ProfileUseCase
	updateUseCase   usecase.UpdateUseCase
}

type UserHandlerInject struct {
	LoginUseCase    usecase.LoginUseCase
	RegisterUseCase usecase.RegisterUseCase
	ProfileUseCase  usecase.ProfileUseCase
	UpdateUseCase   usecase.UpdateUseCase
}

func NewUserHandler(inject UserHandlerInject) UserHandler {
	return &userHandlerImpl{
		loginUseCase:    inject.LoginUseCase,
		registerUseCase: inject.RegisterUseCase,
		profileUseCase:  inject.ProfileUseCase,
		updateUseCase:   inject.UpdateUseCase,
	}
}
