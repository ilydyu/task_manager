package integration

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/pressly/goose/v3"
)

func (s *Suite) ResetMigrations() {
	const (
		migrationsPath = "../../migrations"
		driverName     = "mysql"
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		s.c.Repository.User,
		s.c.Repository.Password,
		s.c.Repository.Host,
		s.c.Repository.Port,
		s.c.Repository.DBName,
	)

	db, err := sql.Open(driverName, dsn)
	s.NoError(err)
	defer db.Close()

	err = db.Ping()
	s.NoError(err)

	err = goose.SetDialect(driverName)
	s.NoError(err)

	err = goose.DownTo(db, migrationsPath, 0)
	if err != nil && !errors.Is(err, goose.ErrNoNextVersion) {
		s.NoError(err)
	}

	err = goose.Up(db, migrationsPath)
	if err != nil && !errors.Is(err, goose.ErrNoNextVersion) {
		s.NoError(err)
	}
}
