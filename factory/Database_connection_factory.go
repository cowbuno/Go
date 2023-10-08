package factory

type DatabaseConnectionFactory struct {
}

func (f *DatabaseConnectionFactory) CreateDatabase(dbType string) DatabaseConnect {
	switch dbType {
	case "MySQL":
		return &MySQLConnection{}
	case "PostgreSql":
		return &PostgreSQLConnection{}
	case "MongoDB":
		return &MongoDBConnection{}
	default:
		return nil
	}

}
