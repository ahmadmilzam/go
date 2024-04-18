package pgclient

import (
	"log"
	"time"

	"github.com/ahmadmilzam/go/config"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
	sqlxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/jmoiron/sqlx"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// type client struct {
// 	DB *sqlx.DB
// }

// func (c *client) Close() {
// 	_ = c.DB.Close()
// }

func New() *sqlx.DB {
	dbConfig := config.GetDBConfig()

	sqltrace.Register("postgres", &pq.Driver{}, sqltrace.WithDBMPropagation(tracer.DBMPropagationModeFull), sqltrace.WithServiceName("ths.db"))
	db, err := sqlxtrace.Open("postgres", dbConfig.GetConnectionURI(), sqltrace.WithDBMPropagation(tracer.DBMPropagationModeFull))
	if err != nil {
		log.Fatalf("failure when opening db connection to: %s err: %v", dbConfig.GetConnectionURI(), err)
	}
	db.SetMaxIdleConns(dbConfig.Connection.MaxIdleConn)
	db.SetMaxOpenConns(dbConfig.Connection.MaxOpenConn)
	lifeTime := time.Second * time.Duration(dbConfig.Connection.MaxLifeTimeConn)
	db.SetConnMaxLifetime(lifeTime)

	return db
}
