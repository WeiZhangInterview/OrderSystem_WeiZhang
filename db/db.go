package db

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "log"

var CONN *sql.DB

func init() {
	var err error
	//CONN, err = sql.Open("mysql", "root:12341234@tcp(127.0.0.1:3306)/order_system?parseTime=true")
	CONN, err := sql.Open("mysql", "order_system:wei_zhang@tcp(db:3306)/order_system")
	//NOTE: Avoid Tmie Wait MYSQL CONN
	CONN.SetMaxOpenConns(20)
	CONN.SetMaxIdleConns(20)

	if err != nil {
		log.Fatal(err.Error())
	}
	err = CONN.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

}
