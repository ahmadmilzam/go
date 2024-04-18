package migration

import (
	"fmt"
	"os"
	"time"

	"github.com/ahmadmilzam/go/pkg/pgclient"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Migrations interface {
	Up() error
	Down() error
	Create(title string) error
}

func CreateMigrate(databaseName string) Migrations {
	return &DBMigrations{
		sourceFile:   "migrations/",
		databaseName: databaseName,
	}
}

type DBMigrations struct {
	migrate      *migrate.Migrate
	sourceFile   string
	databaseName string
}

func (m *DBMigrations) init() error {
	if m.migrate != nil {
		fmt.Println("m.migrate not nil")
		return nil
	}

	sql := pgclient.New()

	defer sql.Close()

	sourceFile := fmt.Sprintf("file://%s", m.sourceFile)
	driver, err := postgres.WithInstance(sql.DB, &postgres.Config{})

	if err != nil {
		return err
	}

	mi, err := migrate.NewWithDatabaseInstance(sourceFile, m.databaseName, driver)
	if err != nil {
		return err
	}

	m.migrate = mi

	return nil
}

func (m *DBMigrations) Up() error {
	if err := m.init(); err != nil {
		fmt.Println("error init PG migrate")
		return err
	}

	return m.migrate.Up()
}

func (m *DBMigrations) Down() error {
	if err := m.init(); err != nil {
		return err
	}

	return m.migrate.Steps(-1)
}

func (m *DBMigrations) Create(title string) error {
	if title == "" {
		return errors.New("Title can't be empty")
	}
	fileNameUp, fileNameDown := m.generateFileName(title)

	if _, err := os.Create(fileNameUp); err != nil {
		return err
	}

	if _, err := os.Create(fileNameDown); err != nil {
		_ = os.Remove(fileNameUp)
		return err
	}

	return nil
}

func (m *DBMigrations) generateFileName(title string) (fileNameUp string, fileNameDown string) {
	now := time.Now()
	unixTime := now.Unix()

	fileNameUp = fmt.Sprintf("%s/%d_%s.up.sql", m.sourceFile, unixTime, title)
	fileNameDown = fmt.Sprintf("%s/%d_%s.down.sql", m.sourceFile, unixTime, title)

	return fileNameUp, fileNameDown
}
