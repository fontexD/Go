package mysql

import (
	"database/sql"

	"github.com/fontexd/go/api/pkg/models"
	_ "github.com/go-sql-driver/mysql"
)

func SqlConn(id string, env string) ([]models.Templatedata, error) {

	db, err := sql.Open("mysql", "health:healthAPI@tcp(192.168.10.200:3306)/healthapi")
	stmt, err := db.Prepare("SELECT * FROM hosts where Type = ? and Env = ?")
	rows, err := stmt.Query(id, env)
	if err != nil {
		return nil, err
	} else {
		defer db.Close()

		products := []models.Templatedata{}
		for rows.Next() {
			var Name string
			var Host string
			var Type string
			var Env string
			err2 := rows.Scan(&Name, &Host, &Type, &Env)
			if err2 != nil {
				return nil, err2
			} else {
				product := models.Templatedata{Name, Host, Type, Env}
				products = append(products, product)
			}
		}
		return products, nil
	}
}
