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
	HandleChangePassword(c echo.Context) error
	HandleFilterUsers(c echo.Context) error
	HandleUserDetail(c echo.Context) error
}

type userHandlerImpl struct {
	loginUseCase       usecase.LoginUseCase
	registerUseCase    usecase.RegisterUseCase
	profileUseCase     usecase.ProfileUseCase
	updateUseCase      usecase.UpdateUseCase
	changePwdUseCase   usecase.ChangePasswordUseCase
	filterUsersUseCase usecase.FilterUsersUseCase
	userDetailUseCase  usecase.UserDetailUseCase
}

type UserHandlerInject struct {
	LoginUseCase       usecase.LoginUseCase
	RegisterUseCase    usecase.RegisterUseCase
	ProfileUseCase     usecase.ProfileUseCase
	UpdateUseCase      usecase.UpdateUseCase
	ChangePwdUseCase   usecase.ChangePasswordUseCase
	FilterUsersUseCase usecase.FilterUsersUseCase
	UserDetailUseCase  usecase.UserDetailUseCase
}

func NewUserHandler(inject UserHandlerInject) UserHandler {
	return &userHandlerImpl{
		loginUseCase:       inject.LoginUseCase,
		registerUseCase:    inject.RegisterUseCase,
		profileUseCase:     inject.ProfileUseCase,
		updateUseCase:      inject.UpdateUseCase,
		changePwdUseCase:   inject.ChangePwdUseCase,
		filterUsersUseCase: inject.FilterUsersUseCase,
		userDetailUseCase:  inject.UserDetailUseCase,
	}
}
