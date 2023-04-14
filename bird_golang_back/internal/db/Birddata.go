package db

import (
	"bird_golang_back/internal/util"

	_ "github.com/go-sql-driver/mysql"
)
func InsertBirdData (id string, name string, des string) error {
	db := GetDatabase()
	_,err := db.Exec(`INSERT INTO bird VALUES(?,?,?)`,id,name,des)
	if err != nil {
		return err
	}
	
	return nil
}

func SearchBird (key string) (util.Bird, error) {
	db := GetDatabase()
	b := util.Bird{}
	err := db.QueryRow(`SELECT idx,name,description FROM bird 
						  WHERE (idx LIKE ? or name LIKE ?)`,key,key).Scan(&b.ID,&b.Name,&b.Description)
	return b,err
}