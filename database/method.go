package database

import (
	"database/sql"
	"fmt"
	"net"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "myuser"
	password = "myuserpassword"
	dbname   = "mydb"
)

var config = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

var DB *sql.DB
var err error

func Connection() sql.DB{
	DB, err = sql.Open("postgres", config)
	if err != nil {
		panic(err)
	}

	return *DB
}

func AddLog(con net.Conn, token string){
	_, err := DB.Exec(
		`INSERT INTO log (server_ip, token) VALUES ($1, $2)`,
		con.RemoteAddr().String(),
		token,
	)

	if err != nil {
		panic(err)
	}
}

func AddResult(con net.Conn, peer string, content string){
	_, err := DB.Exec(
		`INSERT INTO result (peer, content) VALUES ($1, $2)`,
		peer,
		content,
	)

	if err != nil {
		panic(err)
	}
}