package main

import (
	"database/sql"
	"flag"
	"net/http"
	"os"

	"github.com/TheInfamousGrim/snippetbox/internal/models"
	"github.com/charmbracelet/log"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	snippets *models.SnippetModel
}

func main() {
	//* Flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	//* DB
	db, err := openDB(*dsn)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	app := &application{
		snippets: &models.SnippetModel{DB: db},
	}

	format := "%s addr=%s"
	log.Infof(format, "starting server", *addr)

	err = http.ListenAndServe(*addr, app.routes())
	log.Fatal(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
