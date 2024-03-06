// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/ChungKiet/cake-interview/internal/biz"
	conf2 "github.com/ChungKiet/cake-interview/internal/conf"
	"github.com/ChungKiet/cake-interview/internal/data"
	service2 "github.com/ChungKiet/cake-interview/internal/service"
	"github.com/ChungKiet/cake-interview/internal/worker/mapping_voucher"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/indikay/go-core/conf"
	"github.com/indikay/go-core/server"
	"github.com/indikay/go-core/service"
	"github.com/nats-io/nats.go"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf2.Data, logger log.Logger) (service.Service, func(), error) {
	httpServer := server.NewHTTPServer(confServer, logger)
	grpcServer := server.NewGRPCServer(confServer, logger)
	dataData, cleanup, err := data.NewData(confData)
	if err != nil {
		return nil, nil, err
	}

	natsCli, err := nats.Connect(confData.Config.NatsHost)
	if err != nil {
		cleanup()
		return nil, nil, err
	}

	js, err := natsCli.JetStream(nats.Context(context.Background()))
	if err != nil {
		cleanup()
		return nil, nil, err
	}

	err = CreateStream(js, confData.Config.NatsTopic)
	if err != nil {
		cleanup()
		return nil, nil, err
	}

	userRepo := data.NewUserRepo(dataData)
	campaignRepo := data.NewCampaignRepo(dataData)
	voucherCampaignRepo := data.NewVoucherCampaignRepo(dataData)
	userVoucherRepo := data.NewUserVoucherRepo(dataData) // TODO: init for worker
	voucherRepo := data.NewVoucherRepo(dataData)

	userUseCase := biz.NewUserUserCase(userRepo, js, confData.Config.NatsTopic)
	campaignUseCase := biz.NewCampaignUseCase(voucherRepo, voucherCampaignRepo, campaignRepo)
	voucherUseCase := biz.NewVoucherUseCase(voucherRepo)

	worker := mapping_voucher.NewWorker(js, confData.Config.NatsTopic, userVoucherRepo, campaignRepo, voucherCampaignRepo)

	greeterService := service2.NewCAKEServiceService(userUseCase, voucherUseCase, campaignUseCase)
	serviceService := initService(logger, httpServer, grpcServer, greeterService, worker)
	return serviceService, func() {
		cleanup()
	}, nil
}

// Subject name is also use for stream name
func CreateStream(jetStream nats.JetStreamContext, subject string) error {
	stream, err := jetStream.StreamInfo(subject)
	// stream not found, create it
	if stream == nil {
		_, err = jetStream.AddStream(&nats.StreamConfig{
			Name:     subject,
			Subjects: []string{subject},
		})
		if err != nil {
			return err
		}
	}
	return nil
}
