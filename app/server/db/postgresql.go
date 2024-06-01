package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"sync"
)

type Database struct {
	user     string
	password string
	dbname   string
	host     string
	port     string
	Conn     *sql.DB
}

var once sync.Once
var DB *Database

func NewDatabase() *Database {
	created := false
	db := &Database{}
	once.Do(func() {
		db.user = os.Getenv("POSTGRES_USER")
		db.password = os.Getenv("POSTGRES_PASSWORD")
		db.dbname = os.Getenv("POSTGRES_DBNAME")
		db.host = os.Getenv("POSTGRES_HOST")
		db.port = os.Getenv("POSTGRES_PORT")
		created = true
	})
	if !created {
		return nil
	}
	return db
}

func (db *Database) Connect() error {
	ConnectionData := db.ConnectString()
	Conn, err := sql.Open("postgres", ConnectionData)
	if err != nil {
		return err
	}
	db.Conn = Conn
	return nil
}

func (db *Database) ConnectString() string {
	return fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		db.user, db.password, db.dbname, db.host, db.port,
	)
}

func (db *Database) createIndices() error {
	var err error
	_, err = db.Conn.Exec("CREATE INDEX IF NOT EXISTS ExpressionId ON expressions(id)")
	return err
}

func (db *Database) Close() error {
	return db.Conn.Close()
}

func (db *Database) createDatabase() error {
	var err error
	_, err = db.Conn.Exec(
		"DO $$BEGIN IF NOT EXISTS(SELECT FROM pg_database WHERE datname = 'orchestra') THEN CREATE DATABASE orchestra; END IF; END$$",
	)
	return err
}

func (db *Database) createTypes() error {
	var err error
	_, err = db.Conn.Exec(
		"DO $$BEGIN IF NOT EXISTS(SELECT FROM pg_type WHERE typname = 'expression_statuses') THEN CREATE TYPE expression_statuses as ENUM ('waiting', 'done'); END IF; END $$",
	)
	if err != nil {
		return err
	}
	_, err = db.Conn.Exec(
		"DO $$BEGIN IF NOT EXISTS(SELECT FROM pg_type WHERE typname = 'operations') THEN CREATE TYPE operations as ENUM ('/', '*', '+', '-'); END IF; END $$",
	)
	return err
}

func (db *Database) createTables() error {
	var err error
	_, err = db.Conn.Exec(
		"CREATE TABLE IF NOT EXISTS Expressions(id serial PRIMARY KEY, result text, expression text, status expression_statuses default 'waiting', blocked bool default false)",
	)
	if err != nil {
		return err
	}
	_, err = db.Conn.Exec(
		"CREATE TABLE IF NOT EXISTS Task(id serial PRIMARY KEY, arg1 float8, arg2 float8, operation operations NOT NULL, durationMS int, expression_id int REFERENCES Expressions(id) ON DELETE CASCADE ON UPDATE CASCADE)",
	)
	return err
}

func (db *Database) PrepareDatabase() error {
	var err error
	log.Println("Started preparing database")
	err = db.Connect()
	if err != nil {
		return err
	}
	//err = db.createDatabase()
	//if err != nil {
	//	return err
	//}
	err = db.createTypes()
	if err != nil {
		return err
	}
	err = db.createTables()
	if err != nil {
		return err
	}
	err = db.createIndices()
	if err != nil {
		return err
	}
	log.Println("Database prepared")
	return nil
}
