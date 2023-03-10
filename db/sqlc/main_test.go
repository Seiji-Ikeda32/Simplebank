package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Seiji-Ikeda32/simplebank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("connot load config:", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	defer testDB.Close()
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	testQueries = New(testDB)

	os.Exit(m.Run())
}
