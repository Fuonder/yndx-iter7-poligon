package factory

import (
	"fmt"
	"strings"
)

type DatabaseConnector interface {
	Query(q string) error
}

type MysqlConnector struct {
}

func newMysqlConnector(dsn string) *MysqlConnector {
	fmt.Println("Connect to mysql")
	return &MysqlConnector{}
}

func (c *MysqlConnector) Query(q string) error {
	fmt.Printf("Query to mysql: %s\n", q)
	return nil
}

type PostgresqlConnector struct {
}

func newPostgresqlConnector(dsn string) *PostgresqlConnector {
	fmt.Println("Connect to postgresql")
	return &PostgresqlConnector{}
}

func (c *PostgresqlConnector) Query(q string) error {
	fmt.Printf("Query to postgresql: %s\n", q)
	return nil
}

type SQLiteConnector struct {
}

func newSQLiteConnector(dsn string) *SQLiteConnector {
	fmt.Println("Connect to sqlite")
	return &SQLiteConnector{}
}

func (c *SQLiteConnector) Query(q string) error {
	fmt.Printf("Query to sqlite: %s\n", q)
	return nil
}

// NewConnector реализует фабричный метод.
func NewConnector(dsn string) DatabaseConnector {
	switch {
	case strings.HasPrefix(dsn, "mysql://"):
		return newMysqlConnector(dsn)
	case strings.HasPrefix(dsn, "postgresql://"):
		return newPostgresqlConnector(dsn)
	case strings.HasPrefix(dsn, "sqlite3"):
		return newSQLiteConnector(dsn)
	default:
		panic(fmt.Sprintf("unknown dsn protocol: %s", dsn))
	}
}

func Example1() {
	mysql := NewConnector("mysql://...")
	mysql.Query("SELECT something FROM list")

	pg := NewConnector("postgresql://...")
	pg.Query("SELECT something FROM list")
}
