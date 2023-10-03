package httpclient

import (
	"log"
	"os"
	"testing"

	"github.com/fbriansyah/micro-payment-service/util"
)

var testAdapter *HttpAdapter

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../../../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testAdapter = NewHttpAdapter(config.BillerEndpoint)

	os.Exit(m.Run())
}
