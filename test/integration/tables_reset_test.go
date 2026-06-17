package integration

import (
	"database/sql"
	"fmt"
)

func (s *Suite) ResetTables() {
	const (
		driverName = "mysql"
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

	_, err = db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	s.NoError(err)

	tables := []string{
		"users",
		"teams",
		"team_members",
		"tasks",
		"task_history",
		"task_comments",
	}

	for _, table := range tables {
		_, err = db.Exec("TRUNCATE TABLE " + table)
		s.NoError(err)
	}

	_, err = db.Exec("SET FOREIGN_KEY_CHECKS = 1")
	s.NoError(err)
}
