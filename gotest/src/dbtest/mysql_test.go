package dbtest

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"time"
)

// create database test
// use test
// create table airports (id int primary key, iata char(3), iaco char(4), chn_name varchar(50), eng_name varchar(50));
// CREATE USER 'lgq'@'%' IDENTIFIED BY 'lgq123456';
// GRANT ALL PRIVILEGES ON *.* TO 'lgq'@'%' WITH GRANT OPTION;
func TestMysqlUnix(t *testing.T) {
	// sql.Open may just validate its arguments without creating a connection to the database.
	// The returned sql.DB is safe for concurrent use by multiple goroutines and maintains its
	// own pool of idle connections. DB can be seen as a connection pool.
	db, err := sql.Open("mysql", "lgq:lgq123456@unix(/var/run/mysqld/mysqld.sock)/test")
	defer closeResource("db", db)
	if err != nil {
		fmt.Println("open db error: ", err)
		return
	}
	db.SetConnMaxLifetime(time.Minute*30)
	db.SetMaxIdleConns(20)   // control the size of connection pool
	db.SetMaxOpenConns(30)

	// To verify that the data source name is valid, call Ping.
	err = db.Ping()
	if err != nil {
		fmt.Println("ping error: ", err)
	}
}

func TestMysqlTcp(t *testing.T) {
	// create db
	db, err := sql.Open("mysql", "lgq:lgq123456@tcp(localhost:3306)/test")
	defer closeResource("db", db)
	if err != nil {
		fmt.Println("open db error: ", err)
		return
	}

	// ping
	err = db.Ping()
	if err != nil {
		fmt.Println("ping error: ", err)
	}
}

func TestMysqlQuery(t *testing.T) {
	// create db, no connection
	db, err := sql.Open("mysql", "lgq:lgq123456@tcp(localhost:3306)/test")
	defer closeResource("db", db)
	if err != nil {
		fmt.Println("open db error: ", err)
		return
	}

	// when you query, the db create a real connection or get a free connection from pool
	rows, err := db.Query("select * from airports")
	defer closeResource("rows", rows)
	if err != nil {
		fmt.Println("query error.")
		return
	}

	for rows.Next() {  // when rows.Next() return false, rows is automatically closed
		var (
			id   int64
			iata string
			iaco string
			chnName string
			engName string
		)
		if err := rows.Scan(&id, &iata, &iaco, &chnName, &engName); err != nil {
			fmt.Println("row error: ", err)
			continue
		}
		fmt.Printf("%2d, %3s, %4s, %20s, %30s\n", id, iata, iaco, chnName, engName)
	}
}

func TestMysqlPrepare(t *testing.T) {
	// create db
	db, err := sql.Open("mysql", "lgq:lgq123456@tcp(localhost:3306)/test")
	defer closeResource("db", db)
	if err != nil {
		fmt.Println("open db error: ", err)
		return
	}

	// create a transaction bound with a connection. must be end with call of commit or rollback
	tx, _ := db.Begin()
	// statement is safe for use concurrently
	stmt, err := tx.Prepare("insert into airports values (?, ?, ?, ?, ?)")
	defer closeResource("statement", stmt)
	if err != nil {
		fmt.Println("statement error: ", err)
		return
	}
	_, err = stmt.Exec(2, "PEK", "ZBAA", "北京首都国际机场", "Beijing Capital International Airport")
	fmt.Println("exec err1:", err)
	_, err = stmt.Exec(3, "PVG", "ZSPD", "上海浦东国际机场", "Shanghai Pudong International Airport")
	fmt.Println("exec err2:", err)

	err = tx.Commit()
	if err != nil {
		fmt.Println("commit error: ", err)
		return
	}
}
