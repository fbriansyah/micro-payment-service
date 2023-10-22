package main

import (
	httpclient "github.com/fbriansyah/micro-payment-service/internal/adapter/client/http"
	"github.com/fbriansyah/micro-payment-service/internal/adapter/postgresdb"
	rabbitmq "github.com/fbriansyah/micro-payment-service/internal/adapter/rabitmq"
	grpcserver "github.com/fbriansyah/micro-payment-service/internal/adapter/server/grpc"
	"github.com/fbriansyah/micro-payment-service/internal/application"
	"github.com/fbriansyah/micro-payment-service/util"
	"github.com/rs/zerolog/log"
)

func main() {
	config, err := util.LoadConfig("./")
	if err != nil {
		log.Fatal().Msgf("cannot load config: %s", err.Error())
	}

	sqlDB := connectToDB(config.DBDriver, config.DBSource)
	if sqlDB == nil {
		log.Fatal().Msgf("cannot connect to db: %s", err.Error())
	}

	runDBMigration(config.MigrationURL, config.DBSource)
	databaseAdapter := postgresdb.NewDatabaseAdapter(sqlDB)

	billerAdapter := httpclient.NewHttpAdapter(config.BillerEndpoint)

	// setup event broker
	amqClient, err := connectToRabbit(config.EventBrokerAddress)
	if err != nil {
		log.Fatal().Msgf("cannot connect to event broker: %v", err)
	}

	// create event emiter
	eventEmiter, err := rabbitmq.NewEmitter(amqClient)
	if err != nil {
		log.Fatal().Msgf("cannot create event emiter: %v", err)
	}
	paymentService := application.NewService(billerAdapter, databaseAdapter, eventEmiter)

	serverAdapter := grpcserver.NewGrpcServerAdapter(paymentService, config.GrpcPort)
	serverAdapter.Run()
}
