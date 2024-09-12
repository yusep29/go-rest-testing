package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var dbConn sqlx.DB

func initDb() {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=postgres port=5432  password=awsedrftgyhu1234 host=database-1.cx2e4smagi79.eu-north-1.rds.amazonaws.com")
	if err != nil {
		log.Fatalln(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfuly connected")
	}

	dbConn = *db
	// defer db.Close()
}

func dbGetUser() []User {
	arrUser := make([]User, 0)
	user := User{}
	rows, _ := dbConn.Queryx("SELECT id, username, password FROM ms_user")
	i := 0
	for rows.Next() {
		err := rows.StructScan(&user)
		if err != nil {
			log.Fatalln(err)
		}
		arrUser = append(arrUser, user)
		i++
	}
	return arrUser
}

// https://github.com/jmoiron/sqlx
func dbAddUser(user *User) {
	tx := dbConn.MustBegin()
	tx.MustExec("INSERT INTO ms_user VALUES ($1, $2, $3)", user.Id, user.Name, user.PasswordHash)
	tx.Commit()
}

func dbDeleteUser(id string) {
	tx := dbConn.MustBegin()
	tx.MustExec("delete from ms_user where id = $1", id)
	tx.Commit()
}
