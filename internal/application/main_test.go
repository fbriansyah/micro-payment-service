package application

import (
	"database/sql"
	"log"
	"os"
	"testing"

	httpclient "github.com/fbriansyah/micro-payment-service/internal/adapter/client/http"
	"github.com/fbriansyah/micro-payment-service/internal/adapter/postgresdb"
	rabbitmq "github.com/fbriansyah/micro-payment-service/internal/adapter/rabitmq"
	"github.com/fbriansyah/micro-payment-service/util"
	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	testService *Service
	testQueries *postgresdb.Queries
)

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	billerClient := httpclient.NewHttpAdapter(config.BillerEndpoint)

	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = postgresdb.New(db)
	dbAdapter := postgresdb.NewDatabaseAdapter(db)

	// setup event broker
	amqClient, err := amqp.Dial(config.EventBrokerAddress)
	if err != nil {
		if err != nil {
			log.Fatal("cannot connect rabbit mq:", err)
		}
	}

	// create event emiter
	eventEmiter, err := rabbitmq.NewEmitter(amqClient)
	if err != nil {
		log.Fatal("cannot create event emiter", err)
	}

	testService = NewService(billerClient, dbAdapter, eventEmiter)

	os.Exit(m.Run())
}
