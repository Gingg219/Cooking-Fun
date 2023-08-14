package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/CloudyKit/jet/v6"
	"github.com/Gingg219/CookingRecipe/models"
	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	appName string
	server  server
	debug   bool
	errLog  *log.Logger
	infoLog *log.Logger
	view    *jet.Set
	session *scs.SessionManager
	Models models.Models
}

type server struct {
	host string
	port string
	url  string
}

type ConnectionURL struct {
	User     string
	Password string
	Host     string
	Database string
	Options  map[string]string
  }

func main() {

	migrate := flag.Bool("migrate", false, "should drop all tables")

	flag.Parse()

	server := server{
		host: "localhost",
		port: "80",
		url:  "http://localhost:80",
	}

	options := map[string]string{
		"charset": "utf8mb4",
		"timeout": "30s",
	}
	settings := mysql.ConnectionURL{
		User: "root",
		Password: "123456",
		Host: "localhost:3306",
		Database: "CookingFun",
		Options: options,
	  }

	db2, err := openDB(settings.String())
	if err != nil {
		log.Fatal(err)
	}
	defer db2.Close()

	upper, err := mysql.New(db2)
	if err != nil {
		log.Fatal(err)
	}
	defer upper.Close()

	// Application
	app := &application{
		server:  server,
		appName: "Cooking-Fun",
		debug:   true,
		infoLog: log.New(os.Stdout, "INFO\t", log.Ltime|log.Ldate|log.Lshortfile),
		errLog:  log.New(os.Stderr, "ERROR\t", log.Ltime|log.Ldate|log.Llongfile),
		Models: models.New(upper),
	}

	//Migration
	if *migrate{
		err := execMigrate(upper)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Migrated")
	}

	//JET template
	if app.debug {
		app.view = jet.NewSet(jet.NewOSFileSystemLoader("../../views"), jet.InDevelopmentMode())
	} else {
		app.view = jet.NewSet(jet.NewOSFileSystemLoader("../../views"))
	}

	//Session
	app.session = scs.New()
	app.session.Lifetime = 24 * time.Hour
	app.session.Cookie.Name = app.appName
	app.session.Cookie.Domain = app.server.host
	app.session.Cookie.Persist = true
	app.session.Cookie.SameSite = http.SameSiteStrictMode
	app.session.Store = mysqlstore.New(db2)

	if err := app.ListenAndServer(); err != nil {
		log.Fatal(err)
	}
}

func openDB(settings string) (*sql.DB, error) {

	db, err := sql.Open("mysql", settings)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func execMigrate(db db.Session) error {

	script, err := os.ReadFile("../../migrations/tables.sql")
	if err != nil {
		return err
	}

	_, err = db.SQL().Exec(string(script))

	return err
}
