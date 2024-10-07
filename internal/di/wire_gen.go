// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/google/wire"
	"github.com/halcyon-org/kizuna/internal/adapter/api"
	"github.com/halcyon-org/kizuna/internal/adapter/controller"
	"github.com/halcyon-org/kizuna/internal/adapter/interceptor"
	"github.com/halcyon-org/kizuna/internal/adapter/repository/config"
	ent2 "github.com/halcyon-org/kizuna/internal/adapter/repository/ent"
	"github.com/halcyon-org/kizuna/internal/infrastructure/ent"
	"github.com/halcyon-org/kizuna/internal/usecase"
)

// Injectors from wire.go:

func InitializeControllerSet() (*ControllersSet, error) {
	configRepository, err := config.NewConfigRepository()
	if err != nil {
		return nil, err
	}
	client, err := ent.InitDB(configRepository)
	if err != nil {
		return nil, err
	}
	koyoInfomationRepository := ent2.NewKoyoInfomationRepository(client)
	koyoInfomationUsecase := usecase.NewKoyoInfomationUsecase(koyoInfomationRepository)
	clientDataRepository := ent2.NewClientDataRepository(client)
	clientDataUsecase := usecase.NewClientDataUsecase(clientDataRepository)
	adminUserRepository := ent2.NewAdminUserRepository(client)
	authUsecase := usecase.NewAuthUsecase(adminUserRepository)
	beLifelineServiceHandler := api.NewBeLifelineServiceHandler(koyoInfomationUsecase, clientDataUsecase, authUsecase)
	authInterceptorAdapter := interceptor.NewAuthInterceptorAdapter(authUsecase)
	beLifelineController := controller.NewBeLifelineController(beLifelineServiceHandler, authInterceptorAdapter, configRepository)
	controllersSet := &ControllersSet{
		BeLifelineController: beLifelineController,
	}
	return controllersSet, nil
}

// wire.go:

// Adapter
var repositorySet = wire.NewSet(config.NewConfigRepository, ent2.NewAdminUserRepository, ent2.NewClientDataRepository, ent2.NewKoyoInfomationRepository)

var adapterSet = wire.NewSet(api.NewBeLifelineServiceHandler, interceptor.NewAuthInterceptorAdapter)

var controllerSet = wire.NewSet(controller.NewBeLifelineController)

// Infrastructure
var infrastructureSet = wire.NewSet(ent.InitDB)

// Usecase
var usecaseSet = wire.NewSet(usecase.NewAuthUsecase, usecase.NewKoyoInfomationUsecase, usecase.NewClientDataUsecase)

type ControllersSet struct {
	BeLifelineController controller.BeLifelineController
}
