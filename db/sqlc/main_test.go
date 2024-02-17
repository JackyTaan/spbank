package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/JackyTaan/spbank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	///Load config
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("can not load configuration", err)
	}
	///
	testDB, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("can not connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())

}
