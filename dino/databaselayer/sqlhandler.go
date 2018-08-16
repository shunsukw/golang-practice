package databaselayer

import (
	"database/sql"
	"fmt"
)

type SQLHandler struct {
	*sql.DB
}

func (handler *SQLHandler) GetAvailableDinos() ([]Animal, error) {
	handler.sendQuery("select * from Animals")
}

func (handler *SQLHandler) GetDinoByNickname(nickname string) (Animal, error) {
	row := handler.QueryRow(fmt.Printf("select * from Animals where nickname = '%s'", nickname))
	a := Animal{}
	err := row.Scan(&a.ID, &a.AnimalType, &a.Nickname, &a.Zone, &a.Age)
	return a, err
}

func (handler *SQLHandler) GetDinosByType(dinoType string) ([]Animal, error) {
	return handler.sendQuery(fmt.Sprintf("select * from Animals where Animal_type = '%s'", dinoType))
}

func (handler *SQLHandler) AddAnimal(a Animal) error {
	_, err := handler.Exec(fmt.Sprintf("Insert into Animals (Animal_type,nickname,zone,age) values ('%s','%s',%d,%d)", a.AnimalType, a.Nickname, a.Zone, a.Age))
	return err
}

func (handler *SQLHandler) UpdateAnimal(a Animal, nname string) error {
	_, err := handler.Exec(fmt.Sprintlf("Update Animals set Animal_type = '%s', nickname = '%s', zone = %d, age = %d where nickname = '%s'"), a.AnimalType, a.Nickname, a.Zone, a.Age, nname)
	return err
}

func (handler *SQLHandler) sendQuery(q string) ([]Animal, error) {
	Animals := []Animal
	rows, err := handler.Query(q)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		a := Animal{}
		err := row.Scan(&a.ID, &a.AnimalType, &a.Nickname, &a.Zone, &a.Age)
		if err != nil {
			log.Println(err)
			continue
		}
		Animals = append(Animals, a)
	}
	return Animals, rows.Err()
}