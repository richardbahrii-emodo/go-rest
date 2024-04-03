package database

type Database interface {
	Connect() error
	Close() error
	InsertOne(collection string, data interface{}) (interface{}, error)
	DeleteOne(collection string, filter interface{}) error
	UpdateOne(collection string, filter interface{}, query interface{}) (interface{}, error)
	FindAll(collection string, filter interface{}) ([]interface{}, error)
}

func InitDatabase(dbType string) Database {
	switch dbType {
	case "mongo":
		return &MongoDb{}
	default:
		return nil
	}
}
