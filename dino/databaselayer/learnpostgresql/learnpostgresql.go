package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type animal struct {
	id         int
	animalType string
	nickname   string
	zone       int
	age        int
}

func main() {
	db, err := sql.Open("postgres", "user=postgres dbname=dino sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// general query
	rows, err := db.Query("select * from animals where age > $1", 5)
	handleRows(rows, err)

	// insert a single row
	result, err := db.Exec("Insert into animals (animal_type, nickname, zone, age) values ('Carnotaurus', 'Carno', $1, $2)", 3, 22)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.RowsAffected())

	testTransaction(db)
}

func handleRows(rows *sql.Rows, err error) {
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	animals := []animal{}
	for rows.Next() {
		a := animal{}
		err := rows.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
		if err != nil {
			log.Fatal(err)
		}
		animals = append(animals, a)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(animals)
}

func testTransaction(db *sql.DB) {
	fmt.Println("Transactions ...")
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(" select * from animals where age > $1 ")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(5)
	handleRows(rows, err)
	rows, err = stmt.Query(17)
	handleRows(rows, err)
	results, err := tx.Exec("Update animals set age = $1 where id = $2", 18, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(results.RowsAffected())
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
