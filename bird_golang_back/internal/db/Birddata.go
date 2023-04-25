package db

import (
	"bird_golang_back/internal/util"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InsertBirdData(id string, name string, des string) error {
	db := GetDatabase()
	_, err := db.Exec(`INSERT INTO bird VALUES(?,?,?)`, id, name, des)
	if err != nil {
		return err
	}

	return nil
}
func InsertBirdDescription(id string, des string) error {
	db := GetDatabase()
	fmt.Println(id, des)
	_, err := db.Exec(`UPDATE bird SET bird_description = ? WHERE bird_id = ?`, des, id)
	if err != nil {
		return err
	}

	return nil
}

func SearchBird(key string) (util.Bird, error) {
	db := GetDatabase()
	b := util.Bird{}
	err := db.QueryRow(`SELECT bird_id,bird_name,bird_description FROM bird 
						  WHERE (bird_id LIKE ? or bird_name LIKE ?)`, key, key).Scan(&b.ID, &b.Name, &b.Description)
	return b, err
}
