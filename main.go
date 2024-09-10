package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var dbConn sqlx.DB

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getOneAlbums(c *gin.Context) {
	var albumOne = album{ID: "123", Title: "yuseo and the gang", Artist: "yuseo", Price: 12.5}
	c.IndentedJSON(http.StatusOK, albumOne)
}

func getCatFact(c *gin.Context) {
	url := "https://catfact.ninja/fact"

	resp, getErr := http.Get(url)
	if getErr != nil {
		log.Fatal(getErr)
	}
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	c.IndentedJSON(http.StatusOK, string(body))
}

func getUser(c *gin.Context) {
	arrUser := make([]User, 0)
	user := User{}
	rows, _ := dbConn.Queryx("SELECT username, password FROM ms_user")
	i := 0
	for rows.Next() {
		err := rows.StructScan(&user)
		if err != nil {
			log.Fatalln(err)
		}
		arrUser = append(arrUser, user)
		i++
	}
	c.IndentedJSON(http.StatusOK, arrUser)
}

type User struct {
	Name         string `db:"username"`
	PasswordHash string `db:"password"`
}

func main() {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=postgres port=5432  password=awsedrftgyhu1234 host=database-1.cx2e4smagi79.eu-north-1.rds.amazonaws.com")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfuly connected")
	}

	dbConn = *db

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/album-one", getOneAlbums)
	router.GET("/cat", getCatFact)
	router.GET("/user", getUser)

	router.Run(":8081")
}

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

//source : https://go.dev/doc/tutorial/web-service-gin
