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
	clientInformationRepository := ent2.NewClientInformationRepository(client)
	clientInformationUsecase := usecase.NewClientInformationUsecase(clientInformationRepository)
	koyoInformationRepository := ent2.NewKoyoInformationRepository(client)
	koyoInformationUsecase := usecase.NewKoyoInformationUsecase(koyoInformationRepository)
	externalInformationRepository := ent2.NewExternalInformationRepository(client)
	externalInformationUsecase := usecase.NewExternalInformationUsecase(externalInformationRepository)
	adminServiceHandler := api.NewAdminServiceHandler(clientInformationUsecase, koyoInformationUsecase, externalInformationUsecase)
	providerServiceHandler := api.NewProviderServiceHandler()
	externalInformationServiceHandler := api.NewExternalInformationServiceHandler()
	koyoDataRepository := ent2.NewKoyoDataRepository(client)
	koyoDataUsecase := usecase.NewKoyoDataUsecase(koyoDataRepository)
	koyoServiceHandler := api.NewKoyoServiceHandler(koyoInformationUsecase, koyoDataUsecase)
	healthServiceHandler := api.NewHealthServiceHandler()
	adminUserRepository := ent2.NewAdminUserRepository(client)
	authUsecase := usecase.NewAuthUsecase(adminUserRepository, clientInformationRepository, koyoInformationRepository, externalInformationRepository)
	authInterceptorAdapter := interceptor.NewAuthInterceptorAdapter(authUsecase)
	loggingInterceptorAdapter := interceptor.NewLoggingInterceptorAdapter()
	beLifelineController := controller.NewBeLifelineController(adminServiceHandler, providerServiceHandler, externalInformationServiceHandler, koyoServiceHandler, healthServiceHandler, authInterceptorAdapter, loggingInterceptorAdapter, configRepository)
	controllersSet := &ControllersSet{
		BeLifelineController: beLifelineController,
	}
	return controllersSet, nil
}

// wire.go:

// Adapter
var repositorySet = wire.NewSet(config.NewConfigRepository, ent2.NewAdminUserRepository, ent2.NewClientInformationRepository, ent2.NewKoyoDataRepository, ent2.NewKoyoInformationRepository, ent2.NewExternalInformationRepository)

var adapterSet = wire.NewSet(api.NewAdminServiceHandler, api.NewProviderServiceHandler, api.NewExternalInformationServiceHandler, api.NewKoyoServiceHandler, api.NewHealthServiceHandler, interceptor.NewAuthInterceptorAdapter, interceptor.NewLoggingInterceptorAdapter)

var controllerSet = wire.NewSet(controller.NewBeLifelineController)

// Infrastructure
var infrastructureSet = wire.NewSet(ent.InitDB)

// Usecase
var usecaseSet = wire.NewSet(usecase.NewAuthUsecase, usecase.NewKoyoInformationUsecase, usecase.NewKoyoDataUsecase, usecase.NewClientInformationUsecase, usecase.NewExternalInformationUsecase)

type ControllersSet struct {
	BeLifelineController controller.BeLifelineController
}
