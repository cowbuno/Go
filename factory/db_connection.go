package factory

import "fmt"

type DatabaseConnect interface {
	Query(string) ([]string, error)
	Connect() error
}

// ----PostgreSQL----

type PostgreSQLConnection struct{}

func (p *PostgreSQLConnection) Connect() error {
	fmt.Println("PostgresSql database is connected")
	return nil
}

func (p *PostgreSQLConnection) Query(qeury string) ([]string, error) {

	return []string{"PostgreSQL", "res1", "res2", "res3"}, nil
}

// ----MongoDb----
type MongoDBConnection struct{}

func (m *MongoDBConnection) Connect() error {
	fmt.Println("MongoDB database is connected")
	return nil
}

func (m *MongoDBConnection) Query(qeury string) ([]string, error) {

	return []string{"MongoDB", "res1", "res2", "res3"}, nil
}

// ----MySQLConnection----

type MySQLConnection struct{}

func (m *MySQLConnection) Connect() error {
	fmt.Println("MySQLC database is connected")
	return nil
}

func (m *MySQLConnection) Query(qeury string) ([]string, error) {

	return []string{"MySQL", "res1", "res2", "res3"}, nil
}
