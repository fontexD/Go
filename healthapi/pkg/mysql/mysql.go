package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/fontexd/go/api/pkg/models"
	_ "github.com/go-sql-driver/mysql"
)

var connParams = map[string]string{
	"username": os.Getenv("mysql_user"),
	"password": os.Getenv("mysql_password"),
	"host":     os.Getenv("mysql_host"),
	"port":     "3306",
	"database": "healthapi",
}

func SqlConn(id string, env string) ([]models.Templatedata, error) {

	if env == "" {

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", connParams["username"], connParams["password"], connParams["host"], connParams["port"], connParams["database"])

		// Establish a connection to the database
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error connecting to database: %v\n", err)
			log.Fatal(err)
		}
		stmt, _ := db.Prepare("SELECT * FROM hosts where Type = ?")
		rows, err := stmt.Query(id)
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
					products = append(products, models.Templatedata{Name: Name, Host: Host, Type: Type, Env: Env})
				}
			}
			return products, nil
		}
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", connParams["username"], connParams["password"], connParams["host"], connParams["port"], connParams["database"])

	// Establish a connection to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to database: %v\n", err)
		log.Fatal(err)
	}
	stmt, _ := db.Prepare("SELECT * FROM hosts where Type = ? and Env = ?")
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
				products = append(products, models.Templatedata{Name: Name, Host: Host, Type: Type, Env: Env})
			}
		}
		return products, nil
	}
}

func SqlConnPods(id string, env string) ([]models.Templatedata, error) {

	if env == "" {

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", connParams["username"], connParams["password"], connParams["host"], connParams["port"], connParams["database"])

		// Establish a connection to the database
		db, err := sql.Open("mysql", dsn)
		stmt, _ := db.Prepare("SELECT * FROM pods where Type = ?")
		rows, err := stmt.Query(id)
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
					products = append(products, models.Templatedata{Name: Name, Host: Host, Type: Type, Env: Env})
				}
			}
			return products, nil
		}
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", connParams["username"], connParams["password"], connParams["host"], connParams["port"], connParams["database"])

	// Establish a connection to the database
	db, err := sql.Open("mysql", dsn)
	stmt, _ := db.Prepare("SELECT * FROM pods where Type = ? and Env = ?")
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
				products = append(products, models.Templatedata{Name: Name, Host: Host, Type: Type, Env: Env})
			}
		}
		return products, nil
	}
}
