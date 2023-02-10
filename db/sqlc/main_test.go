package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/maan19/bank-app-go/util"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.Loadconfig("../..")
	if err != nil {
		log.Fatal("Error loading configs", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
