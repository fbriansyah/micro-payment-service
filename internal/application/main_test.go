package application

import (
	"database/sql"
	"log"
	"os"
	"testing"

	httpclient "github.com/fbriansyah/micro-payment-service/internal/adapter/client/http"
	"github.com/fbriansyah/micro-payment-service/internal/adapter/postgresdb"
	"github.com/fbriansyah/micro-payment-service/util"
	_ "github.com/lib/pq"
)

var testService *PaymentService
var testQueries *postgresdb.Queries

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

	testService = NewPaymentService(billerClient, dbAdapter)

	os.Exit(m.Run())
}
