package database

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

const driverMysql="mysql"

func NewMySQLDatabase(host, port, user, password, dbname string) (*sql.DB, error) {
	connStr := "%s:%s@tcp(%s:%s)/%s"

	db, err := sql.Open(driverMysql,fmt.Sprintf(connStr, user, password,host,port,dbname))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
    

	return db, nil
}