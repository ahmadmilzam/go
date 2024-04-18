package store

import (
	"log"
	"os"
	"testing"

	"github.com/ahmadmilzam/go/config"
	"github.com/ahmadmilzam/go/pkg/pgclient"
)

var testStore *SQLStore

func TestMain(m *testing.M) {
	_ = config.Load("config", "../../config")
	sql := pgclient.New()

	if err := sql.DB.Ping(); err != nil {
		log.Fatal("cannot ping db: ", err)
	}

	testStore = &SQLStore{
		DB:      sql,
		Queries: NewQueries(sql),
	}

	os.Exit(m.Run())
}
