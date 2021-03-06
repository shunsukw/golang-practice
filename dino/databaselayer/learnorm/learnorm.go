package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type animal struct {
	gorm.Model
	ID         int    `gorm:"primary_key;not null;unique;AUTO_INCREMENT"`
	Animaltype string `gorm:"type:TEXT"`
	Nickname   string `gorm:"type:TEXT"`
	Zone       int    `gorm:"type:INTEGER"`
	Age        int    `gorm:"type:INTEGER"`
}

func main() {
	db, err := gorm.Open("sqlite3", "dino.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.DropTableIfExists(&animal{})
	db.Table("dinos").DropTableIfExists(&animal{})
	db.AutoMigrate(&animal{})
	db.Table("dinos").AutoMigrate(&animal{})

	// insert
	a := animal{
		Animaltype: "Tyrannosaurus rex",
		Nickname:   "rex",
		Zone:       1,
		Age:        11,
	}

	db.Create(&a)
	db.Table("dinos").Create(&a)

	a = animal{
		Animaltype: "Velociraptor",
		Nickname:   "rapto",
		Zone:       2,
		Age:        15,
	}
	db.Save(&a)

	// queries
	animals := []animal{}
	db.Table("dinos").Find(&animals, "age > ?", 10)
	fmt.Println(animals)
}
