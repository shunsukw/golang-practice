package databaselayer

import (
	"errors"
)

var DBTypeNotSupported = errors.New("this db type is not supported")

const (
	MYSQL uint8 = iota
	SQLITE
	POSTGRESQL
	MONGODB
)

type DinoDBHandler interface {
	GetAvailableDinos() ([]Animal, error)
	GetDinoByNickname(string) (Animal, error)
	GetDinosByType(string) ([]Animal, error)
	AddAnimal(Animal) error
	UpdateAnimal(Animal, string) error
}

type Animal struct {
	ID         int    `bson:"-"`
	AnimalType string `bson:"animal_type"`
	Nickname   string `bson:"nickname"`
	Zone       int    `bson:"zone"`
	Age        int    `bson:"age"`
}

func GetDatabaseHandler(dbtype uint8, connection string) (DinoDBHandler, error) {
	switch dbtype {
	case MYSQL:
		return NewMySQLHandler(connection)
	case POSTGRESQL:
		return NewPQHandler(connection)
	case MONGODB:
		return NewMongodbHandler(connection)
	}

	return nil, DBTypeNotSupported
}
