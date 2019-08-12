package dbtest

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"testing"
	"time"
)

// docker run --name postgres -h postgres -e POSTGRES_PASSWORD=postgres123 -e POSTGRES_USER postgres -e POSTGRES_DB postgres -v /opt/postgres:/var/lib/postgresql/data -d postgres:11.4
// d exec -it postgres bash
// usermod -s /bin/bash postgres
// su postgres
// psql
// CREATE USER test WITH PASSWORD 'test123';
// CREATE SCHEMA aviation AUTHORIZATION test;
func TestPostgresPing(t *testing.T) {
	db, err := sql.Open("postgres", "postgres://test:test123@172.17.48.4:5432/postgres?sslmode=disable")
	if err != nil {
		fmt.Println("open err:", err)
		return
	}
	defer closeResource("pq db", db)

	err = db.Ping()
	if err != nil {
		fmt.Println("ping err:", err)
	}
}

func TestPostgresUnixSocket(t *testing.T) {
	// postgres must run on the same host to connect via unix socket
	db, err := sql.Open("postgres",
		"dbname=postgres user=test password=test123 host=/var/run/postgresql/ port=5432 sslmode=disable")
	if err != nil {
		fmt.Println("open err:", err)
		return
	}
	defer closeResource("pq db", db)

	err = db.Ping()
	if err != nil {
		fmt.Println("ping err:", err)
		return
	}
}

func TestCreateTableInsertQuery(t *testing.T) {
	db, err := sql.Open("postgres", "postgres://test:test123@172.17.48.4:5432/postgres?sslmode=disable")
	if err != nil {
		fmt.Println("open err:", err)
		return
	}
	defer closeResource("pq db", db)

	// create table
	err = testPqExec(db, `create table aviation.airports(
    	id integer,
    	iata char(3),
    	iaco char(4),
    	chn_name varchar(60),
    	eng_name varchar(60),
    	update_time timestamptz,
    	primary key(id)
	)`)
	if err != nil {
		fmt.Println("create table err:", err)
		return
	}
	fmt.Println()

	// prepare statement and insert
	stmt, err := db.Prepare("insert into aviation.airports values ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		fmt.Println("create statement err:", err)
		return
	}
	defer closeResource("statement", stmt)

	now := time.Now()
	data := [][]interface{}{
		{1, "SHA", "ZSSS", "上海虹桥国际机场", "Shanghai Hongqiao International Airport", now},
		{4, "CAN", "ZGGG", "广州白云国际机场", "Guangzhou Baiyun International Airport", now},
		{5, "CGO", "ZHCC", "郑州新郑国际机场", "Zhengzhou Xinzheng International Airport", now},
	}
	for i := range data {
		re, err := stmt.Exec(data[i]...)
		if err != nil {
			fmt.Println("insert err:", err)
			continue
		}
		ra, err := re.RowsAffected()
		fmt.Println("row affected:", ra, err)
	}
	fmt.Println()

	// query table
	rows, err := db.Query("select * from aviation.airports")
	if err != nil {
		fmt.Println("select err:", err)
		return
	}
	defer closeResource("rows", rows)

	// get column names
	cols, err := rows.Columns()
	if err != nil {
		fmt.Println("get columns err:", err)
		return
	}
	fmt.Println("columns:", cols)

	// get column types
	types, err := rows.ColumnTypes()
	if err != nil {
		fmt.Println("get columns err:", err)
		return
	}
	for i := range types {
		fmt.Println(i, types[i].Name(), types[i].DatabaseTypeName(), types[i].ScanType())
	}
	fmt.Println()

	// scan rows
	for rows.Next() { // when rows.Next() return false, rows is automatically closed
		var (
			id      int64
			iata    string
			iaco    string
			chnName string
			engName string
			utime   time.Time
		)
		if err := rows.Scan(&id, &iata, &iaco, &chnName, &engName, &utime); err != nil {
			fmt.Println("row error: ", err)
			continue
		}
		fmt.Printf("%2d, %3s, %4s, %20s, %30s, %v\n", id, iata, iaco, chnName, engName, utime)
	}
	fmt.Println()

	// query single row
	var (
		iata    string
		iaco    string
		chnName string
	)
	row := db.QueryRow("select iata, iaco, chn_name from aviation.airports where id=$1 and iata=$2", 4, "SHA")
	err = row.Scan(&iata, &iaco, &chnName)
	if err != nil {
		fmt.Println("query row scan err:", err)
	} else {
		fmt.Println(iata, iaco, chnName)
	}
	row = db.QueryRow("select iata, iaco, chn_name from aviation.airports where id=$1 and iata=$2", 4, "CAN")
	err = row.Scan(&iata, &iaco, &chnName)
	if err != nil {
		fmt.Println("query row scan err:", err)
	} else {
		fmt.Println(iata, iaco, chnName)
	}
	fmt.Println()

	// delete table to clean
	err = testPqExec(db, "drop table aviation.airports")
	if err != nil {
		fmt.Println("create table err:", err)
		return
	}
	fmt.Println()
}

func testPqExec(db *sql.DB, sql string) error {
	result, rerr := db.Exec(sql)
	if rerr != nil {
		return rerr
	}
	ra, err := result.RowsAffected()
	fmt.Println("row affected:", ra, err)
	return rerr
}
