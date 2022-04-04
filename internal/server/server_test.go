package server_test

import (
	"log"
	"os"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/wrs-news/bff-api-getaway/internal/config"
)

var (
	testConfig = config.NewConfig()
)

func TestMain(m *testing.M) {
	if _, err := toml.DecodeFile("../../config/config.test.toml", testConfig); err != nil {
		log.Fatalf(err.Error())
	}

	os.Exit(m.Run())
}
